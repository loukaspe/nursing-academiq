import React from "react";
import "./Homepage.css";
import PageTitle from "../Utilities/PageTitle";
import LimitedCoursesList from "../CoursesList/LimitedCoursesList";
import LimitedQuizzesList from "../QuizzesList/LimitedQuizzesList";


const Homepage = (props) => {
    return (
        <>
            <div>
                <PageTitle title={"Αρχική Σελίδα"}/>
            </div>
            <div className="homepageContainer">
                <div className="coursesListContainer">
                    <LimitedCoursesList/>
                </div>
                <div className="quizListContainer">
                    <LimitedQuizzesList/>
                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;