import React, {useEffect, useState} from "react";
import "./SingleCourse.css";
import Cookies from "universal-cookie";
import {Link, useLocation} from "react-router-dom";

import {useParams} from "react-router-dom";
import PageTitle from "../Utilities/PageTitle";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faUser} from "@fortawesome/free-solid-svg-icons";
import SectionTitle from "../Utilities/SectionTitle";
import LimitedRecentCourseQuizzesList from "../QuizzesList/LimitedRecentCourseQuizzesList";
import axios from "axios";
import LimitedRecentCourseChaptersList from "../ChaptersList/LimitedCourseChaptersList";

const cookies = new Cookies();

const SingleCourse = () => {
    const [course, setCourse] = useState({});
    const [quizzes, setQuizzes] = useState([]);
    const [chapters, setChapters] = useState([]);

    const params = useParams();
    let courseID = params.id;

    useEffect(() => {
        fetchCourse();
    }, []);

    const fetchCourse = () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}/extended`

        axios.get(apiUrl, {
            headers: {
                Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
            },
        })
            .then(response => {
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

    return (
        <React.Fragment>
            <div className="singleCoursePageHeader">
                <div className="singleCoursePageInfo">
                    <span className="singleCoursePageCourseName">{course.title}</span>
                    <span className="singleCoursePageTeacherName">{course.tutorName}</span>
                </div>
                <Link className="unregisterButton" to="/change-password">Απεγγραφή</Link>
            </div>
            <div className="singleCourseDescription">
                <div>{course.description}</div>
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
                            <SectionTitle title="Πρόσφατα Quizs"/>
                        </div>
                        <LimitedRecentCourseQuizzesList quizzes={quizzes} courseID={courseID}/>
                    </div>
                </div>
            </div>
        </React.Fragment>
    );
};

export default SingleCourse;