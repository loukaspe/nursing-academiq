import React, {useState} from 'react';
import {useParams} from "react-router-dom";
import "./EditChapter.css";

import api from "../Utilities/APICaller";

const CreateChapter = () => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [chapterID, setChapterID] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();
    let courseID = params.courseID;


    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        // Basic validation
        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή ενότητας.');
            setIsSubmitting(false);
            return;
        }

        setError('');
        setIsSubmitting(true);

        try {
            let apiUrl = `/chapter`

            await api.post(apiUrl, {
                title: title,
                description: description,
                courseID: parseInt(courseID),
            }).then((response) => {
                console.log(response.data);
                setChapterID(response.data.insertedID);
            }).then(() => {
                window.location.href = `/courses/${courseID}/chapters`;
            });
        } catch (error) {
            console.error('Error creating the chapter', error);
            setError('Υπήρξε πρόβλημα κατά την δημιουργία της Ενότητας. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };

    return (
        <div className="edit-chapter-center">
            <div className="edit-chapter-container">
                <h2 className="edit-chapter-title">Δημιουργία Ενότητας</h2>
                <form onSubmit={handleSubmit}>
                    <div className="edit-chapter-form-row">
                        <label htmlFor="title">Τίτλος:</label>
                        <input
                            type="text"
                            id="title"
                            name="title"
                            value={title}
                            onChange={(e) => {
                                setTitle(e.target.value)
                                setError('')
                            }
                            }
                        />
                    </div>
                    <div className="edit-chapter-form-row">
                        <label htmlFor="description">Περιγραφή:</label>
                        <input
                            type="text"
                            id="description"
                            name="description"
                            value={description}
                            onChange={(e) => {
                                setDescription(e.target.value)
                                setError('')
                            }}
                        />
                    </div>
                    <div className="edit-chapter-form-row">
                        <button type="submit" className="edit-chapter-submit" disabled={isSubmitting}>
                            Υποβολή
                        </button>
                    </div>
                    {error && <div className="edit-chapter-error-row">{error}</div>}
                </form>
            </div>
        </div>
    )
        ;
};

export default CreateChapter;
