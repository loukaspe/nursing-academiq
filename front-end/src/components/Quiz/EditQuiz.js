import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {useNavigate, useParams} from "react-router-dom";
import "./EditQuiz.css";

import Cookies from "universal-cookie";
import Breadcrumb from "../Utilities/Breadcrumb";

const cookies = new Cookies();
const EditQuiz = ({}) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [isVisible, setIsVisible] = useState(false);
    const [isShowSubsetChecked, setIsShowSubsetChecked] = useState(false);
    const [subsetSize, setSubsetSize] = useState(0);

    //TODO: change courseID
    const [courseID, setCourseID] = useState(1);
    const [courseTitle, setCourseTitle] = useState('');
    const [questions, setQuestions] = useState([]);

    const [error, setError] = useState("");
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();
    let quizID = params.id;

    let navigate = useNavigate();

    useEffect(() => {
        const fetchQuiz = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${quizID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setTitle(response.data.quiz.Title);
                setDescription(response.data.quiz.Description);
                setCourseID(response.data.quiz.Course.ID);
                setIsVisible(response.data.quiz.Visibility);
                setIsShowSubsetChecked(response.data.quiz.ShowSubset);
                setQuestions(response.data.quiz.Questions)
            } catch (error) {
                console.error('Error fetching the quiz data', error);
            }
        };

        const fetchCourse = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setCourseTitle(response.data.course.Course.title);
            } catch (error) {
                console.error('Error fetching the course data', error);
            }
        };

        fetchCourse();
        fetchQuiz();
    }, [quizID]);

    const handleShowSubsetChange = (e) => {
        setIsShowSubsetChecked(e.target.checked);
        if (!e.target.checked) {
            setSubsetSize(0);
        }
    };

    const handleIsVisibleChange = (e) => {
        setIsVisible(e.target.checked);
    };

    const handleSubsetSizeChange = (e) => {
        const value = e.target.value;
        const numValue = parseInt(value, 10);

        if (!isNaN(numValue) && numValue <= questions.length) {
            setSubsetSize(numValue);
        } else if (!isNaN(numValue) && numValue > questions.length) {
            setSubsetSize(questions.length);
            alert(`Το μέγιστο υποσύνολο ερωτήσεων είναι ${questions.length}.`);
        } else {
            setSubsetSize(0);
        }
    };


    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        // Basic validation
        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή quiz.');
            return;
        }

        try {
            let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${quizID}`

            await axios.put(apiUrl, {
                    title: title,
                    description: description,
                    courseID: courseID,
                },
                {
                    headers: {
                        Authorization: `Bearer ${cookies.get("token")}`,
                    },
                });

            window.location.href = `/courses/${courseID}/quizzes/`;
        } catch (error) {
            console.error('Error updating the quiz', error);
            setError('Υπήρξε πρόβλημα κατά την επεξαργασία του quiz. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };

    return (
        <div>
            <Breadcrumb
                actualPath={`/courses/${courseID}/quizzes`}
                namePath={`/Διαχείριση Μαθημάτων/${courseTitle}/Επεξεργασία Quiz`}
            />
            <div className="editQuizContainer">
                <div className="editQuizHeaderRow">
                    <div className="editQuizHeader">
                        <div className="editQuizInfo">
                            <span className="singleChapterQuizzesPageChapterName">Επεξεργασία Quiz</span>
                            <button className="editQuizHeaderButton" onClick={() => navigate(-1)}>Πίσω</button>
                        </div>
                        <button className="editQuizHeaderButton" onClick={() => {
                            alert("questions")
                        }}>Επιλογή Ερωτήσεων
                        </button>
                    </div>
                </div>
                <div className="editQuizDetailsRow">
                    <div className="editQuizDetailsRowColumn">
                        <div className="editQuizDetailsRowInputGroup">
                            <label>Όνομα Quiz</label>
                            <input type="text"
                                   value={title}
                                   className="editQuizDetailsRowInputText"
                            />
                        </div>
                        <div className="editQuizDetailsRowInputGroup">
                            <label>Περιγραφή</label>
                            <input type="text"
                                   value={description}
                                   className="editQuizDetailsRowInputText"
                            />
                        </div>
                    </div>

                    <div className="editQuizDetailsRowColumn">
                        <div className="editQuizCheckboxRow">
                            <label>
                                Ορατό <input type="checkbox"
                                             checked={isVisible}/>
                            </label>
                            <span> Αριθμός Ερωτήσεων: {questions.length}</span>
                        </div>

                        <div className="editQuizCheckboxRow">
                            <label>
                                Τυχαίο Υποσύνολο Ανά Συμπλήρωση
                                <input type="checkbox"
                                       checked={isShowSubsetChecked}
                                       onChange={handleShowSubsetChange}/>
                            </label>
                        </div>

                        <div className={isShowSubsetChecked ? "" : "disabledInput"}>
                            <label className={isShowSubsetChecked ? "" : "disabledInput"}>
                                Πλήθος Ερωτήσεων Ανά Συμπλήρωση: </label>
                            <input type="number" value={subsetSize}
                                   onChange={handleSubsetSizeChange}
                                   disabled={!isShowSubsetChecked}
                                   className="editQuizDetailsRowInputText"
                            />
                        </div>
                    </div>
                </div>
                <div className="editQuizQuestionsRow">
                    <div className="editQuizQuestionsList">
                        <div className="editQuizQuestionsListTitle">Questions List</div>
                        <ul>
                            {questions.map((question, index) => (
                                <li className="editQuizQuestionsListItem" key={index}>{question.Text}</li>
                            ))}
                        </ul>
                    </div>

                    <div className="editQuizButtonsColumn">
                        <button className="editQuizSaveButton">Αποθήκευση</button>
                        <button className="editQuizDeleteButton">Διαγραφή</button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default EditQuiz;
