import React, {useEffect, useState} from "react";
import "./MyCoursesList.css";
import Cookies from "universal-cookie";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBookmark} from "@fortawesome/free-solid-svg-icons";
import {Link} from "react-router-dom";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const LimitedMyCoursesList = () => {
    const [courses, setCourses] = useState([]);
    const [visibleCourses, setVisibleCourses] = useState(2);

    useEffect(() => {
        const fetchUserCourses = async () => {
            let userCookie = cookies.get("user");
            let specificID = userCookie.specificID;

            let apiUrl =  `/tutor/${specificID}/courses`;

            try {
                const response = await api.get(apiUrl);
                // TODO if 401 show unauthorized
                // TODO if 500 show server error
                if (response.status === 500) {
                    throw Error(response.data.message);
                }

                if (response.status === 401) {
                    throw Error("unauthorized: 401");
                }

                if (response.data.courses === undefined) {
                    throw Error("error getting courses for tutor");
                }
                setCourses(response.data.courses);
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
                {courses.length > 0 ? (
                    courses.slice(0, visibleCourses).map((item) => {
                        return (
                            <div className="mySingleCourseContainer">
                                <FontAwesomeIcon icon={faBookmark} className="bookmarkIcon"/>
                                <div className="mySingleCourseTextContainer">
                                    <Link className="mySingleCourseTitle" to={`/courses/${item.id}`}>{item.title}</Link>
                                    <div className="mySingleCourseDetails">{item.description}</div>
                                </div>
                            </div>
                        );
                    })
                ) : (
                    <div className="mySingleCourseTitle">Δεν υπάρχουν διαθέσιμα μαθήματα.</div>
                )}
                <div className={`coursesButtonContainer ${courses.length > visibleCourses ? 'multiple' : 'single'}`}>
                    <Link className="myCoursesListButton" to="/courses/create">+ Δημιουργία Μαθήματος</Link>
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