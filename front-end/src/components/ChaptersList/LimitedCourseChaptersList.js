import React, {useState} from "react";
import "./LimitedCourseChaptersList.css";
import Cookies from "universal-cookie";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";

const cookies = new Cookies();

const LimitedCourseChaptersList = (props) => {
    const [visibleChapters, setVisibleChapters] = useState(2);

    const token = cookies.get("token");

    const isTutorSignedIn = () => {
        return !!token;
    }

    return (
        <React.Fragment>
            <ul className="chaptersList">
                {props.chapters.slice(0, visibleChapters).map((item) => {
                    return (
                        <div className="singleQuizContainer">
                            <div className="quizContent">
                                <div className="singleChapterTextContainer">
                                    <Link className="singleChapterTitle"
                                          to={`/courses/${props.courseID}/chapters/${item.ID}/quizzes`}>{item.Title}</Link>
                                    <div className="singleChapterDetails">{item.Description}</div>
                                </div>
                            </div>
                            {
                                isTutorSignedIn() && <div className="chapterIcons">
                                    <FontAwesomeIcon icon={faPenToSquare} className="chapterIcon" onClick={() => {
                                        alert("edit")
                                    }}/>
                                    <FontAwesomeIcon icon={faTrashCan} className="chapterIcon" onClick={() => {
                                        alert("delete")
                                    }}/>
                                </div>
                            }
                        </div>
                    );
                })}
                <div
                    className={`quizzesButtonContainer ${props.chapters.length > visibleChapters ? 'multiple' : 'single'}`}>
                    {
                        isTutorSignedIn() && <Link className="myCoursesListButton" to="/my-courses">+ Νέο Quiz</Link>
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