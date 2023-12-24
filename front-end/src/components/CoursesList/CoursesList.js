import React from "react";
import "./CoursesList.css";

const CoursesList = ({courses}) => {
    return (
        <React.Fragment>
            <ul className="coursesList">
                {(courses).map((item) => {
                    return (
                        <div>
                            <li className="singleCourse" key={item.id}>
                                <div className="singleCourseTitle">{item.title}</div>
                                <div className="singleCourseDetails">{item.details}</div>
                            </li>
                            <hr/>
                        </div>
                    );
                })}
            </ul>
        </React.Fragment>
    );
};

export default CoursesList;