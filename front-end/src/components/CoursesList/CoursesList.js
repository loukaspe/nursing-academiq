import React, {useEffect, useState} from "react";
import "./CoursesList.css";
import Cookies from "universal-cookie";

const cookies = new Cookies();

const CoursesList = () => {
    const [courses, setCourses] = useState([]);
    const [myCourses, setMyCourses] = useState([]);
    // const [selectedCourses, setSelectedCourses] = useState([]);
    const [isSubmitting, setSubmitting] = useState(false);

    const handleCheckboxChange = (id, isChecked) => {
        if (isChecked) {
            // setSelectedCourses((prevSelectedCourses) => {
            //     return [...prevSelectedCourses, id];
            // });
            setMyCourses((prevMyCourses) => {
                return [...prevMyCourses, id];
            });

        } else {
            setMyCourses((prevMyCourses) => {
                return prevMyCourses.filter((course) => course !== id);
            });
            // setSelectedCourses((prevMyCourses) => {
            //     return prevMyCourses.filter((course) => course !== id);
            // });
        }
    };

    const handleCourseRegistration = async () => {
        let userCookie = cookies.get("user");
        let specificID = userCookie.specificID;

        let apiUrl = process.env.REACT_APP_API_URL + `/student/${specificID}/courses`;

        let courses = myCourses;
        let requestBody = JSON.stringify({courses});
        console.log(requestBody);

        setSubmitting(true);
        try {
            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${cookies.get("token")}`,
                },
                credentials: 'include',
                body: requestBody,
            });

            if (response.ok) {
                console.log('Courses registered successfully!');
            } else {
                console.error('Failed to register courses:', response.statusText);
            }
        } catch (error) {
            console.error('Error registering courses:', error);
        } finally {
            setSubmitting(false);
        }
    };

    useEffect(() => {
        const fetchCourses = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/courses`;

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

        // This page is only for students
        const fetchStudentCourses = async () => {
            let userCookie = cookies.get("user");
            let specificID = userCookie.specificID;

            let apiUrl = process.env.REACT_APP_API_URL + `/student/${specificID}/courses`;


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

                const coursesIDs = result.courses.map(course => course.id);

                setMyCourses(coursesIDs);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchCourses();
        fetchStudentCourses()
    }, []);

    return (
        <React.Fragment>
            <ul className="coursesList">
                <div className="coursesListTitle">Κατάλογος Μαθημάτων</div>
                <div className="headerContainer">
                    <div className="singleCourseTextContainer">
                    </div>
                    <div className="registerTitle">
                        <span>Εγγραφή</span>
                    </div>
                    <div style={{clear: 'both'}}></div>
                </div>
                {courses.map((item) => {
                    return (
                        <div className="singleCourseContainer">
                            <div className="singleCourseTextContainer">
                                <span className="singleCourseTitle">{item.title}</span>
                                <div className="singleCourseDetails">{item.description}</div>
                            </div>
                            <div className="singleCourseCheckbox">
                                <input
                                    type="checkbox"
                                    id={`checkbox-${item.id}`}
                                    checked={myCourses.includes(item.id)}
                                    // checked={selectedCourses.includes(item.id) || myCourses.includes(item.id)}
                                    onChange={(event) => handleCheckboxChange(item.id, event.target.checked)}
                                />
                            </div>
                            <div style={{clear: 'both'}}></div>
                        </div>
                    );
                })}
                <button
                    className="registerButton"
                    onClick={handleCourseRegistration}
                    style={{
                        backgroundColor: isSubmitting ? "#C3C3C3" : "#220D6A",
                    }}
                    disabled={isSubmitting}
                >
                    {isSubmitting ? 'Υποβολή...' : 'Υποβολή'}
                </button>
            </ul>
        </React.Fragment>
    );
};

export default CoursesList;