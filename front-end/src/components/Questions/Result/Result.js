import React from "react";
import "./Result.css";

const Result = ({ score, questions, restartHandler }) => {
    return (
        <React.Fragment>
            <div className="finalScore">
                You scored {score} out of {questions.length}
            </div>

            <button onClick={restartHandler} className="restart">
                RESTART{" "}
            </button>
        </React.Fragment>
    );
};

export default Result;