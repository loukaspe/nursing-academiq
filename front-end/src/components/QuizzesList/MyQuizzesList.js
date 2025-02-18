import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";
import axios from "axios";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const MyQuizzesList = () => {
    const [quizzes, setQuizzes] = useState([]);

    const fetchUserQuizzes = async () => {
        let userCookie = cookies.get("user");
        let specificID = userCookie.specificID;

        let apiUrl = `/tutor/${specificID}/quizzes`;

        try {
            const response = await api.get(apiUrl);
            // TODO if 401 show unauthorized
            // TODO if 500 show server error
            if (response.status === 500) {
                throw Error(response.data.message);
            }

            if (response.status === 401) {
                throw Error("unauthorized: 401");
            }

            if (response.data.quizzes === undefined) {
                throw Error("error getting quizzes for tutor");
            }
            setQuizzes(response.data.quizzes);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    useEffect(() => {
        fetchUserQuizzes();
    }, []);

    const deleteQuiz = (id, title, courseID) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε το quiz ${title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/quiz/${id}`

            api.delete(apiUrl).then(
                () => {
                    fetchUserQuizzes()
                }
            ).catch(error => {
                console.error('Error deleting quiz', error);
            });
        }
    };

    return (
        <React.Fragment>
            <ul className="quizzesList">
                <div className="quizzesListTitle">Τα Quiz Μου</div>
                {quizzes.map((item) => {
                    return (
                        <div className="singleQuizContainer">
                            <div className="quizContent">
                                <div className="singleQuizTextContainer">
                                    <div className="singleQuizTitle">{item.Title}</div>
                                    <div className="singleQuizDetails">{item.CourseName}</div>
                                    <div className="singleQuizDetails">{item.NumberOfQuestions} ερωτήσεις</div>
                                </div>
                            </div>
                            <div className="quizIcons">
                                <Link to={`/courses/${item.CourseID}/quizzes/${item.ID}/edit`}>
                                    <FontAwesomeIcon icon={faPenToSquare} className="quizIcon"/>
                                </Link>
                                <FontAwesomeIcon icon={faTrashCan} className="quizIcon" onClick={() => {
                                    deleteQuiz(item.ID, item.Title, item.CourseID)
                                }}/>
                            </div>
                        </div>
                    );
                })}
                <div className={`quizzesButtonContainer single`}>
                    <Link className="myQuizzesListButton" to="/quizzes/create">+ Δημιουργία Quiz</Link>
                </div>
            </ul>
        </React.Fragment>
    );
};

export default MyQuizzesList;