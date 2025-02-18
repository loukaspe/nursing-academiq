import React, {useEffect, useState} from "react";
import "./CourseQuestionsManager.css";
import Breadcrumb from "../Utilities/Breadcrumb";
import {Link, useNavigate, useParams} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faFileExport, faFileImport, faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import axios from "axios";
import api from "../Utilities/APICaller";

const CourseQuestionsManager = () => {
    const [chapters, setChapters] = useState([]);
    const [selectedChaptersIDs, setSelectedChaptersIDs] = useState([]);
    const [questions, setQuestions] = useState([]);
    const [selectedQuestions, setSelectedQuestions] = useState([]);
    const [course, setCourse] = useState({});

    const params = useParams();
    let courseID = params.courseID;

    let navigate = useNavigate();

    useEffect(() => {
        fetchCourseChaptersQuestions();
    }, []);

    useEffect(() => {
        const filteredQuestions = chapters
            .filter((chapter) => selectedChaptersIDs.includes(chapter.id))
            .flatMap((chapter) => chapter.questions);
        setQuestions(filteredQuestions);
    }, [selectedChaptersIDs, chapters]);

    const fetchCourseChaptersQuestions = () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/courses/${courseID}/questions`

        axios.get(apiUrl, {
            headers: {
                Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
            },
        })
            .then(response => {
                if (response.data.course) {
                    setCourse(response.data.course);
                }

                if (response.data.course.Chapters) {
                    const fetchedChapters = response.data.course.Chapters;
                    setChapters(fetchedChapters);
                    setSelectedChaptersIDs(fetchedChapters.map((chapter) => chapter.id));
                }
            })
            .catch(error => {
                console.error('Error fetching course questions data', error);
            });
    };

    const handleChapterCheckbox = (chapterId) => {
        if (selectedChaptersIDs.includes(chapterId)) {
            setSelectedChaptersIDs(selectedChaptersIDs.filter((id) => id !== chapterId));
        } else {
            setSelectedChaptersIDs([...selectedChaptersIDs, chapterId]);
        }
    };

    const handleQuestionCheckboxChange = (question) => {
        setSelectedQuestions((prevSelected) => {
            if (prevSelected.includes(question)) {
                return prevSelected.filter(q => q !== question);
            } else {
                return [...prevSelected, question];
            }
        });
    };

    const downloadCSV = () => {
        const csvRows = [
            ["Ερώτηση"],
            ...selectedQuestions.map(q => [q.Text])
        ];

        const csvContent = csvRows.map(row => row.join(",")).join("\n");
        const blob = new Blob([csvContent], {type: 'text/csv'});
        const url = URL.createObjectURL(blob);

        const a = document.createElement('a');
        a.href = url;
        a.download = 'questions.csv';
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
    };

    const deleteSelectedQuestions = () => {
        const numToDelete = selectedQuestions.length;
        const confirmDelete = window.confirm(`Είστε σίγουρος/η ότι θέλετε να διαγράψετε ${numToDelete} ερώτηση/εις;`);

        if (confirmDelete) {
            let apiUrl = `/questions/bulk`

            api.post(apiUrl,
                {
                    IDs: selectedQuestions.map(q => q.ID),
                })
                .then(() => {
                    setQuestions((prevQuestions) => prevQuestions.filter(q => !selectedQuestions.includes(q)));
                    setSelectedQuestions([]);
                })
                .catch(error => {
                    console.error('Error deleting questions', error);
                });
        }
    };

    const deleteQuestion = (question) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε την ερώτηση ${question.Text};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = `/questions/${question.ID}`

            api.delete(apiUrl)
                .then(() => {
                    setQuestions((prevQuestions) => prevQuestions.filter(q => q.ID !== question.ID));
                })
                .catch(error => {
                    console.error('Error deleting question', error);
                });
        }
    };

    return (
        <React.Fragment>
            <Breadcrumb
                actualPath={`/courses/${courseID}/questions/manage`}
                namePath={`/Μαθήματα/${course.Title}/Ερωτήσεις`}
            />
            <div className="questionsManagerPageHeader">
                <div className="questionsManagerPageInfo">
                    <span className="questionsManagerPageTitle">Ερωτήσεις</span>

                    <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
                </div>
            </div>
            <div className="questionsManagerSubtitle">
                <div>{course.Title} - {course.NumberOfQuestions} ερωτήσεις</div>
            </div>
            <div className="questionsChaptersContainer">
                <div className="questionsSection">
                    <div className="questionsList">
                        {questions.map((question, index) => (
                            <div key={index} className="questionRow">
                                <div>
                                    <span>
                                        <input
                                            className="questionCheckbox"
                                            type="checkbox"
                                            onChange={() => handleQuestionCheckboxChange(question)}
                                            checked={selectedQuestions.includes(question)}
                                        />
                                    </span>
                                    <span>
                                        {question.Text}
                                    </span>

                                </div>
                                <span className="questionCheckboxContainer">
                                    <Link
                                        to={`/courses/${courseID}/chapters/${question.ChapterID}/questions/${question.ID}/edit`}>
                                        <FontAwesomeIcon icon={faPenToSquare} className="questionIcon"/>
                                    </Link>
                                    <FontAwesomeIcon icon={faTrashCan} className="questionIcon" onClick={() => {
                                        deleteQuestion(question)
                                    }}/>
                                </span>
                            </div>
                        ))}
                    </div>
                </div>
                <div className="chaptersSection">
                    <h2 className="questionsManagerPageTitle">Θεματικές Ενότητες</h2>
                    {chapters.map((chapter) => (
                        <div key={chapter.id} className="chapterRow">
                            <span>{chapter.title}</span>
                            <input
                                type="checkbox"
                                onChange={() => handleChapterCheckbox(chapter.id)}
                                checked={selectedChaptersIDs.includes(chapter.id)}
                            />
                        </div>
                    ))}
                </div>
            </div>
            <div className="questionsChaptersButtonContainer">
                <div className="questionsChaptersLeftButtons">
                    <Link
                        className="questionsChaptersButton"
                        // TODO: change hard coded chapter ID
                        to={`/courses/${courseID}/chapters/1/questions/create`}>
                        + Προσθήκη Νέας
                    </Link>
                    <Link
                        className="questionsChaptersButton"
                        to={`/courses/${courseID}/questions/import`}>
                        <FontAwesomeIcon icon={faFileImport} className="questionsIcon"/> Εισαγωγή Αρχείου
                    </Link>
                    <button
                        className="questionsChaptersButton"
                        onClick={downloadCSV}
                        disabled={selectedQuestions.length === 0}
                    >
                        <FontAwesomeIcon icon={faFileExport} className="questionsIcon"/> Εξαγωγή Αρχείου
                    </button>
                </div>
                <button
                    className="questionsChaptersRightButton"
                    onClick={deleteSelectedQuestions}
                    disabled={selectedQuestions.length === 0}
                >
                    <FontAwesomeIcon icon={faTrashCan} className="questionsIcon"/> Διαγραφή Επιλεγμένων
                </button>
            </div>
        </React.Fragment>
    );
};

export default CourseQuestionsManager;