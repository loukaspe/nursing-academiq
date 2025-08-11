import {useQuiz} from "../../context/QuizContext";
import {Link} from "react-router-dom";
import CreationProgressBar from "./CreationProgressBar";
import React from "react";
import "./CreateQuizStepTwo.css";
import Breadcrumb from "../Utilities/Breadcrumb";

export default function CreateQuizStepTwo() {
    const {quiz, setQuiz} = useQuiz();

    return (
        <div className="createQuizStepTwoContainer">
            <Breadcrumb actualPath={`/quizzes/create/step-two`} namePath={`/Quiz/Δημιουργία - Βήμα 2`}/>
            <CreationProgressBar/>
            <div className="createQuizStepTwoHeaderRow">
                <div className="createQuizStepTwoHeader">
                    <div className="createQuizStepTwoInfo">
                        <span className="singleChapterQuizzesPageChapterName">2. Λεπτομέρειες Quiz</span>
                    </div>
                </div>
            </div>
            <div className="createQuizStepTwoDetailsRow">
                <div className="createQuizStepTwoDetailsRowColumn">
                    <div className="createQuizStepTwoDetailsRowInputGroup">
                        <label>Όνομα Quiz</label>
                        <input type="text"
                               value={quiz.title}
                               className="createQuizStepTwoDetailsRowInputText"
                               onChange={(e) => setQuiz({...quiz, title: e.target.value || ""})}
                        />
                    </div>
                    <div className="createQuizStepTwoDetailsRowInputGroup">
                        <label>Περιγραφή</label>
                        <input type="text"
                               value={quiz.description}
                               className="createQuizStepTwoDetailsRowInputText"
                               onChange={(e) => setQuiz({...quiz, description: e.target.value || ""})}
                        />
                    </div>
                </div>
            </div>

            <div className="createQuizStepTwoButtonsContainer">
                <Link  className="createQuizStepTwoButton" to="/quizzes/create">Προηγούμενο</Link>
                <Link
                    className={`createQuizStepTwoButton ${(quiz.title.trim() === "" || quiz.description.trim() === "") ? "disabled" : ""}`}
                    to={(quiz.title && quiz.description) ? "/quizzes/create/step-three" : "#"}
                    onClick={(e) => {
                        if (quiz.title.trim() === "" || quiz.description.trim() === "") {
                            e.preventDefault();
                        }
                    }}
                >
                    Επόμενο
                </Link>
            </div>
        </div>
    );
}
