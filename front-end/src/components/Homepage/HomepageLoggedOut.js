import React, {useEffect, useState} from "react";
import "./Homepage.css";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare} from "@fortawesome/free-solid-svg-icons";
import {Link, useNavigate} from "react-router-dom";
import apiWithAPIKey from "../Utilities/APICallerAPIKey";
import api from "../Utilities/APICaller";

const Homepage = () => {
    const [courses, setCourses] = useState([]);
    const [selectedCourse, setSelectedCourse] = useState("Όλα τα Μαθήματα");
    const [searchText, setSearchText] = useState("");
    const [error, setError] = useState('');

    const navigate = useNavigate();

    const handleSearchSubmit = async () => {
        try {
            let apiUrl = `/quizzes/search`

            const response = await apiWithAPIKey.post(apiUrl, {
                title: searchText,
                courseName: selectedCourse === "Όλα τα Μαθήματα" ? "" : selectedCourse,
            });

            if (response.status === 200 && response.data) {
                console.log(response.data)
                navigate("/quizzes/search", {state: {quizzes: response.data.quizzes}});
            } else if (response.status === 204) {
                setError('Δεν βρέθηκαν αποτελέσματα.');
            }
        } catch (error) {
            setError('Υπήρξε πρόβλημα κατά την αναζήτηση Quiz. Παρακαλώ δοκιμάστε ξανά.');
            console.error('Error searching the quiz', error);
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
                    throw Error("error getting courses list");
                }
                setCourses(result.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchCourses();
    }, []);

    const filteredCourses =
        selectedCourse === "Όλα τα Μαθήματα"
            ? courses
            : courses.filter((course) => course.title === selectedCourse);


    return (
        <>
            <div className="homepageContainer">
                <div className="homepageContainer">
                    {/* Search Section */}
                    <section className="homepageSearchSection">
                        <h1 className="homepageSectionTitle">Αναζήτηση Κουίζ</h1>
                        <div className="homepageSearchControls">
                            <input
                                type="text"
                                placeholder="Αναζήτηση Κουίζ..."
                                className="homepageSearchInput"
                                value={searchText}
                                onChange={(e) => setSearchText(e.target.value)}
                            />
                            <select
                                className="homepageFilterSelect"
                                value={selectedCourse}
                                onChange={(e) => setSelectedCourse(e.target.value)}
                            >
                                <option>Όλα τα Μαθήματα</option>
                                {courses.map((course) => (
                                    <option key={course.id} value={course.title}>
                                        {course.title}
                                    </option>
                                ))}
                            </select>
                            <button className="homepageSearchButton" onClick={handleSearchSubmit}>Αναζήτηση</button>
                        </div>
                    </section>
                    {error && <div className="homepageInfoRow">{error}</div>}

                    {/* Courses Section */}
                    {courses &&
                        <section className="homepageCoursesSection">
                            <h2 className="homepageSectionTitle">Περιήγηση Μαθημάτων</h2>
                            <div className="homepageCoursesGrid">
                                {courses.slice(0, 3).map((course) => (
                                    <div className="homepageCoursesCard">
                                        <img src="/images/courseThumbnail.png" alt="Course Image"
                                             className="homepageCourseImage"/>
                                        <div className="homepageCourseName">{course.title}</div>
                                    </div>
                                ))}

                            </div>
                            <div className="homepageSeeMoreContainer">
                                <Link className="homepageSeeMoreButton" to={`/courses`}>
                                    Δείτε Περισσότερα
                                </Link>
                            </div>
                        </section>
                    }

                </div>
            </div>
            <div style={{clear: 'both'}}></div>
        </>
    );
};

export default Homepage;