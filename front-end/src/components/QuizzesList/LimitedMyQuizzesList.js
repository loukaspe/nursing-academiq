import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";
import axios from "axios";

const cookies = new Cookies();

const LimitedMyQuizzesList = () => {
    const [quizzes, setQuizzes] = useState([]);
    const [visibleQuizzes, setVisibleQuizzes] = useState(2);

    const fetchUserQuizzes = async () => {
        let userCookie = cookies.get("user");
        let specificID = userCookie.specificID;

        let apiUrl = process.env.REACT_APP_API_URL + `/tutor/${specificID}/quizzes`;

        try {
            const response = await fetch(apiUrl, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${cookies.get("token")}`,
                },
                credentials: 'include',
            });
            const result = await response.json();
            // TODO if 401 show unauthorized
            // TODO if 500 show server error
            if (response.status === 500) {
                throw Error(result.message);
            }

            if (response.status === 401) {
                throw Error("unauthorized: 401");
            }

            if (result.quizzes === undefined) {
                throw Error("error getting quizzes for student");
            }
            setQuizzes(result.quizzes);
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
            let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${id}`

            axios.delete(apiUrl, {
                headers: {
                    Authorization: `Bearer ${cookies.get("token")}`,
                },
            }).then(
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
            <ul className="limitedQuizzesList">
                <div className="quizzesListTitle">Τα Quiz Μου</div>
                {quizzes.slice(0, visibleQuizzes).map((item) => {
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
                <div className={`quizzesButtonContainer ${quizzes.length > visibleQuizzes ? 'multiple' : 'single'}`}>
                    <Link className="myQuizzesListButton" to="/quizzes/create">+ Δημιουργία Quiz</Link>
                    {
                        quizzes.length > visibleQuizzes
                        &&
                        <Link className="myQuizzesListButton" to="/my-quizzes">+ Περισσότερα Quiz</Link>
                    }
                </div>
            </ul>
        </React.Fragment>
    );
};

export default LimitedMyQuizzesList;