import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";
import axios from "axios";

const QuizzesList = () => {
    const [quizzes, setQuizzes] = useState([]);

    const fetchQuizzes = async () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/quizzes`;

        try {
            const response = await fetch(apiUrl, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
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
        fetchQuizzes();
    }, []);

    return (
        <React.Fragment>
            <ul className="quizzesList">
                <div className="quizzesListTitle">Τα Quiz Μου</div>
                {quizzes.map((item) => {
                    return (
                        <div className="singleQuizContainer">
                            <div className="quizContent">
                                <div className="singleQuizTextContainer">
                                    <Link className="singleQuizTitle"
                                          to={`/courses/${item.Course.ID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                    <div className="singleQuizDetails">{item.CourseName}</div>
                                    <div className="singleQuizDetails">{item.NumberOfQuestions} ερωτήσεις</div>
                                </div>
                            </div>
                        </div>
                    );
                })}
            </ul>
        </React.Fragment>
    );
};

export default QuizzesList;