import React, { useState} from "react";
import "./LimitedCourseChaptersList.css";
import Cookies from "universal-cookie";
import {Link} from "react-router-dom";

const cookies = new Cookies();

const LimitedCourseChaptersList = (props) => {
    const [visibleChapters, setVisibleChapters] = useState(2);

    return (
        <React.Fragment>
            <ul className="chaptersList">
                {props.chapters.slice(0, visibleChapters).map((item) => {
                    return (
                        <div className="singleChapterTextContainer">
                            <div className="singleChapterTitle">{item.Title}</div>
                            <div className="singleChapterDetails">{item.Description}</div>
                        </div>
                    );
                })}
                {
                    props.chapters.length > visibleChapters &&
                    (
                        <Link className="moreButton"
                              to={`/courses/${props.courseID}/chapters`}>+ Όλες οι Ενότητες</Link>
                    )
                }
            </ul>
        </React.Fragment>
    );
};

export default LimitedCourseChaptersList;