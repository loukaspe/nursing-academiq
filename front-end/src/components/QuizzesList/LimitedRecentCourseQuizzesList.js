import React, {useState} from "react";
import "./QuizzesList.css";
import {Link} from "react-router-dom";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import axios from "axios";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const LimitedRecentCourseQuizzesList = (props) => {
    const [visibleQuizzes, setVisibleQuizzes] = useState(2);

    const token = cookies.get("access_token");

    const isTutorSignedIn = () => {
        return !!token;
    }

    const deleteQuiz = (id, title) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε το quiz ${title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/quiz/${id}`

            api.delete(apiUrl)
                .then(() => {
                    window.location.href = `/courses/${props.courseID}/quizzes`;
                })
                .catch(error => {
                    console.error('Error deleting quiz', error);
                });
        }
    };

    return (
        <React.Fragment>
            <ul className="recentQuizzesList">
                {
                    (isTutorSignedIn() && props.quizzes.length > 0) || props.quizzes.some(quiz => quiz.Visibility) ? (
                        props.quizzes.slice(0, visibleQuizzes).map((item) => {
                            return (
                                <div className="singleQuizContainer">
                                    <div className="quizContent">
                                        <div className="singleQuizTextContainer">
                                            <Link className="singleQuizTitle"
                                                  to={`/courses/${props.courseID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                            <div className="singleQuizDetails">{item.CourseName}</div>
                                            <div className="singleQuizDetails">
                                                {
                                                    (!isTutorSignedIn() && item.ShowSubset)
                                                    ? `${item.SubsetSize} ερωτήσεις`
                                                    : `${item.NumberOfQuestions} ερωτήσεις`
                                                }
                                            </div>
                                        </div>
                                    </div>
                                    {
                                        isTutorSignedIn() && <div className="quizIcons">
                                            <Link to={`/courses/${props.courseID}/quizzes/${item.ID}/edit`}>
                                                <FontAwesomeIcon icon={faPenToSquare} className="quizIcon"/>
                                            </Link>
                                            <FontAwesomeIcon icon={faTrashCan} className="quizIcon" onClick={() => {
                                                deleteQuiz(item.ID, item.Title)
                                            }}/>
                                        </div>
                                    }
                                </div>
                            );
                        })
                    ) : (
                        <div className="singleQuizTitle">Δεν υπάρχουν διαθέσιμα quiz.</div>
                    )}
                <div
                    className={`quizzesButtonContainer ${props.quizzes.length > visibleQuizzes ? 'multiple' : 'single'}`}>
                    {
                        isTutorSignedIn() &&
                        <Link className="myCoursesListButton" to={`/courses/${props.courseID}/quizzes/create`}>+ Νέο
                            Quiz</Link>
                    }
                    {
                        props.quizzes.length > visibleQuizzes &&
                        <Link className="moreButton" to={`/courses/${props.courseID}/quizzes`}>+ Όλα τα Quiz</Link>
                    }
                </div>
            </ul>
        </React.Fragment>
    );
};

export default LimitedRecentCourseQuizzesList;