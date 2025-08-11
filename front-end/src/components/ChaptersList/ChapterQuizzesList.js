import React, {useEffect, useState} from "react";
import "./ChapterQuizzesList.css";
import Cookies from "universal-cookie";
import {useParams, useNavigate, Link} from "react-router-dom";
import axios from "axios";
import SectionTitle from "../Utilities/SectionTitle";
import Breadcrumb from "../Utilities/Breadcrumb";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const ChapterQuizzesList = (props) => {
    const [quizzes, setQuizzes] = useState([]);
    const [chapter, setChapter] = useState({});
    const [course, setCourse] = useState({});

    const params = useParams();
    let chapterID = params.chapterID;
    let courseID = params.courseID;

    const token = cookies.get("access_token");

    const isTutorSignedIn = () => {
        return !!token;
    }

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

    const deleteQuiz = (id, title) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε το quiz ${title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/quiz/${id}`

            api.delete(apiUrl)
                .then(() => {
                    window.location.href = `/courses/${courseID}/chapters`;
                })
                .catch(error => {
                    console.error('Error deleting quiz', error);
                });
        }
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
                    {
                        isTutorSignedIn() &&
                        <Link to={`/courses/${courseID}/chapters/${chapterID}/edit`}>
                            <FontAwesomeIcon icon={faPenToSquare} className="chapterIcon"/>
                        </Link>
                    }
                    
                </div>
                {
                    isTutorSignedIn()
                    &&
                    <>
                        <Link className="courseButton" to={`/courses/${courseID}/quizzes/create`}>
                            + Προσθήκη Quiz
                        </Link>
                        <button className="courseDangerButton" onClick={() => {
                            deleteChapter(chapterID, chapter.title)
                        }}>Διαγραφή Ενότητας
                        </button>
                    </>
                }
            </div>
            <div className="singleChapterDescription">
                <div>{chapter.description}</div>
            </div>
            <div className="singleChapterQuizzes">
                <div className="singleChapterPageSectionTitle">
                    <SectionTitle title="Quiz Ενότητας"/>
                </div>
                <ul className="chapterQuizzesList">
                    {quizzes.length > 0 ? (
                        quizzes.map((item) => {
                            return (
                                <div className="chaptersSingleQuizContainer">
                                    <div className="singleQuizRowContainer">
                                        <Link className="singleQuizTitle"
                                              to={`/courses/${courseID}/quizzes/${item.ID}`}>{item.Title}</Link>
                                        {
                                            isTutorSignedIn() && <div className="chapterIcons">

                                                <Link to={`/courses/${props.courseID}/quizzes/${item.ID}/edit`}>
                                                    <FontAwesomeIcon icon={faPenToSquare} className="chapterIcon"/>
                                                </Link>
                                                <FontAwesomeIcon icon={faTrashCan} className="chapterIcon" onClick={() => {
                                                    deleteQuiz(item.ID, item.Title)
                                                }}/>
                                            </div>
                                        }
                                    </div>
                                    <div className="singleQuizRowContainer">
                                        <div className="singleQuizDetails">{item.Description}</div>
                                    </div>
                                    <div className="singleQuizRowContainer">
                                        <div className="singleQuizDetails">{item.NumberOfQuestions} Ερωτήσεις</div>
                                    </div>
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

export default ChapterQuizzesList;