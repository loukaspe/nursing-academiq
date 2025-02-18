import React, {useEffect, useState} from "react";
import "./MyCoursesList.css";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBookmark} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";

const LimitedRecentCoursesList = () => {
    const [courses, setCourses] = useState([]);
    const [visibleCourses, setVisibleCourses] = useState(2);

    useEffect(() => {
        const fetchRecentCourses = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/courses/recent`;

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
                    throw Error("error getting courses for tutor");
                }
                setCourses(result.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchRecentCourses();
    }, []);

    return (
        <React.Fragment>
            <ul className="limitedMyCoursesList">
                <div className="myCoursesListTitle">Πρόσφατα Μαθήματά</div>
                {courses.slice(0, visibleCourses).map((item) => {
                    return (
                        <div className="mySingleCourseContainer">
                            <FontAwesomeIcon icon={faBookmark} className="bookmarkIcon"/>
                            <div className="mySingleCourseTextContainer">
                                <Link className="mySingleCourseTitle" to={`/courses/${item.id}`}>{item.title}</Link>
                                <div className="mySingleCourseDetails">{item.description}</div>
                            </div>
                        </div>
                    );
                })}
                <div className={`coursesButtonContainer ${courses.length > visibleCourses ? 'multiple' : 'single'}`}>
                    {
                        courses.length > visibleCourses
                        &&
                        <Link className="myCoursesListButton" to="/courses">+ Περισσότερα Μαθήματα</Link>
                    }
                </div>
            </ul>
        </React.Fragment>
    );
};

export default LimitedRecentCoursesList;