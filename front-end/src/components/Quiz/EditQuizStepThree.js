import React, {useEffect, useState} from "react";
import "./CreateQuizStepFour.css";
import {useQuiz} from "../../context/QuizContext";
import api from "../Utilities/APICaller";
import Cookies from "universal-cookie";
import EditProgressBar from "./EditProgressBar";
import Breadcrumb from "../Utilities/Breadcrumb";

const cookies = new Cookies();
const EditQuizStepThree = () => {
    const {quiz, setQuiz, resetQuiz} = useQuiz();

    let userCookie = cookies.get("user");
    let tutorID = userCookie.specificID;
    const handleSubsetSizeChange = (e) => {
        const value = e.target.value;
        const numValue = parseInt(value, 10);

        if (!isNaN(numValue) && numValue <= quiz.questions.length) {
            setQuiz({...quiz, subsetSize: numValue})
        } else if (!isNaN(numValue) && numValue > quiz.questions.length) {
            setQuiz({...quiz, subsetSize: quiz.questions.length})
            alert(`Το μέγιστο υποσύνολο ερωτήσεων είναι ${quiz.questions.length}.`);
        } else {
            setQuiz({...quiz, subsetSize: 0})
        }
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        const questionIDs = quiz.questions.map(q => q.ID);

        try {
            let apiUrl = `/quiz/${quiz.id}`

            const response = await api.put(apiUrl, {
                Title: quiz.title,
                Description: quiz.description,
                CourseID: parseInt(quiz.course.id),
                Visibility: quiz.isVisible,
                ShowSubset: quiz.isShowSubsetChecked,
                SubsetSize: quiz.subsetSize,
                QuestionsIDs: questionIDs
            });

            resetQuiz()
            window.location.href = `/courses/${quiz.course.ID}/quizzes/`;
        } catch (error) {
            console.error('Error update the quiz', error);
            // setError('Υπήρξε πρόβλημα κατά την δημιουργία του quiz. Παρακαλώ δοκιμάστε ξανά.');
        }
    };

    const handleDelete = async (event) => {
        event.preventDefault();

        resetQuiz()
        window.location.href = `/my-quizzes`;
    };

    return (
        <div>
            <Breadcrumb
                actualPath={`/courses/${quiz.course.ID}/quizzes/${quiz.id}/edit/step-three`}
                namePath={`/Μαθήματα/${quiz.course.Title}/Quiz/${quiz.title}/Επεξεργασία - Βήμα 3`}
            />
            <EditProgressBar/>
            <h2 className="createQuizStepFourPageTitle">3. Ολοκλήρωση</h2>
            <div className="createQuizStepFourDetailsRow">
                <div className="createQuizStepFourDetailsRowColumn">
                    <div> Αριθμός Ερωτήσεων: {quiz.questions.length}</div>
                    <div className="createQuizStepFourCheckboxRow">
                        <label>
                            Ορατό <input type="checkbox"
                                         checked={quiz.isVisible}
                                         onChange={(e) => setQuiz({...quiz, isVisible: e.target.checked})}/>
                        </label>

                    </div>
                    <div className="createQuizStepFourCheckboxRow">
                        <label>
                            Τυχαίο Υποσύνολο Ανά Συμπλήρωση
                            <input type="checkbox"
                                   checked={quiz.isShowSubsetChecked}
                                   onChange={(e) => setQuiz({...quiz, isShowSubsetChecked: e.target.checked})}/>
                        </label>
                    </div>

                    <div className={quiz.isShowSubsetChecked ? "" : "disabledInput"}>
                        <label className={quiz.isShowSubsetChecked ? "" : "disabledInput"}>
                            Πλήθος Ερωτήσεων Ανά Συμπλήρωση: </label>
                        <input type="number" value={quiz.subsetSize}
                               onChange={handleSubsetSizeChange}
                               disabled={!quiz.isShowSubsetChecked}
                               className="createQuizStepFourDetailsRowInputText"
                        />
                    </div>
                </div>
            </div>
            <div className="createQuizStepTwoButtonsContainer">
                <button className="createQuizStepFourButtonSubmit" onClick={handleSubmit}>Αποθήκευση</button>
                <button className="createQuizStepFourButtonDelete" onClick={handleDelete}>Διαγραφή</button>
            </div>
        </div>
    );
};

export default EditQuizStepThree;