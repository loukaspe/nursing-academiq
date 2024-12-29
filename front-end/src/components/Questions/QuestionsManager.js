import React, {useEffect, useState} from "react";
import "./QuestionsManager.css";
import {Link} from "react-router-dom";
import Breadcrumb from "../Utilities/Breadcrumb";

const QuestionsManager = () => {
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
            <ul className="questionsManagerCoursesList">
                <div className="questionsManagerCoursesListTitle">Διαχείριση Ερωτήσεων</div>
                <div className="questionsManagerCoursesListSubTitle">Επιλογή Μαθήματος</div>
                {courses.map((item) => {
                    return (
                        <div className="questionsManagerSingleCourseContainer">
                            <div className="questionsManagerSingleCourseTextContainer">
                                <Link className="questionsManagerSingleCourseTitle" to={`/courses/${item.id}/questions/manage`}>{item.title} - {item.numberOfQuestions} Ερωτήσεις</Link>
                                <div className="questionsManagerSingleCourseDetails">{item.description}</div>
                            </div>
                            <div style={{clear: 'both'}}></div>
                        </div>
                    );
                })}
            </ul>
        </React.Fragment>
    );
};

export default QuestionsManager;