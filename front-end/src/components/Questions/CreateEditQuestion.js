import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {useNavigate, useParams} from "react-router-dom";
import "./EditQuestion.css";

import Cookies from "universal-cookie";

const cookies = new Cookies();
const CreateEditQuestion = ({}) => {
    const [text, setText] = useState('');
    const [explanation, setExplanation] = useState('');
    const [source, setSource] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();
    let questionID = params.id;

    let navigate = useNavigate();

    useEffect(() => {
        const fetchQuestion = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/questions/${questionID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setText(response.data.question.Text);
                setExplanation(response.data.question.Explanation);
                setSource(response.data.question.Source);
            } catch (error) {
                console.error('Error fetching the question data', error);
            }
        };

        fetchQuestion();
    }, [questionID]);

    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        if (text.trim() === '' || explanation.trim() === '' || source.trim() === '') {
            setError('Παρακαλώ συμπληρώστε κείμενο, εξήγηση και πηγή ερώτησης.');
            return;
        }

        try {
            let apiUrl = process.env.REACT_APP_API_URL + `/questions/${questionID}`

            await axios.put(apiUrl, {
                    text: text,
                    explanation: explanation,
                    source: source,
                },
                {
                    headers: {
                        Authorization: `Bearer ${cookies.get("token")}`,
                    },
                });

            navigate(-1)
        } catch (error) {
            console.error('Error updating the question', error);
            setError('Υπήρξε πρόβλημα κατά την επεξαργασία της Ερώτησης. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };

    return (
        <div className="edit-question-center">
            <div className="edit-question-container">
                <h2 className="edit-question-text">Επεξεργασία Ερώτησης</h2>
                <form onSubmit={handleSubmit}>
                    <div className="edit-question-form-row">
                        <label htmlFor="text">Κείμενο:</label>
                        <input
                            type="text"
                            id="text"
                            name="text"
                            value={text}
                            onChange={(e) => setText(e.target.value)}
                        />
                    </div>
                    <div className="edit-question-form-row">
                        <label htmlFor="explanation">Εξήγηση:</label>
                        <input
                            type="text"
                            id="explanation"
                            name="explanation"
                            value={explanation}
                            onChange={(e) => setExplanation(e.target.value)}
                        />
                    </div>
                    <div className="edit-question-form-row">
                        <label htmlFor="explanation">Πηγή:</label>
                        <input
                            type="text"
                            id="source"
                            name="source"
                            value={source}
                            onChange={(e) => setSource(e.target.value)}
                        />
                    </div>
                    <div className="edit-question-form-row">
                        <button type="submit" className="edit-question-submit" disabled={isSubmitting}>
                            Υποβολή
                        </button>
                    </div>
                    {error && <div className="edit-question-error-row">{error}</div>}
                </form>
            </div>
        </div>
    )
        ;
};

export default CreateEditQuestion
