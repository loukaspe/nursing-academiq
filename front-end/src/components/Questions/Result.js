import React from "react";
import "./Result.css";

const Result = ({score, questions, restartHandler}) => {
    return (
        <div className="resultWrapper">
            <div className="finalScore">
                Απαντήσατε σωστά σε {score} από τις {questions.length} ερωτήσεις.
            </div>

            <button onClick={restartHandler} className="restart">
                Επανάληψη{" "}
            </button>
        </div>
    );
};

export default Result;