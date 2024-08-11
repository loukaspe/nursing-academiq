import React, {useEffect, useState} from "react";
import "./MyCoursesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBookmark} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";

const cookies = new Cookies();

const LimitedMyCoursesList = () => {
    const [courses, setCourses] = useState([]);
    const [visibleCourses, setVisibleCourses] = useState(2);

    useEffect(() => {
        const fetchUserCourses = async () => {
            let userCookie = cookies.get("user");
            let userType = userCookie.type;
            let specificID = userCookie.specificID;

            let apiUrl = "";
            if (userType === "student") {
                apiUrl = process.env.REACT_APP_API_URL + `/student/${specificID}/courses`;
            } else if (userType === "tutor") {
                apiUrl = process.env.REACT_APP_API_URL + `/tutor/${specificID}/courses`;
            }


            try {
                const response = await fetch(apiUrl, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${cookies.get("token")}`,
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

        fetchUserCourses();
    }, []);

    return (
        <React.Fragment>
            <ul className="limitedMyCoursesList">
                <div className="myCoursesListTitle">Τα Μαθήματά Μου</div>
                {courses.slice(0, visibleCourses).map((item) => {
                    return (
                        <div className="mySingleCourseContainer">
                            <FontAwesomeIcon icon={faBookmark} className="bookmarkIcon"/>
                            <div className="mySingleCourseTextContainer">
                                <span className="mySingleCourseTitle">{item.title}</span>
                                <div className="mySingleCourseDetails">{item.description}</div>
                            </div>
                        </div>
                    );
                })}
                <div className={`coursesButtonContainer ${courses.length > visibleCourses ? 'multiple' : 'single'}`}>
                    <Link className="myCoursesListButton" to="/my-courses">+ Δημιουργία Μαθήματος</Link>
                    {
                        courses.length > visibleCourses
                        &&
                        <Link className="myCoursesListButton" to="/my-courses">+ Περισσότερα Μαθήματα</Link>
                    }
                </div>
            </ul>
        </React.Fragment>
    );
};

export default LimitedMyCoursesList;