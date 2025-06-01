import React, {useEffect, useState} from "react";
import "./QuestionsSelect.css";
import Breadcrumb from "../Utilities/Breadcrumb";
import {Link, useNavigate, useParams} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare, faTrashCan} from "@fortawesome/free-solid-svg-icons";
import axios from "axios";
import Cookies from "universal-cookie";
import api from "../Utilities/APICaller";

const QuestionsSelect = () => {
    const [chapters, setChapters] = useState([]);
    const [selectedChaptersIDs, setSelectedChaptersIDs] = useState([]);
    const [questions, setQuestions] = useState([]);
    const [selectedQuestions, setSelectedQuestions] = useState([]);
    const [course, setCourse] = useState({});
    const [quizName, setQuizName] = useState("");
    const [error, setError] = useState("");

    const params = useParams();
    let courseID = params.courseID;
    let quizID = params.quizID

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

    useEffect(() => {
        fetchQuiz();
    }, [questions])

    const fetchQuiz = async () => {
        let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${quizID}`

        try {
            const response = await axios.get(apiUrl, {
                headers: {
                    Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                },
            });
            setQuizName(response.data.quiz.Title);

            const quizQuestionIds = response.data.quiz.Questions.map(q => q.ID);

            setSelectedQuestions(questions.filter(question =>
                quizQuestionIds.includes(question.ID)
            ));
        } catch (error) {
            console.error('Error fetching the quiz data', error);
        }
    };

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

    const saveQuizQuestions = async (event) => {
        event.preventDefault();

        try {
            let apiUrl = `/quiz/${quizID}/questions`

            await api.post(apiUrl, {
                questionsIDs: selectedQuestions.map((question) => question.ID),
            });

            window.location.href = `/courses/${courseID}/quizzes`;
        } catch (error) {
            console.error('Error updating the quiz questions', error);
            setError('Υπήρξε πρόβλημα κατά την επεξαργασία των ερωτήσεων του quiz. Παρακαλώ δοκιμάστε ξανά.');
        }
    }

    const deleteQuestion = (question) => {
        const confirmMessage = `Είστε σίγουρος ότι θέλετε να διαγράψετε την ερώτηση ${question.Text};`;

        if (window.confirm(confirmMessage)) {
            let apiUrl = process.env.REACT_APP_API_URL + `/questions/${question.ID}`

            api.delete(apiUrl,)
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
                actualPath={`/courses/${courseID}/questions/quizzes/${quizID}/select`}
                namePath={`/Μαθήματα/${course.Title}/Quiz/Επεξεργασία Quiz/Ερωτήσεις`}
            />
            <div className="questionsSelectPageHeader">
                <div className="questionsSelectPageInfo">
                    <span className="questionsSelectPageTitle">Ερωτήσεις Quiz</span>

                    <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
                </div>
            </div>
            <div className="questionsSelectSubtitle">
                <div>{quizName}</div>
            </div>
            <div className="questionsSelectChaptersContainer">
                <div className="questionsSection">
                    <div className="questionsList">
                        {questions.length > 0 ? (
                            questions.map((question, index) => (
                                <div key={index} className="questionSelectRow">
                                    <div className="questionSelectRowTop">
                                        <div className="questionSelectDetails">
                                        <span>
                                            <input
                                                className="questionSelectCheckbox"
                                                type="checkbox"
                                                onChange={() => handleQuestionCheckboxChange(question)}
                                                checked={selectedQuestions.includes(question)}
                                            />
                                        </span>
                                            <span>{question.Text}</span>
                                        </div>
                                        <span className="questionSelectCheckboxContainer">
                                        <Link
                                            to={`/courses/${courseID}/chapters/${question.ChapterID}/questions/${question.ID}/edit`}
                                        >
                                            <FontAwesomeIcon icon={faPenToSquare} className="questionIcon"/>
                                        </Link>
                                        <FontAwesomeIcon
                                            icon={faTrashCan}
                                            className="questionIcon"
                                            onClick={() => deleteQuestion(question)}
                                        />
                                    </span>
                                    </div>
                                    <div className="questionSelectChapterName">Θεματική
                                        Ενότητα: {question.Chapter.title}</div>
                                </div>

                            ))
                        ) : (
                            <div className="questionSelectDetails">Δεν υπάρχουν διαθέσιμες ερωτήσεις.</div>
                        ) }
                    </div>
                </div>
                <div className="chaptersSection">
                    <h2 className="questionsSelectPageTitle">Θεματικές Ενότητες</h2>
                    {chapters.length > 0 ? (
                        chapters.map((chapter) => (
                            <div key={chapter.id} className="chapterRow">
                                <span>{chapter.title}</span>
                                <input
                                    type="checkbox"
                                    onChange={() => handleChapterCheckbox(chapter.id)}
                                    checked={selectedChaptersIDs.includes(chapter.id)}
                                />
                            </div>
                        ))
                    ) : (
                        <div className="mySingleCourseTitle">Δεν υπάρχουν διαθέσιμες θεματικές ενότητες.</div>
                    )}
                </div>
            </div>
            {error && <div className="questionsSelectErrorRow">{error}</div>}
            <div className="questionsSelectChaptersButtonContainer">
                <div className="questionsSelectChaptersLeft questionsSelectPageTitle">
                    Τρέχων Αριθμός Ερωτήσεων : {selectedQuestions.length}
                </div>
                <div className="questionsSelectChaptersRightButtons">
                    <button
                        className="questionsSelectChaptersSaveButton"
                        onClick={saveQuizQuestions}
                    >
                        Αποθήκευση
                    </button>
                    <button
                        className="questionsSelectChaptersCancelButton"
                        onClick={() => navigate(-1)}
                    >
                        Ακύρωση
                    </button>
                </div>
            </div>
        </React.Fragment>
    );
};

export default QuestionsSelect;