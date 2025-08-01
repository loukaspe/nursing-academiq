import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import {Link} from "react-router-dom";
import Breadcrumb from "../Utilities/Breadcrumb";

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
                throw Error("error getting quizzes for tutor");
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
            <Breadcrumb actualPath={`/quizzes`}
                        namePath={`/Quizzes`}/>
            <ul className="quizzesList">
                <div className="quizzesListTitle">Quizzes</div>
                {quizzes.length > 0 ? (
                    quizzes.map((item) => {
                        return (
                            <div className="singleQuizContainer">
                                <div className="quizContent">
                                    <div className="singleQuizTextContainer">
                                        <Link className="singleQuizTitle"
                                              to={`/courses/${item.Course.ID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                        <div className="singleQuizDetails">{item.Course.Title}</div>
                                        <div className="singleQuizDetails">{item.Description}</div>
                                        <div className="singleQuizDetails">{item.ShowSubset
                                            ? `${item.SubsetSize} ερωτήσεις`
                                            : `${item.NumberOfQuestions} ερωτήσεις`
                                        }</div>
                                    </div>
                                </div>
                            </div>
                        );
                    })
                ) : (
                    <div className="singleQuizTitle">Δεν υπάρχουν διαθέσιμα quiz.</div>
                )}
            </ul>
        </React.Fragment>
    );
};

export default QuizzesList;