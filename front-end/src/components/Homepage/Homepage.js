import React from "react";
import "./Homepage.css";
import PageTitle from "../Utilities/PageTitle";


const Homepage = () => {
    return (
        <>
            <div>
                <PageTitle title={"Αρχική Σελίδα"}/>
            </div>
            <div className="homepageContainer">
                <div className="coursesListContainer">
                    {/*<LimitedMyCoursesList/>*/}
                    TI NA DEIKSW 1
                </div>
                <div className="quizListContainer">
                    {/*<LimitedMyQuizzesList/>*/}
                    TI NA DEIKSW 2
                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;