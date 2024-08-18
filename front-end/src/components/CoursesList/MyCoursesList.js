import React, {useEffect, useState} from "react";
import "./MyCoursesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBookmark} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";

const cookies = new Cookies();

const MyCoursesList = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const fetchTutorCourses = async () => {
            let userCookie = cookies.get("user");
            let specificID = userCookie.specificID;

            console.log(specificID);

            let apiUrl = process.env.REACT_APP_API_URL + `/tutor/${specificID}/courses`;

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

        fetchTutorCourses();
    }, []);

    return (
        <React.Fragment>
            <div className="myCoursesListTitle">Διαχείριση Μαθημάτων</div>
            <ul className="myCoursesList">
                <div className="myCoursesListSubTitle">Επιλογή Μαθήματος</div>
                {courses.map((item) => {
                    return (
                        <div className="mySingleCourseContainer">
                            {/*<FontAwesomeIcon icon={faBookmark} className="bookmarkIcon"/>*/}
                            <div className="mySingleCourseTextContainer">
                                <span className="mySingleCourseTitle">{item.title}</span>
                                <div className="mySingleCourseDetails">{item.description}</div>
                            </div>
                        </div>
                    );
                })}
                <Link className="registerButton" to="/courses/create">+ Δημιουργία Μαθήματος</Link>
            </ul>
        </React.Fragment>
    );
};

export default MyCoursesList;