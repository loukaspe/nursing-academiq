import React from "react";
import "./Homepage.css";
import Logo from "../Logo/Logo";
import PageTitle from "../Utilities/PageTitle";
import CoursesList from "../CoursesList/CoursesList";
import LoginForm from "../Login/LoginForm";


const Homepage = (props) => {
    return (
        <>
            <div>
                <PageTitle title={"Αρχική Σελίδα"}/>
            </div>
            <div className="homepageContainer">
                <div className="coursesListContainer">
                    <CoursesList coursesList={props.coursesList}/>
                </div>
                <div className="quizListContainer">
                    <CoursesList coursesList={props.coursesList}/>
                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;