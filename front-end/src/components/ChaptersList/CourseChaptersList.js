import React, {useEffect, useState} from "react";
import "./CourseChaptersList.css";
import Cookies from "universal-cookie";
import {Link, useParams, useNavigate} from "react-router-dom";
import axios from "axios";
import SectionTitle from "../Utilities/SectionTitle";
import LimitedRecentCourseChaptersList from "./LimitedCourseChaptersList";
import {useHistory} from "react-router-dom";
import Breadcrumb from "../Utilities/Breadcrumb";

const cookies = new Cookies();

const CourseChaptersList = (props) => {
    const [chapters, setChapters] = useState([]);
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
            <Breadcrumb actualPath={`/courses/${courseID}/chapters`} namePath={`/Μαθήματα/${course.title}/Θεματικές Ενότητες`}/>
            <div className="singleCourseChaptersPageHeader">
                <div className="singleCourseChaptersPageInfo">
                    <span className="singleCourseChaptersPageCourseName">{course.title}</span>
                </div>
                <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
            </div>
            <div className="singleCourseDescription">
                <div>{course.description}</div>
            </div>
            <div className="singleCourseChapters">
                <div className="singleCoursePageSectionTitle">
                    <SectionTitle title="Θεματικές Ενότητες"/>
                </div>
                <ul className="courseChaptersList">
                    {chapters.map((item) => {
                        return (
                            <div className="singleChapterTextContainer">
                                <Link className="singleChapterTitle"
                                      to={`/courses/${courseID}/chapters/${item.ID}/quizzes`}>{item.Title}</Link>
                                <div className="singleChapterDetails">{item.Description}</div>
                            </div>
                        );
                    })}
                </ul>
            </div>
        </React.Fragment>

    );
};

export default CourseChaptersList;