import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCheck} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";

const cookies = new Cookies();

const MyQuizzesList = () => {
    const [quizzes, setQuizzes] = useState([]);
    const [quizSessions, setQuizSessions] = useState([]);
    const [finalQuizzes, setFinalQuizzes] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            const userCookie = cookies.get("user");
            const userType = userCookie.type;
            const specificID = userCookie.specificID;

            const apiUrlQuizzes = userType === "student"
                ? `${process.env.REACT_APP_API_URL}/student/${specificID}/quizzes`
                : `${process.env.REACT_APP_API_URL}/tutor/${specificID}/quizzes`;

            const apiUrlQuizSessions = `${process.env.REACT_APP_API_URL}/student/${specificID}/quiz_sessions`;

            try {
                const [quizzesResponse, quizSessionsResponse] = await Promise.all([
                    fetch(apiUrlQuizzes, {
                        method: 'GET',
                        headers: {
                            'Content-Type': 'application/json',
                            Authorization: `Bearer ${cookies.get("token")}`,
                        },
                        credentials: 'include',
                    }),
                    fetch(apiUrlQuizSessions, {
                        method: 'GET',
                        headers: {
                            'Content-Type': 'application/json',
                            Authorization: `Bearer ${cookies.get("token")}`,
                        },
                        credentials: 'include',
                    }),
                ]);

                const [quizzesResult, quizSessionsResult] = await Promise.all([
                    quizzesResponse.json(),
                    quizSessionsResponse.json(),
                ]);

                if (quizzesResponse.status === 500 || quizSessionsResponse.status === 500) {
                    throw Error(quizzesResult.message || quizSessionsResult.message);
                }

                if (quizzesResponse.status === 401 || quizSessionsResponse.status === 401) {
                    throw Error("unauthorized: 401");
                }

                if (quizzesResult.quizzes === undefined || quizSessionsResult.quizSessions === undefined) {
                    throw Error("error getting quizzes or quiz sessions");
                }

                setQuizzes(quizzesResult.quizzes);
                setQuizSessions(quizSessionsResult.quizSessions);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    const checkIfQuizHasBeenDone = () => {
        const updatedFinalQuizzes = quizzes.map(quiz => {
            let quizSession = quizSessions.find(session => session.quizName === quiz.Title);
            return {
                Title: quiz.Title,
                CourseName: quiz.CourseName,
                NumberOfQuestions: quiz.NumberOfQuestions,
                hasBeenDoneOnce: !!quizSession,
            };
        });

        setFinalQuizzes(updatedFinalQuizzes);
    };

    useEffect(() => {
        checkIfQuizHasBeenDone()
    }, [quizzes, quizSessions]);

    return (
        <React.Fragment>
            <ul className="quizzesList">
                <div className="quizzesListTitle">Διαθέσιμα Quiz</div>
                {finalQuizzes.map((item) => {
                    return (
                        <div className="singleQuizTextContainer">
                            <div className="singleQuizTitle">
                                {item.Title}
                                {item.hasBeenDoneOnce ? (
                                    <FontAwesomeIcon icon={faCheck} className="hasBeenDoneOnceCheckmark"/>
                                ) : (
                                    ""
                                )}
                            </div>
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