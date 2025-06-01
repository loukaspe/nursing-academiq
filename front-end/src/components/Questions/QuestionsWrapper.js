import React, {useEffect, useState} from "react";
import "./QuestionsWrapper.css";
import Result from "./Result";
import {useParams} from "react-router-dom";

const QuestionsWrapper = () => {
    const [questions, setQuestions] = useState([]);
    const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
    const [selectedAnswers, setSelectedAnswers] = useState({});
    const [checkedQuestions, setCheckedQuestions] = useState({});
    const [quizFinished, setQuizFinished] = useState(false);
    const [score, setScore] = useState(0);

    const params = useParams();
    let quizID = params.quizID;

    useEffect(() => {
        const fetchQuestions = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${quizID}`

            try {
                const response = await fetch(apiUrl, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                    credentials: 'include',
                });
                const result = await response.json();
                // TODO if 401 show unauthorized
                // TODO if 500 show server error
                if (response.status === 500) {
                    throw Error(result.message);
                }

                if (response.status === 401) {
                    throw Error("unauthorized: 401");
                }

                if (result.quiz.Questions === undefined) {
                    throw Error("error getting quiz questions");
                }

                if (result.quiz.ShowSubset) {
                    setQuestions(getRandomSubset(result.quiz.Questions, result.quiz.SubsetSize))
                } else {
                    setQuestions(result.quiz.Questions);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        }

        fetchQuestions();
    }, [quizID]);
    const handleAnswerClick = (answer) => {
        setSelectedAnswers(prev => ({
            ...prev,
            [currentQuestionIndex]: answer
        }));
    };

    const handleNext = () => {
        setCurrentQuestionIndex((prev) => Math.min(prev + 1, questions.length - 1));
    };

    const handlePrevious = () => {
        setCurrentQuestionIndex((prev) => Math.max(prev - 1, 0));
    };

    const handleSubmit = () => {

        const unansweredCount = questions.length - Object.keys(selectedAnswers).length;
        const confirmMessage = `Έχετε ${unansweredCount} αναπάντητες ερωτήσεις. Θέλετε να προχωρήσετε ;`;

        if (unansweredCount === 0 || window.confirm(confirmMessage)) {
            questions.forEach((question, index) => {
                let correctAnswer = question.Answers.find((answer) => answer.IsCorrect);

                if (selectedAnswers[index] === correctAnswer) {
                    setScore((preValue) => preValue + 1);
                }
            });

            setQuizFinished(true);
        }
    };

    const handleCheck = () => {
        if (!selectedAnswers.hasOwnProperty(currentQuestionIndex)) {
            alert('Please select an answer before checking.');
            return;
        }
        setCheckedQuestions(prev => ({
            ...prev,
            [currentQuestionIndex]: true
        }));
    };

    const handleCircleClick = (index) => {
        setCurrentQuestionIndex(index);
    };

    const restartHandler = () => {
        setCurrentQuestionIndex(0);
        setQuizFinished(false);
        setSelectedAnswers({});
        setCheckedQuestions({});
        setScore(0);
    };

    if (questions.length === 0) {
        return <div> Παρακαλώ Περιμένετε ...</div>;
    }

    return (
        quizFinished ? (
            <Result
                score={score}
                restartHandler={restartHandler}
                questions={questions}
            />
        ) : (<React.Fragment>
            <div className="quiz-container">
                <div className="progress-line">
                    {questions.map((_, index) => (
                        <span
                            key={index}
                            className={`${selectedAnswers[index] ? 'completed' : ''} ${currentQuestionIndex === index ? 'current' : ''}`}
                            onClick={() => handleCircleClick(index)}
                        >
            {index + 1}
          </span>
                    ))}
                </div>
                <div className="questionCard">
                    <div className="question-section">
                        <h2>{questions[currentQuestionIndex].Text}</h2>
                        <hr/>
                        <ul>
                            {questions[currentQuestionIndex].Answers.map((answer, idx) => {
                                let className = '';

                                let correctAnswer = questions[currentQuestionIndex].Answers.find((answer) => answer.IsCorrect);

                                if (checkedQuestions[currentQuestionIndex]) {
                                    if (answer === correctAnswer) {
                                        className = 'correct';
                                    } else if (answer === selectedAnswers[currentQuestionIndex]) {
                                        className = 'incorrect';
                                    }
                                }

                                return (
                                    <li
                                        key={idx}
                                        className={`${selectedAnswers[currentQuestionIndex] === answer ? 'selected' : ''} ${className}`}
                                        onClick={() => handleAnswerClick(answer)}
                                    >
                                        {answer.Text}
                                    </li>
                                );
                            })}
                        </ul>
                        {
                            checkedQuestions[currentQuestionIndex] && (
                                <div className="Explanation">
                                    <div className="ExplanationResult">
                                        {selectedAnswers[currentQuestionIndex].Text === questions[currentQuestionIndex].Answers.find(option => option.IsCorrect).Text
                                            ? <span className="correct">Σωστό</span>
                                            : <span className="incorrect">Λάθος</span>}
                                    </div>
                                    <br/>
                                    <div className="ExplanationCorrectAnswer">
                                        Σωστή
                                        Απάντηση: <strong>{questions[currentQuestionIndex].Answers.find(option => option.IsCorrect).Text}</strong>
                                    </div>
                                    <br/>
                                    <div className="ExplanationDetails">
                                        Εξήγηση: {questions[currentQuestionIndex].Explanation}
                                    </div>
                                    <br/>
                                    <div>
                                        Πηγή: {questions[currentQuestionIndex].Source}
                                    </div>
                                </div>
                            )}
                    </div>
                    <div className="questionButtons">
                        <button className="questionsSimpleButton" onClick={handlePrevious}
                                disabled={currentQuestionIndex === 0}>
                            Προηγούμενη
                        </button>
                        <button className="questionsSimpleButton" onClick={handleNext}
                                disabled={currentQuestionIndex === questions.length - 1}>
                            Eπόμενη
                        </button>
                        <button className="questionsSubmitButton" onClick={handleCheck}>
                            Έλεγχος Απάντησης
                        </button>
                    </div>
                    <button className="questionsSubmitButton" onClick={handleSubmit}>
                        Οριστική Υποβολή Quiz
                    </button>
                </div>
            </div>
        </React.Fragment>)
    );
};


function shuffleArray(arr) {
    const a = arr.slice(); // copy it, so you don’t clobber the original
    for (let i = a.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [a[i], a[j]] = [a[j], a[i]];
    }
    return a;
}

// returns a random “slice” of length `size`
function getRandomSubset(array, size) {
    if (size >= array.length) {
        return array.slice(); // or just return array if you don’t mind mutating it
    }
    const shuffled = shuffleArray(array);
    return shuffled.slice(0, size);
}

export default QuestionsWrapper;