import React, {useEffect, useState} from "react";
import "./CoursesList.css";
import Cookies from "universal-cookie";
import {Link} from "react-router-dom";
import Breadcrumb from "../Utilities/Breadcrumb";

const cookies = new Cookies();

const CoursesList = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const fetchCourses = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/courses`;

            try {
                const response = await fetch(apiUrl, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                    credentials: 'include',
                });
                const result = await response.json();
                // TODO if 401 show unauthorized
                // TODO if 500 show server error
                if (response.status === 500) {
                    throw Error(result.message);
                }

                if (response.status === 401) {
                    throw Error("unauthorized: 401");
                }

                if (result.courses === undefined) {
                    throw Error("error getting courses for student");
                }
                setCourses(result.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchCourses();
    }, []);

    return (
        <React.Fragment>
            <Breadcrumb actualPath="/courses" namePath="Μαθήματα"/>
            <ul className="coursesList">
                <div className="coursesListTitle">Κατάλογος Μαθημάτων</div>
                <div className="headerContainer">
                    <div className="singleCourseTextContainer">
                    </div>
                    <div style={{clear: 'both'}}></div>
                </div>
                {courses.map((item) => {
                    return (
                        <div className="singleCourseContainer">
                            <div className="singleCourseTextContainer">
                                <Link className="singleCourseTitle" to={`/courses/${item.id}`}>{item.title}</Link>
                                <div className="singleCourseDetails">{item.description}</div>
                            </div>
                            <div style={{clear: 'both'}}></div>
                        </div>
                    );
                })}
            </ul>
        </React.Fragment>
    );
};

export default CoursesList;