import React, {useEffect, useState} from "react";
import "./QuizzesList.css";
import {Link, useLocation, useNavigate} from "react-router-dom";
import Breadcrumb from "../Utilities/Breadcrumb";

const QuizzesListSearch = () => {
    const location = useLocation();
    const quizzes = location.state?.quizzes || [];

    return (
        <React.Fragment>
            <Breadcrumb actualPath={`/quizzes`}
                        namePath={`/Quizzes`}/>
            <ul className="quizzesList">
                <div className="quizzesListTitle">Τα Quiz Μου</div>
                {quizzes.length > 0 ? (
                    quizzes.map((item) => {
                        return (
                            <div className="singleQuizContainer">
                                <div className="quizContent">
                                    <div className="singleQuizTextContainer">
                                        <Link className="singleQuizTitle"
                                              to={`/courses/${item.Course.ID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                        <div className="singleQuizDetails">{item.CourseName}</div>
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

export default QuizzesListSearch;