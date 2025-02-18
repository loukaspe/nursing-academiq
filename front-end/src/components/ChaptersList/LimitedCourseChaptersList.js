import React, {useState} from "react";
import "./LimitedCourseChaptersList.css";
import Cookies from "universal-cookie";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const LimitedCourseChaptersList = (props) => {
    const [visibleChapters, setVisibleChapters] = useState(2);

    const token = cookies.get("access_token");

    const isTutorSignedIn = () => {
        return !!token;
    }

    const deleteChapter = (id, title) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε την ενότητα ${title};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/chapter/${id}`

            api.delete(apiUrl)
                .then(() => {
                    window.location.href = `/courses/${props.courseID}/chapters`;
                })
                .catch(error => {
                    console.error('Error deleting chapter', error);
                });
        }
    };

    return (
        <React.Fragment>
            <ul className="chaptersList">
                {props.chapters.slice(0, visibleChapters).map((item) => {
                    return (
                        <div className="singleChapterContainer">
                            <div className="singleChapterRowContainer">
                                <Link className="singleChapterTitle"
                                      to={`/courses/${props.courseID}/chapters/${item.ID}/quizzes`}>{item.Title}</Link>
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
                })}
                <div
                    className={`quizzesButtonContainer ${props.chapters.length > visibleChapters ? 'multiple' : 'single'}`}>
                    {
                        isTutorSignedIn() &&
                        <Link className="myCoursesListButton" to={`/courses/${props.courseID}/chapters/create`}>+ Νέα
                            Ενότητα
                        </Link>
                    }
                    {
                        props.chapters.length > visibleChapters &&
                        <Link className="moreButton" to={`/courses/${props.courseID}/chapters`}>+ Όλες οι
                            Ενότητες</Link>
                    }
                </div>
            </ul>
        </React.Fragment>
    );
};

export default LimitedCourseChaptersList;