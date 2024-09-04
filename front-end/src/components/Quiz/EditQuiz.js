import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {useParams} from "react-router-dom";
import "./EditQuiz.css";

import Cookies from "universal-cookie";

const cookies = new Cookies();
const EditQuiz = ({}) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [courseID, setCourseID] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();
    let quizID = params.id;

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
            } catch (error) {
                console.error('Error fetching the quiz data', error);
            }
        };

        fetchQuiz();
    }, [quizID]);

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
        <div className="edit-quiz-center">
            <div className="edit-quiz-container">
                <h2 className="edit-quiz-title">Επεξεργασία Quiz</h2>
                <form onSubmit={handleSubmit}>
                    <div className="edit-quiz-form-row">
                        <label htmlFor="title">Τίτλος:</label>
                        <input
                            type="text"
                            id="title"
                            name="title"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div className="edit-quiz-form-row">
                        <label htmlFor="description">Περιγραφή:</label>
                        <input
                            type="text"
                            id="description"
                            name="description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </div>
                    <div className="edit-quiz-form-row">
                        <button type="submit" className="edit-quiz-submit" disabled={isSubmitting}>
                            Υποβολή
                        </button>
                    </div>
                    {error && <div className="edit-quiz-error-row">{error}</div>}
                </form>
            </div>
        </div>
    )
        ;
};

export default EditQuiz;
