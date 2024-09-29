import React, { useState } from 'react';
import './EditQuestion.css';
import {Link, useNavigate} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPenToSquare} from "@fortawesome/free-solid-svg-icons"; // Import the CSS file

const EditQuestion = () => {
    // Move chapters to state with two random strings as initial values
    const [chapters] = useState(['Chapter 1', 'Chapter 2']);
    const [selectedChapter, setSelectedChapter] = useState('');
    const [questionText, setQuestionText] = useState('');
    const [answers, setAnswers] = useState([{ text: '', isCorrect: false }, { text: '', isCorrect: false }]);
    const [explanation, setExplanation] = useState('');
    const [source, setSource] = useState('');

    let navigate = useNavigate();

    const handleChapterChange = (e) => setSelectedChapter(e.target.value);
    const handleQuestionTextChange = (e) => setQuestionText(e.target.value);
    const handleAnswerChange = (index, field, value) => {
        const updatedAnswers = [...answers];
        updatedAnswers[index][field] = value;
        setAnswers(updatedAnswers);
    };
    const addAnswer = () => setAnswers([...answers, { text: '', isCorrect: false }]);
    const removeAnswer = (index) => setAnswers(answers.filter((_, i) => i !== index));
    const handleExplanationChange = (e) => setExplanation(e.target.value);
    const handleSourceChange = (e) => setSource(e.target.value);
    const handleSave = () => {
        // Save logic here
        console.log('Question Saved');
    };
    const handleDelete = () => {
        // Delete logic here
        console.log('Question Deleted');
    };

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
                        <option key={index} value={chapter}>{chapter}</option>
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
                                    value={answer.text}
                                    onChange={(e) => handleAnswerChange(index, 'text', e.target.value)}
                                />
                            </div>
                            <label className="editQuestionPageCheckbox">
                                <input
                                    type="checkbox"
                                    checked={answer.isCorrect}
                                    onChange={(e) => handleAnswerChange(index, 'isCorrect', e.target.checked)}
                                />
                            </label>
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
                        <button onClick={handleDelete} className="editQuestionPageDeleteButton">Διαγραφή</button>
                        <button onClick={handleSave} className="editQuestionPageButton">Αποθήκευση</button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default EditQuestion
