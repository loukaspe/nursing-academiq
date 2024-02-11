import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBookmark} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";

const cookies = new Cookies();

const MyQuizzesList = () => {
    const [quizzes, setQuizzes] = useState([]);

    useEffect(() => {
        const fetchUserQuizzes = async () => {
            let userCookie = cookies.get("user");
            let userType = userCookie.type;
            let specificID = userCookie.specificID;

            let apiUrl = "";
            if (userType === "student") {
                apiUrl = process.env.REACT_APP_API_URL + `/student/${specificID}/quizzes`;
            } else if (userType === "tutor") {
                apiUrl = process.env.REACT_APP_API_URL + `/tutor/${specificID}/quizzes`;
            }


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

        fetchUserQuizzes();
    }, []);

    return (
        <React.Fragment>
            <ul className="quizzesList">
                <div className="quizzesListTitle">Διαθέσιμα Quiz</div>
                {quizzes.map((item) => {
                    return (
                        <div className="singleQuizTextContainer">
                            <div className="singleQuizTitle">{item.Title}</div>
                            <div className="singleQuizDetails">{item.CourseName}</div>
                            <div className="singleQuizDetails">{item.NumberOfQuestions} ερωτήσεις</div>
                        </div>
                    );
                })}
            </ul>
        </React.Fragment>
    );
};

export default MyQuizzesList;