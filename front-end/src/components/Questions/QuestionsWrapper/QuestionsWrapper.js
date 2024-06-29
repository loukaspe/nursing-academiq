import React, {useState} from "react";
import "./QuestionsWrapper.css";
import Result from "../Result/Result";


const QuestionsWrapper = ({questions}) => {
    const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
    const [selectedAnswers, setSelectedAnswers] = useState({});
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

    const handleCircleClick = (index) => {
        setCurrentQuestionIndex(index);
    };

    const restartHandler = () => {
        setCurrentQuestionIndex(0);
        setQuizFinished(false);
        setSelectedAnswers({});
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
                        {questions[currentQuestionIndex].answerOptions.map((answer, idx) => (
                            <li
                                key={idx}
                                className={selectedAnswers[currentQuestionIndex] === answer ? 'selected' : ''}
                                onClick={() => handleAnswerClick(answer)}
                            >
                                {answer.answerText}
                            </li>
                        ))}
                    </ul>
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
                    <button className="questionsSubmitButton" onClick={handleSubmit}>
                        Υποβολή
                    </button>
                </div>
            </div>
        </div>)
    );
};


export default QuestionsWrapper;