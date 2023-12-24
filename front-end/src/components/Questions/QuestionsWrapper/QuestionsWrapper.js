import React, {useState} from "react";
import Questions from "../Questions/Questions";
import "./QuestionWrapper.css";
import Result from "../Result/Result";
import Circle from "../../Utilities/Circle";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faStopwatch} from "@fortawesome/free-solid-svg-icons";

const QuestionsWrapper = ({questions}) => {
    const [currentIndex, setCurrentIndex] = useState(0);
    const [quizFinished, setQuizFinished] = useState(false);
    const [score, setScore] = useState(0);

    function onAnswer(isCorrect) {
        if (isCorrect) {
            setScore((preValue) => preValue + 1);
        }

        if (currentIndex === questions.length - 1) {
            setQuizFinished(true);
        } else {
            setCurrentIndex((value) => value + 1);
        }
    }

    const restartHandler = () => {
        setCurrentIndex(0);
        setQuizFinished(false);
        setScore(0);
    };

    function renderQuestionCircles() {
        return questions.map((key, value) => {
            return (
                <Circle text={value + 1}/>
            );
        })
    }

    return (
        <div className="wrapper">
            {
                quizFinished ? (
                    <Result
                        score={score}
                        restartHandler={restartHandler}
                        questions={questions}
                    />
                ) : (
                    <React.Fragment>
                        {renderQuestionCircles()}
                        <div className="timer">
                            <FontAwesomeIcon icon={faStopwatch} size="2xl" className="timerIcon"/>
                            <span className="timerText">00:00</span>
                        </div>
                        <div className="questions">
                            <Questions
                                index={currentIndex}
                                question={questions}
                                onAnswer={onAnswer}
                            />
                        </div>
                    </React.Fragment>
                )}
        </div>
    )
};

export default QuestionsWrapper;