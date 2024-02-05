import React from "react";
import "./Homepage.css";
import PageTitle from "../Utilities/PageTitle";
import CoursesList from "../CoursesList/CoursesList";
import QuizzesList from "../QuizzesList/QuizzesList";


const Homepage = (props) => {
    return (
        <>
            <div>
                <PageTitle title={"Αρχική Σελίδα"}/>
            </div>
            <div className="homepageContainer">
                <div className="coursesListContainer">
                    <CoursesList/>
                </div>
                <div className="quizListContainer">
                    <QuizzesList/>
                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;