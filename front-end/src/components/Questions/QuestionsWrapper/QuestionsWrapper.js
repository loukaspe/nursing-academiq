import React, {useState} from "react";
import "./QuestionsWrapper.css";
import Result from "../Result/Result";


const QuestionsWrapper = ({questions}) => {
    const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
    const [selectedAnswers, setSelectedAnswers] = useState({});
    const [checkedQuestions, setCheckedQuestions] = useState({});
    const [quizFinished, setQuizFinished] = useState(false);
    const [score, setScore] = useState(0);

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
        const confirmMessage = `You have ${unansweredCount} unanswered questions. Are you sure you want to submit?`;

        if (window.confirm(confirmMessage)) {
            questions.forEach((question, index) => {
                let correctAnswer = question.answerOptions.find((answer) => answer.isCorrect);

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

    return (
        quizFinished ? (
            <Result
                score={score}
                restartHandler={restartHandler}
                questions={questions}
            />
        ) : (<div className="quiz-container">
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
                    <h2>{questions[currentQuestionIndex].questionText}</h2>
                    <hr/>
                    <ul>
                        {questions[currentQuestionIndex].answerOptions.map((answer, idx) => {
                            let className = '';

                            let correctAnswer = questions[currentQuestionIndex].answerOptions.find((answer) => answer.isCorrect);

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
                                    {answer.answerText}
                                </li>
                            );
                        })}
                    </ul>
                    {
                        checkedQuestions[currentQuestionIndex] && (
                        <div className="explanation">
                            <div className="explanationResult">
                                {selectedAnswers[currentQuestionIndex].answerText === questions[currentQuestionIndex].answerOptions.find(option => option.isCorrect).answerText
                                    ? <span className="correct">Σωστό</span>
                                    : <span className="incorrect">Λάθος</span>}
                            </div>
                            <br/>
                            <div className="explanationCorrectAnswer">
                                Σωστή Απάντηση: <strong>{questions[currentQuestionIndex].answerOptions.find(option => option.isCorrect).answerText}</strong>
                            </div>
                            <br/>
                            <div className="explanationDetails">
                                Εξήγηση: {questions[currentQuestionIndex].explanation}
                            </div>
                            <br/>
                            <div>
                                Πηγή: {questions[currentQuestionIndex].explanation}
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
        </div>)
    );
};


export default QuestionsWrapper;