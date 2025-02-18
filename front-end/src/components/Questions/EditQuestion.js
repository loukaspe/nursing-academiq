import React, {useEffect, useState} from 'react';
import './EditQuestion.css';
import {useNavigate, useParams} from "react-router-dom";
import axios from "axios";

import Cookies from "universal-cookie";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const EditQuestion = () => {
    const [chapters, setChapters] = useState([]);
    const [selectedChapter, setSelectedChapter] = useState('');
    const [questionText, setQuestionText] = useState('');
    const [answers, setAnswers] = useState([{Text: '', IsCorrect: false}, {Text: '', IsCorrect: false}]);
    const [explanation, setExplanation] = useState('');
    const [source, setSource] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    let navigate = useNavigate();

    const params = useParams();

    let questionID = params.id;
    let courseID = params.courseID;
    let chapterID = params.chapterID;

    const handleChapterChange = (e) => setSelectedChapter(e.target.value);
    const handleQuestionTextChange = (e) => setQuestionText(e.target.value);
    const handleAnswerChange = (index, field, value) => {
        const updatedAnswers = [...answers];
        updatedAnswers[index][field] = value;
        setAnswers(updatedAnswers);
    };
    const addAnswer = () => setAnswers([...answers, {Text: '', IsCorrect: false}]);
    const removeAnswer = (index) => setAnswers(answers.filter((_, i) => i !== index));
    const handleExplanationChange = (e) => setExplanation(e.target.value);
    const handleSourceChange = (e) => setSource(e.target.value);
    const handleSave = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        if (explanation.trim() === '' || source.trim() === '' || questionText.trim() === '' || selectedChapter === '') {
            setError('Παρακαλώ συμπληρώστε όλα τα πεδία.');
            return;
        }

        if (answers.length < 2) {
            setError('Παρακαλώ συμπληρώστε τουλάχιστον 2 απαντήσεις για την ερώτηση.');
            return;
        }

        try {
            let apiUrl = `/questions/${questionID}`

            let multipleCorrectAnswers = answers.filter(answer => answer.IsCorrect).length > 1;

            let filteredAnswers = answers.filter(answer => answer.Text.trim() !== '');
            setAnswers(filteredAnswers)

            await api.put(apiUrl, {
                text: questionText,
                explanation: explanation,
                source: source,
                multipleCorrectAnswers: multipleCorrectAnswers,
                numberOfAnswers: answers.length,
                answers: filteredAnswers,
                courseID: parseInt(courseID),
                chapterID: chapters.find(chapter => chapter.Title === selectedChapter).ID
            });

            setIsSubmitting(false);
            alert("Η Ερώτηση άλλαξε με επιτυχία.");
        } catch (error) {
            console.error('Error updating the question', error);
            setError('Υπήρξε πρόβλημα κατά την επεξαργασία της Ερώτησης . Παρακαλώ δοκιμάστε ξανά.');
        }
    };
    const handleDelete = async () => {
        try {
            let apiUrl = `/questions/${questionID}`

            await api.delete(apiUrl);
            navigate(-1);
        } catch (error) {
            console.error('Error deleting the question', error);
            setError('Υπήρξε πρόβλημα κατά την διαγραφή της Ερώτησης . Παρακαλώ δοκιμάστε ξανά.');
        }
    };

    useEffect(() => {
        const fetchQuestion = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/questions/${questionID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setAnswers(response.data.question.Answers)
                setExplanation(response.data.question.Explanation)
                setSource(response.data.question.Source)
                setQuestionText(response.data.question.Text)
                setSelectedChapter(response.data.question.Chapter.title)
            } catch (error) {
                console.error('Error fetching the question data', error);
            }
        };
        const fetchCourseChapters = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}/chapters`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setChapters(response.data.chapters)
            } catch (error) {
                console.error('Error fetching the course chapters', error);
            }
        };

        fetchQuestion();
        fetchCourseChapters();
    }, []);

    return (
        <div className="editQuestionPageContainer">

            <div className="editQuestionPageInfo">
                <span className="editQuestionPageTitle">Διαχείριση Ερώτησης</span>
                <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
            </div>
            <div className="editQuestionPageChapterSection">
                <label htmlFor="chapter-select">Θεματική Ενότητα:</label>
                <select
                    id="chapter-select"
                    value={selectedChapter}
                    onChange={handleChapterChange}
                >
                    <option value="">Επιλέξτε Ενότητα</option>
                    {chapters.map((chapter, index) => (
                        <option key={index} value={chapter.Title}>{chapter.Title}</option>
                    ))}
                </select>
            </div>

            <div className="editQuestionPageQuestionSection">
                <label htmlFor="question-text">
                    Εκφώνηση <span className="editQuestionPageQuestionSectionLabelSpanText">(εώς 500 χαρακτήρες)</span>
                </label>
                <input
                    id="question-text"
                    type="text"
                    value={questionText}
                    onChange={handleQuestionTextChange}
                    maxLength={500}
                />
            </div>

            <div className="editQuestionPageAnswersMetadataSection">
                <div className="editQuestionPageAnswersSection">
                    <div className="editQuestionPageAnswersHeader">
                        <span>Απαντήσεις</span>
                        <span>Σωστό/Λάθος</span>
                    </div>

                    {answers.map((answer, index) => (
                        <div key={index} className="editQuestionPageAnswerRow">
                            <div className="editQuestionPageAnswerLabelAndInput">
                                <label htmlFor={`answer-text-${index}`}>Απάντηση {index + 1}</label>
                                <input
                                    id={`answer-text-${index}`}
                                    type="text"
                                    value={answer.Text}
                                    onChange={(e) => handleAnswerChange(index, 'Text', e.target.value)}
                                />
                            </div>
                            <label className="editQuestionPageCheckbox">
                                <input
                                    type="checkbox"
                                    checked={answer.IsCorrect}
                                    onChange={(e) => handleAnswerChange(index, 'IsCorrect', e.target.checked)}
                                />
                            </label>
                            {answers.length > 2 && (
                                <button
                                    className="editQuestionPageDeleteAnswerButton"
                                    onClick={() => removeAnswer(index)}
                                >
                                    Αφαίρεση
                                </button>
                            )}
                        </div>
                    ))}

                    <div className="editQuestionPageCenteredButton">
                        <button className="editQuestionPageButton" onClick={addAnswer}>+ Προσθήκη Απάντησης</button>
                    </div>
                </div>
                <div className="editQuestionPageMetadataSection">
                    <label htmlFor="explanation-input">Επεξήγηση</label>
                    <input
                        id="explanation-input"
                        type="text"
                        value={explanation}
                        onChange={handleExplanationChange}
                    />
                    <label htmlFor="source-input">Πηγή</label>
                    <input
                        id="source-input"
                        type="text"
                        value={source}
                        onChange={handleSourceChange}
                    />
                    <div className="editQuestionPageActionButtons">
                        {error && <div className="editQuestionErrorRow">{error}</div>}
                        <button onClick={handleDelete} className="editQuestionPageDeleteButton">Διαγραφή</button>
                        <button onClick={handleSave} className="editQuestionPageButton"
                                disabled={isSubmitting}>Αποθήκευση
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default EditQuestion
