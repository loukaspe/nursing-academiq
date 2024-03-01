import React, {useEffect, useState} from "react";
import "./QuizHistoryList.css";
import Cookies from "universal-cookie";

const cookies = new Cookies();

const QuizHistoryList = () => {
    const [quizSessions, setQuizSessions] = useState([]);

    useEffect(() => {
        const fetchStudentQuizSessions = async () => {
            const userCookie = cookies.get("user");
            const specificID = userCookie.specificID;

            const apiUrlQuizSessions = `${process.env.REACT_APP_API_URL}/student/${specificID}/quiz_sessions`;

            try {
                const response = await fetch(apiUrlQuizSessions, {
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

                if (result.quizSessions === undefined) {
                    throw Error("error getting courses for student");
                }
                setQuizSessions(result.quizSessions);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchStudentQuizSessions();
    }, []);

    return (
        <React.Fragment>
            <ul className="quizHistoryList">
                <div className="quizSessionsListTitle">Ιστορικό Quiz</div>
                {quizSessions.map((item) => {
                    return (
                        <div className="singleQuizSessionTextContainer">
                            <div className="singleQuizSessionInfo">
                                <span className="singleQuizSessionTitle">{item.quizName}</span>
                                <span className="singleQuizSessionDate"> {item.date}</span>
                            </div>
                            <div className="singleQuizSessionResultsContainer">
                                <div className="singleQuizSessionResult">
                                    <span className="singleQuizSessionResultLabel">Σκορ:</span>{item.score}/{item.maxScore}
                                </div>
                                <div className="singleQuizSessionResult">
                                    <span className="singleQuizSessionResultLabel">Χρονος:</span> {item.duration}
                                </div>
                            </div>
                            <div className="singleQuizSessionButtonsContainer">
                                <button className="singleQuizSessionButton">Αποτελέσματα</button>
                                <button className="singleQuizSessionButton">Επανάληψη</button>
                            </div>
                        </div>
                    );
                })}
            </ul>
        </React.Fragment>
    );
};

export default QuizHistoryList;