import React, {useEffect, useState} from "react";
import "./CourseQuizzesList.css";
import Cookies from "universal-cookie";
import {Link, useParams, useNavigate} from "react-router-dom";
import axios from "axios";
import SectionTitle from "../Utilities/SectionTitle";
import Breadcrumb from "../Utilities/Breadcrumb";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const CourseQuizzesList = (props) => {
    const [quizzes, setQuizzes] = useState([]);
    const [course, setCourse] = useState({});

    const params = useParams();
    let courseID = params.id;

    let navigate = useNavigate();

    const token = cookies.get("access_token");

    const isTutorSignedIn = () => {
        return !!token;
    }

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
            })
            .catch(error => {
                console.error('Error fetching course quizzes data', error);
            });
    };

    const deleteQuiz = (id, title) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε το quiz ${title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/quiz/${id}`

            api.delete(apiUrl)
                .then(() => {
                    window.location.href = `/courses/${courseID}/quizzes`;
                })
                .catch(error => {
                    console.error('Error deleting quiz', error);
                });
        }
    };

    return (
        <React.Fragment>
            <Breadcrumb
                actualPath={`/courses/${courseID}/quizzes/`}
                namePath={`/Μαθήματα/${course.title}/Quizzes`}
            />
            <div className="singleCourseQuizzesPageHeader">
                <div className="singleCourseQuizzesPageInfo">
                    <span className="singleCourseQuizzesPageCourseName">{course.title}</span>
                    
                </div>
                {
                    isTutorSignedIn()
                    &&
                    <>
                        <Link className="courseButton" to={`/courses/${courseID}/quizzes/create`}>
                            + Προσθήκη Quiz
                        </Link>
                    </>
                }
            </div>
            <div className="courseQuizzesPageCourseDescription">
                <div>{course.description}</div>
            </div>
            <div className="courseQuizzesPageCourseQuizzes">
                <div className="singleCoursePageSectionTitle">
                    <SectionTitle title="Quiz Μαθήματος"/>
                </div>
                <ul className="courseQuizzesList">
                    {
                        (
                            isTutorSignedIn() && quizzes.length > 0) || (!isTutorSignedIn() && quizzes.some(quiz => quiz.Visibility)
                        ) ? (
                            quizzes.map((item) => {
                                return (
                                    <div className="singleQuizContainer">
                                        <div className="quizContent">
                                            <div className="singleQuizTextContainer">
                                                <Link className="singleQuizTitle"
                                                      to={`/courses/${courseID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                                <div className="singleQuizDetails">{item.CourseName}</div>
                                                <div
                                                    className="singleQuizDetails">{(!isTutorSignedIn() && item.ShowSubset)
                                                    ? `${item.SubsetSize} ερωτήσεις`
                                                    : `${item.NumberOfQuestions} ερωτήσεις`
                                                }</div>
                                            </div>
                                        </div>
                                        {
                                            isTutorSignedIn() && <div className="quizIcons">
                                                <Link to={`/courses/${courseID}/quizzes/${item.ID}/edit`}>
                                                    <FontAwesomeIcon icon={faPenToSquare} className="quizIcon"/>
                                                </Link>
                                                <FontAwesomeIcon icon={faTrashCan} className="quizIcon" onClick={() => {
                                                    deleteQuiz(item.ID, item.Title)
                                                }}/>
                                            </div>
                                        }
                                    </div>
                                );
                            })
                        ) : (
                            <div className="singleQuizTitle">Δεν υπάρχουν διαθέσιμα quiz.</div>
                        )}
                </ul>
            </div>
        </React.Fragment>

    );
};

export default CourseQuizzesList;