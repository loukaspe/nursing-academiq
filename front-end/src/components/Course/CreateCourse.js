import React, {useState} from 'react';
import "./EditCourse.css";

import Cookies from "universal-cookie";
import api from "../Utilities/APICaller";

const cookies = new Cookies();

const CreateCourse = () => {
    let userCookie = cookies.get("user");
    let specificID = userCookie.specificID;

    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [courseID, setCourseID] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        if (title.trim() === '' || description.trim() === '') {
            setError('Παρακαλώ συμπληρώστε τίτλο και περιγραφή μαθήματος.');
            return;
        }

        try {
            let apiUrl = `/course`

            await api.post(apiUrl, {
                course: {
                    title: title,
                    description: description,
                    tutorID: specificID,
                }
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
                        <label htmlFor="title">*Τίτλος:</label>
                        <input
                            type="text"
                            id="title"
                            name="title"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div className="edit-course-form-row">
                        <label htmlFor="description">*Περιγραφή:</label>
                        <input
                            type="text"
                            id="description"
                            name="description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </div>
                    <div className="edit-course-form-row">
                        <p>Τα πεδία με αστερίσκο είναι υποχρεωτικά.</p>
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
