import React, {useEffect, useState} from "react";
import "./ChapterQuizzesList.css";
import Cookies from "universal-cookie";
import {useParams, useNavigate} from "react-router-dom";
import axios from "axios";
import SectionTitle from "../Utilities/SectionTitle";
import Breadcrumb from "../Utilities/Breadcrumb";

const cookies = new Cookies();

const ChapterQuizzesList = (props) => {
    const [quizzes, setQuizzes] = useState([]);
    const [chapter, setChapter] = useState({});
    const [course, setCourse] = useState({});

    const params = useParams();
    let chapterID = params.chapterID;
    let courseID = params.courseID;

    let navigate = useNavigate();

    useEffect(() => {
        fetchChapter();
    }, []);

    const fetchChapter = () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/chapter/${chapterID}`

        axios.get(apiUrl, {
            headers: {
                Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
            },
        })
            .then(response => {
                console.log(response);
                if (response.data.chapter) {
                    setChapter(response.data.chapter);
                }

                if (response.data.chapter.quizzes) {
                    setQuizzes(response.data.chapter.quizzes);
                }

                if (response.data.chapter.course) {
                    setCourse(response.data.chapter.course);
                }
            })
            .catch(error => {
                console.error('Error fetching chapter quizzes data', error);
            });
    };

    return (
        <React.Fragment>
            <Breadcrumb
                actualPath={`/courses/${courseID}/chapters/${courseID}`}
                namePath={`/Μαθήματα/${course.Title}/Θεματικές Ενότητες/${chapter.title}`}
            />
            <div className="singleChapterQuizzesPageHeader">
                <div className="singleChapterQuizzesPageInfo">
                    <span className="singleChapterQuizzesPageChapterName">{chapter.title}</span>
                </div>
                <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
            </div>
            <div className="singleChapterDescription">
                <div>{chapter.description}</div>
            </div>
            <div className="singleChapterQuizzes">
                <div className="singleChapterPageSectionTitle">
                    <SectionTitle title="Quiz Ενότητας"/>
                </div>
                <ul className="chapterQuizzesList">
                    {quizzes.map((item) => {
                        return (
                            <div className="singleQuizTextContainer">
                                <div className="singleQuizTitle">{item.Title}</div>
                                <div className="singleQuizDetails">{item.Description}</div>
                                <div className="singleQuizDetails">{item.NumberOfQuestions} Ερωτήσεις</div>
                            </div>
                        );
                    })}
                </ul>
            </div>
        </React.Fragment>

    );
};

export default ChapterQuizzesList;