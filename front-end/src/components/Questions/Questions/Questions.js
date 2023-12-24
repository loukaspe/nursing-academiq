import React from "react";
import "./Questions.css";

const Questions = ({index, onAnswer, question}) => {
    return (
        <div className="questionCard">
            <div className="question">
                <span>Ερώτηση {index + 1}: {question[index].questionText}</span>
            </div>
            <hr/>
            <div className="answersList">
                {question[index].answerOptions.map((option) => {
                    return (
                        <button className="singleAnswer"
                                onClick={() => onAnswer(option.isCorrect)}
                                key={option.answerText}
                        >
                            {option.answerText}
                        </button>
                    );
                })}
            </div>
            <div className="questionButtons">
                <button className="simple" style={{float: "left"}} onClick={() => onAnswer(null)}>Προηγούμενη</button>
                <button className="simple" onClick={() => onAnswer(null)}>Eπόμενη</button>
                <button className="submit" style={{float: "right"}} onClick={() => onAnswer(null)}>Υποβολή</button>
            </div>
        </div>
    );
};

export default Questions;