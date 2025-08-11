import React, {useEffect, useState} from "react";
import "./SingleCourse.css";
import Cookies from "universal-cookie";

import {Link, useParams} from "react-router-dom";
import SectionTitle from "../Utilities/SectionTitle";
import LimitedRecentCourseQuizzesList from "../QuizzesList/LimitedRecentCourseQuizzesList";
import axios from "axios";
import LimitedRecentCourseChaptersList from "../ChaptersList/LimitedCourseChaptersList";
import Breadcrumb from "../Utilities/Breadcrumb";
import {faPenToSquare} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const SingleCourse = () => {
    const [course, setCourse] = useState({});
    const [quizzes, setQuizzes] = useState([]);
    const [chapters, setChapters] = useState([]);

    const token = cookies.get("access_token");

    const isTutorSignedIn = () => {
        return !!token;
    }

    const params = useParams();
    let courseID = params.id;

    useEffect(() => {
        fetchCourse();
    }, [courseID]);

    const fetchCourse = () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}/extended`

        axios.get(apiUrl, {
            headers: {
                Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
            },
        })
            .then(response => {
                console.log(response.data);
                if (response.data) {
                    setCourse(response.data);
                }

                if (response.data.quizzes) {
                    setQuizzes(response.data.quizzes);
                }

                if (response.data.chapters) {
                    setChapters(response.data.chapters);
                }
            })
            .catch(error => {
                console.error('Error fetching course data', error);
            });
    };

    const deleteCourse = () => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε το μάθημα ${course.title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/course/${courseID}`

            api.delete(apiUrl)
                .then(() => {
                    window.location.href = "/courses";
                })
                .catch(error => {
                    console.error('Error deleting course', error);
                });
        }
    };

    return (
        <React.Fragment>
            <Breadcrumb actualPath={`/courses/${courseID}`} namePath={`/Μαθήματα/${course.title}`}/>
            <div className="singleCoursePageHeader">
                <div className="singleCoursePageInfo">
                    <span className="singleCoursePageCourseName">{course.title}</span>
                    {
                        isTutorSignedIn() ? (
                            <Link className="link" to={`/courses/${courseID}/edit`}>
                                <FontAwesomeIcon icon={faPenToSquare} className="courseIcon"/>
                            </Link>
                        ) : (
                            <span className="singleCoursePageTeacherName">Υπεύθυνος/η: {course.tutorName}</span>
                        )
                    }
                </div>
                {
                    isTutorSignedIn()
                    &&
                    <>
                        <Link className="courseButton" to={`/courses/${courseID}/questions/manage`}>
                            Διαχείριση Ερωτήσεων
                        </Link>
                        <button className="courseDangerButton" onClick={() => {
                            deleteCourse()
                        }}>Διαγραφή
                        </button>
                    </>
                }
            </div>
            <div className="singleCourseDescription">
                <div>{course.description}</div>
                {
                    isTutorSignedIn() &&
                    <Link className="link" to={`/courses/${courseID}/edit`}>
                        <FontAwesomeIcon icon={faPenToSquare} className="courseIcon"/>
                    </Link>
                }
            </div>
            <div className="singleCoursePageContainer">
                <div className="singleCoursePageDetails">
                    <div className="singleCourseChapters">
                        <div className="singleCoursePageSectionTitle">
                            <SectionTitle title="Θεματικές Ενότητες"/>
                        </div>
                        <LimitedRecentCourseChaptersList chapters={chapters} courseID={courseID}/>
                    </div>
                    <div className="singleCourseQuizzes">
                        <div className="singleCoursePageSectionTitle">
                            <SectionTitle title="Quizzes"/>
                        </div>
                        <LimitedRecentCourseQuizzesList quizzes={quizzes} courseID={courseID}/>
                    </div>
                </div>
            </div>
        </React.Fragment>
    )
        ;
};

export default SingleCourse;