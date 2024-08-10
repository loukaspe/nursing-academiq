import React, {useState} from "react";
import "./QuizzesList.css";
import {Link} from "react-router-dom";

const LimitedRecentCourseQuizzesList = (props) => {
    const [visibleQuizzes, setVisibleQuizzes] = useState(2);

    return (
        <React.Fragment>
            <ul className="recentQuizzesList">
                {props.quizzes.slice(0, visibleQuizzes).map((item) => {
                    return (
                        <div className="singleQuizTextContainer">
                            <Link className="singleQuizTitle"
                                  to={`/courses/${props.courseID}/quizzes/${item.ID}`}>{item.Title}</Link>
                            <div className="singleQuizDetails">{item.CourseName}</div>
                            <div className="singleQuizDetails">{item.NumberOfQuestions} ερωτήσεις</div>
                        </div>
                    );
                })}
                {
                    props.quizzes.length > visibleQuizzes &&
                    (
                        <Link className="moreButton" to={`/courses/${props.courseID}/quizzes`}>+ Όλα τα Quiz</Link>
                    )
                }
            </ul>
        </React.Fragment>
    );
};

export default LimitedRecentCourseQuizzesList;