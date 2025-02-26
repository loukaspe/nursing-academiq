import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {Link, useNavigate, useParams} from "react-router-dom";
import "./EditQuiz.css";

import Cookies from "universal-cookie";
import CoursesListModal from "../CoursesList/CoursesListModal";
import Breadcrumb from "../Utilities/Breadcrumb";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const CreateQuiz = () => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [isVisible, setIsVisible] = useState(false);
    const [isShowSubsetChecked, setIsShowSubsetChecked] = useState(false);
    const [subsetSize, setSubsetSize] = useState(0);

    //TODO: change courseID
    const [courseTitle, setCourseTitle] = useState('');
    const [questions, setQuestions] = useState([]);

    // TODO: error handling in form works ?
    const [error, setError] = useState("");
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();

    const [isModalOpen, setIsModalOpen] = useState(true);
    const [selectedCourseID, setSelectedCourseID] = useState(params.courseID || "");

    let navigate = useNavigate();


    useEffect(() => {
        console.log("useEffect " + selectedCourseID);
        const fetchCourse = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/course/${selectedCourseID}`

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
    }, [selectedCourseID]);

    const handleCourseSelect = (courseId) => {
        console.log("aaaaaaaa" + courseId);
        setSelectedCourseID(courseId);
        setIsModalOpen(false);
    };

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

        if (!selectedCourseID) {
            setError("Παρακαλώ επιλέξτε μάθημα.");
            return;
        }

        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή quiz.');
            return;
        }

        try {
            let apiUrl = `/quiz`

            const response = await api.post(apiUrl, {
                Title: title,
                Description: description,
                CourseID: parseInt(selectedCourseID),
                Visibility: isVisible,
                ShowSubset: isShowSubsetChecked,
                SubsetSize: subsetSize,
            });


            if (response.status === 201 && response.data.insertedID) {
                const newQuizID = response.data.insertedID;

                window.location.href = `/courses/${selectedCourseID}/quizzes/${newQuizID}/questions/select`;
            } else {
                console.error("Quiz creation failed, unexpected response:", response);
            }
        } catch (error) {
            console.error('Error creating the quiz', error);
            setError('Υπήρξε πρόβλημα κατά την δημιουργία του quiz. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };


    return (
        <div>
            {
                selectedCourseID === "" &&
                <CoursesListModal isOpen={isModalOpen} onCourseSelect={handleCourseSelect}/>
            }
            {/*<CoursesListModal isOpen={isModalOpen} onCourseSelect={handleCourseSelect}/>*/}
            <Breadcrumb
                actualPath={`/courses/${selectedCourseID}/quizzes`}
                namePath={`/Διαχείριση Μαθημάτων/${courseTitle}/Δημιουργία Quiz`}
            />
            <div className="editQuizContainer">
                <div className="editQuizHeaderRow">
                    <div className="editQuizHeader">
                        <div className="editQuizInfo">
                            <span className="singleChapterQuizzesPageChapterName">Δημιουργία Quiz</span>
                            <button className="editQuizHeaderButton" onClick={() => navigate(-1)}>Πίσω</button>
                        </div>
                    </div>
                </div>
                <div className="editQuizDetailsRow">
                    <div className="editQuizDetailsRowColumn">
                        <div className="editQuizDetailsRowInputGroup">
                            <label>Όνομα Quiz</label>
                            <input type="text"
                                   value={title}
                                   className="editQuizDetailsRowInputText"
                                   onChange={(e) => setTitle(e.target.value)}
                            />
                        </div>
                        <div className="editQuizDetailsRowInputGroup">
                            <label>Περιγραφή</label>
                            <input type="text"
                                   value={description}
                                   className="editQuizDetailsRowInputText"
                                   onChange={(e) => setDescription(e.target.value)}
                            />
                        </div>
                    </div>

                    <div className="editQuizDetailsRowColumn">
                        <div className="editQuizCheckboxRow">
                            <label>
                                Ορατό <input type="checkbox"
                                             checked={isVisible}
                                             onChange={handleIsVisibleChange}/>
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
                        <div className="editQuizQuestionsListTitle">Ερωτήσεις</div>
                        <ul>
                            {questions.map((question, index) => (
                                <li className="editQuizQuestionsListItem" key={index}>{question.Text}</li>
                            ))}
                        </ul>
                    </div>

                    <div className="editQuizButtonsColumn">
                        <button className="editQuizSaveButton" onClick={handleSubmit}>Αποθήκευση και Επιλογή Ερωτήσεων
                        </button>
                        <button className="editQuizDeleteButton">Διαγραφή</button>
                    </div>
                </div>
            </div>
        </div>
    )
        ;
};

export default CreateQuiz;
