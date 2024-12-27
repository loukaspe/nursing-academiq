import React from "react";
import "./Homepage.css";
import PageTitle from "../Utilities/PageTitle";
import LimitedRecentCourseQuizzesList from "../QuizzesList/LimitedRecentCourseQuizzesList";
import LimitedRecentQuizzesList from "../QuizzesList/LimitedRecentQuizzesList";
import LimitedRecentCoursesList from "../CoursesList/LimitedRecentCoursesList";

const Homepage = () => {
    return (
        <>
            <div>
                <PageTitle title={"Αρχική Σελίδα"}/>
            </div>
            <div className="homepageContainer">
                <div className="coursesListContainer">
                    <LimitedRecentCoursesList/>
                </div>
                <div className="quizListContainer">
                    <LimitedRecentQuizzesList/>
                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;