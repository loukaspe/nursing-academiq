import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {useParams} from "react-router-dom";
import "./EditCourse.css";

import api from "../Utilities/APICaller";

const EditCourse = ({courseId}) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();
    let courseID = params.id;

    useEffect(() => {
        const fetchCourse = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setTitle(response.data.course.Course.title);
                setDescription(response.data.course.Course.description);
            } catch (error) {
                console.error('Error fetching the course data', error);
            }
        };

        fetchCourse();
    }, [courseID]);

    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή μαθήματος.');
            return;
        }

        try {
            let apiUrl = `/course/${courseID}`

            await api.put(apiUrl, {
                    title: title,
                    description: description
                });

            window.location.href = `/courses/${courseID}`;
        } catch (error) {
            console.error('Error updating the course', error);
            setError('Υπήρξε πρόβλημα κατά την επεξαργασία του Μαθήματος. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };

    return (
        <div className="edit-course-center">
            <div className="edit-course-container">
                <h2 className="edit-course-title">Επεξεργασία Μαθήματος</h2>
                <form onSubmit={handleSubmit}>
                    <div className="edit-course-form-row">
                        <label htmlFor="title">Τίτλος:</label>
                        <input
                            type="text"
                            id="title"
                            name="title"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div className="edit-course-form-row">
                        <label htmlFor="description">Περιγραφή:</label>
                        <input
                            type="text"
                            id="description"
                            name="description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </div>
                    <div className="edit-course-form-row">
                        <button type="submit" className="edit-course-submit" disabled={isSubmitting}>
                            Υποβολή
                        </button>
                    </div>
                    {error && <div className="edit-course-error-row">{error}</div>}
                </form>
            </div>
        </div>
    )
        ;
};

export default EditCourse;
