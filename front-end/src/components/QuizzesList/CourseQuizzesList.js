import React, {useEffect, useState} from "react";
import "./CourseQuizzesList.css";
import Cookies from "universal-cookie";
import {Link, useParams, useNavigate} from "react-router-dom";
import axios from "axios";
import SectionTitle from "../Utilities/SectionTitle";
import {useHistory} from "react-router-dom";

const cookies = new Cookies();

const CourseQuizzesList = (props) => {
    const [quizzes, setQuizzes] = useState([]);
    const [course, setCourse] = useState({});

    const params = useParams();
    let courseID = params.id;

    let navigate = useNavigate();

    useEffect(() => {
        fetchCourse();
    }, []);

    const fetchCourse = () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}/extended`

        axios.get(apiUrl, {
            headers: {
                'Authorization': `Bearer ${cookies.get("token")}`,
            },
        })
            .then(response => {
                if (response.data) {
                    setCourse(response.data);
                }

                if (response.data.quizzes) {
                    setQuizzes(response.data.quizzes);
                }
            })
            .catch(error => {
                console.error('Error fetching course quizzes data', error);
            });
    };

    return (
        <React.Fragment>
            <div className="singleCourseQuizzesPageHeader">
                <div className="singleCourseQuizzesPageInfo">
                    <span className="singleCourseQuizzesPageCourseName">{course.title}</span>
                </div>
                <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
            </div>
            <div className="singleCourseDescription">
                <div>{course.description}</div>
            </div>
            <div className="singleCourseQuizzes">
                <div className="singleCoursePageSectionTitle">
                    <SectionTitle title="Quiz Μαθήματος"/>
                </div>
                <ul className="courseQuizzesList">
                    {quizzes.map((item) => {
                        return (
                            <div className="singleQuizTextContainer">
                                <div className="singleQuizTitle">{item.Title}</div>
                                <div className="singleQuizDetails">{course.title}</div>
                                <div className="singleQuizDetails">{item.NumberOfQuestions} ερωτήσεις</div>
                            </div>
                        );
                    })}
                </ul>
            </div>
        </React.Fragment>

    );
};

export default CourseQuizzesList;