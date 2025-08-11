import React, {useEffect, useState} from "react";
import "./CourseChaptersList.css";
import Cookies from "universal-cookie";
import {Link, useParams, useNavigate} from "react-router-dom";
import axios from "axios";
import SectionTitle from "../Utilities/SectionTitle";
import Breadcrumb from "../Utilities/Breadcrumb";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const CourseChaptersList = (props) => {
    const [chapters, setChapters] = useState([]);
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

    const deleteChapter = (id, title) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε την ενότητα ${title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/chapter/${id}`

            api.delete(apiUrl)
                .then(() => {
                    window.location.href = `/courses/${courseID}/chapters`;
                })
                .catch(error => {
                    console.error('Error deleting chapter', error);
                });
        }
    };

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

                if (response.data.chapters) {
                    setChapters(response.data.chapters);
                }
            })
            .catch(error => {
                console.error('Error fetching course chapters data', error);
            });
    };

    return (
        <React.Fragment>
            <Breadcrumb actualPath={`/courses/${courseID}/chapters`}
                        namePath={`/Μαθήματα/${course.title}/Θεματικές Ενότητες`}/>
            <div className="singleCourseChaptersPageHeader">
                <div className="singleCourseChaptersPageInfo">
                    <span className="singleCourseChaptersPageCourseName">{course.title}</span>
                    
                </div>
                <Link className="chapterButton" to={`/courses/${courseID}/chapters/create`}>
                    + Νέα Θεματική Ενότητα
                </Link>
            </div>
            <div className="singleCourseDescription">
                <div>{course.description}</div>
            </div>
            <div className="singleCourseChapters">
                <div className="singleCoursePageSectionTitle">
                    <SectionTitle title="Θεματικές Ενότητες"/>
                </div>
                <ul className="courseChaptersList">
                    {chapters.length > 0 ? (
                        chapters.map((item) => {
                            return (
                                <div className="singleChapterContainer">
                                    <div className="singleChapterRowContainer">
                                        <Link className="singleChapterTitle"
                                              to={`/courses/${courseID}/chapters/${item.ID}/quizzes`}>{item.Title}</Link>
                                        {
                                            isTutorSignedIn() && <div className="chapterIcons">

                                                <Link to={`/courses/${props.courseID}/chapters/${item.ID}/edit`}>
                                                    <FontAwesomeIcon icon={faPenToSquare} className="chapterIcon"/>
                                                </Link>
                                                <FontAwesomeIcon icon={faTrashCan} className="chapterIcon" onClick={() => {
                                                    deleteChapter(item.ID, item.Title)
                                                }}/>
                                            </div>
                                        }
                                    </div>
                                    <div className="singleChapterRowContainer">
                                        <div className="singleChapterDetails">{item.Description}</div>
                                    </div>
                                </div>
                            );
                        })
                    ) : (
                        <div className="singleChapterTitle">Δεν υπάρχουν διαθέσιμες θεματικές ενότητες.</div>
                    )}
                </ul>
            </div>
        </React.Fragment>

    );
};

export default CourseChaptersList;