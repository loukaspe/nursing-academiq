import React, {useEffect, useState} from "react";
import "./MyCoursesList.css";
import Cookies from "universal-cookie";
import {Link} from "react-router-dom";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const MyCoursesList = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const fetchTutorCourses = async () => {
            let userCookie = cookies.get("user");
            let specificID = userCookie.specificID;
            
            let apiUrl = `/tutor/${specificID}/courses`;

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

        fetchTutorCourses();
    }, []);

    return (
        <React.Fragment>
            <div className="myCoursesListTitle">Διαχείριση Μαθημάτων</div>
            <ul className="myCoursesList">
                <div className="myCoursesListSubTitle">Επιλογή Μαθήματος</div>
                {courses.length > 0 ? (
                    courses.map((item) => {
                        return (
                            <div className="mySingleCourseContainer">
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
                <Link className="registerButton" to="/courses/create">+ Δημιουργία Μαθήματος</Link>
            </ul>
        </React.Fragment>
    );
};

export default MyCoursesList;