import React from "react";
import "./Result.css";

const Result = ({score, questions, restartHandler}) => {
    return (
        <div className="resultWrapper">
            <div className="finalScore">
                You scored {score} out of {questions.length}
            </div>

            <button onClick={restartHandler} className="restart">
                RESTART{" "}
            </button>
        </div>
    );
};

export default Result;