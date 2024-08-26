import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {useParams} from "react-router-dom";
import "./EditChapter.css";

import Cookies from "universal-cookie";

const cookies = new Cookies();
const EditChapter = ({}) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [courseID, setCourseID] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();
    let chapterID = params.id;

    useEffect(() => {
        const fetchChapter = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/chapter/${chapterID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setTitle(response.data.chapter.title);
                setDescription(response.data.chapter.description);
                setCourseID(response.data.chapter.course.ID);
            } catch (error) {
                console.error('Error fetching the chapter data', error);
            }
        };

        fetchChapter();
    }, [chapterID]);

    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        // Basic validation
        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή μαθήματος.');
            return;
        }

        try {
            let apiUrl = process.env.REACT_APP_API_URL + `/chapter/${chapterID}`

            await axios.put(apiUrl, {
                    title: title,
                    description: description
                },
                {
                    headers: {
                        Authorization: `Bearer ${cookies.get("token")}`,
                    },
                });

            window.location.href = `/courses/${courseID}/chapters/${chapterID}/quizzes`;
        } catch (error) {
            console.error('Error updating the chapter', error);
            setError('Υπήρξε πρόβλημα κατά την επεξαργασία του Μαθήματος. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };

    return (
        <div className="edit-chapter-center">
            <div className="edit-chapter-container">
                <h2 className="edit-chapter-title">Επεξεργασία Ενότητας</h2>
                <form onSubmit={handleSubmit}>
                    <div className="edit-chapter-form-row">
                        <label htmlFor="title">Τίτλος:</label>
                        <input
                            type="text"
                            id="title"
                            name="title"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div className="edit-chapter-form-row">
                        <label htmlFor="description">Περιγραφή:</label>
                        <input
                            type="text"
                            id="description"
                            name="description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
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

export default EditChapter;
