import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import {Link} from "react-router-dom";

const LimitedRecentQuizzesList = () => {
    const [quizzes, setQuizzes] = useState([]);
    const [visibleQuizzes, setVisibleQuizzes] = useState(2);

    const fetchRecentQuizzes = async () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/quizzes/recent`;

        try {
            const response = await fetch(apiUrl, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                },
                credentials: 'include',
            })
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
        fetchRecentQuizzes();
    }, []);

    return (
        <React.Fragment>
            <ul className="limitedQuizzesList">
                <div className="quizzesListTitle">Πρόσφατα Quizzes</div>
                {quizzes.some(quiz => quiz.Visibility) ? (
                    quizzes.slice(0, visibleQuizzes).map((item) => {
                        return (
                            <div className="singleQuizContainer">
                                <div className="quizContent">
                                    <div className="singleQuizTextContainer">
                                        <Link className="singleQuizTitle"
                                              to={`/courses/${item.Course.ID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                        <div className="singleQuizDetails">{item.CourseName}</div>
                                        <div className="singleQuizDetails">
                                            {item.ShowSubset
                                                ? `${item.SubsetSize} ερωτήσεις`
                                                : `${item.NumberOfQuestions} ερωτήσεις`
                                            }
                                        </div>
                                    </div>
                                </div>
                            </div>
                        );
                    })
                ) : (
                    <div className="singleQuizTitle">Δεν υπάρχουν διαθέσιμα quiz.</div>
                )}
                <div className={`quizzesButtonContainer ${quizzes.length > visibleQuizzes ? 'multiple' : 'single'}`}>
                    {
                        quizzes.length > visibleQuizzes
                        &&
                        <Link className="myQuizzesListButton" to="/quizzes">+ Περισσότερα Quiz</Link>
                    }
                </div>
            </ul>
        </React.Fragment>
    );
};

export default LimitedRecentQuizzesList;