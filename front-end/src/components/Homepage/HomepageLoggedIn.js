import React from "react";
import "./Homepage.css";
import PageTitle from "../Utilities/PageTitle";
import LimitedMyCoursesList from "../CoursesList/LimitedMyCoursesList";
import LimitedMyQuizzesList from "../QuizzesList/LimitedMyQuizzesList";

const Homepage = () => {
    return (
        <>
            <div>
                <PageTitle title={"Αρχική Σελίδα"}/>
            </div>
            <div className="homepageContainer">
                <div className="coursesListContainer">
                    <LimitedMyCoursesList/>
                </div>
                <div className="quizListContainer">
                    <LimitedMyQuizzesList/>
                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;