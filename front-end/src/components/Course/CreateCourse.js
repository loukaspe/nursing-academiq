import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {useParams} from "react-router-dom";
import "./EditCourse.css";

import Cookies from "universal-cookie";

const cookies = new Cookies();
const CreateCourse = () => {
    let userCookie = cookies.get("user");
    let specificID = userCookie.specificID;

    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [courseID, setCourseID] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const params = useParams();

    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        // Basic validation
        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή μαθήματος.');
            return;
        }

        try {
            let apiUrl = process.env.REACT_APP_API_URL + `/course`

            await axios.post(apiUrl, {
                    course: {
                        title: title,
                        description: description,
                        tutorID: specificID,
                    }
                },
                {
                    headers: {
                        Authorization: `Bearer ${cookies.get("token")}`,
                    },
                }).then((response) => {
                console.log(response.data);
                setCourseID(response.data.insertedID);
            }).then(() => {
                window.location.href = `/courses/${courseID}`;
            });


        } catch (error) {
            console.error('Error creating the course', error);
            setError('Υπήρξε πρόβλημα κατά την δημιουργία του Μαθήματος. Παρακαλώ δοκιμάστε ξανά.');
        }
        setIsSubmitting(false);
    };

    return (
        <div className="edit-course-center">
            <div className="edit-course-container">
                <h2 className="edit-course-title">Δημιουργία Μαθήματος</h2>
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

export default CreateCourse;
