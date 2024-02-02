import React, {useEffect, useState} from "react";
import "./CoursesList.css";
import Cookies from "universal-cookie";

const cookies = new Cookies();

const CoursesList = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const fetchStudentCourses = async () => {
            let userCookie = cookies.get("user");
            let userType = userCookie.type;
            if (userType !== "student") {
                throw Error("Not a student");
            }
            let studentID = userCookie.specificID;

            const apiUrl = process.env.REACT_APP_API_URL + `/student/${studentID}/courses`;

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

        fetchStudentCourses();
    }, []);

    return (
        <React.Fragment>
            <ul className="coursesList">
                {(courses).map((item) => {
                    return (
                        <div>
                            <li className="singleCourse" key={item.id}>
                                <div className="singleCourseTitle">{item.title}</div>
                                <div className="singleCourseDetails">{item.description}</div>
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