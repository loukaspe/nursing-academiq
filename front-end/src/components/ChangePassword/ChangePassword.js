import React, {useState} from 'react';
import './ChangePassword.css';
import axios from "axios"; // Import CSS file for styling

import Cookies from "universal-cookie";

const cookies = new Cookies();
const ChangePassword = () => {
    const [formData, setFormData] = useState({
        oldPassword: '',
        newPassword: '',
        newPasswordConfirm: '',
    });
    const [errors, setErrors] = useState({});
    const [apiError, setApiError] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleChange = (e) => {
        const {name, value} = e.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsSubmitting(true);

        // Basic validation
        const errors = {};
        if (!formData.oldPassword) {
            errors.oldPassword = 'Απαιτείται Παλιός Κωδικός';
        }
        if (!formData.newPassword) {
            errors.newPassword = 'Απαιτείται Νέος Κωδικός';
        } else if (formData.newPassword.length < 6) {
            errors.newPassword = 'Ο Κωδικός πρέπει να είναι τουλάχιστον 6 χαρακτήρες';
        }
        if (formData.newPassword !== formData.newPasswordConfirm) {
            errors.newPasswordConfirm = 'Οι κωδικοί πρέπει να ταιριάζουν';
        }

        setErrors(errors);

        if (Object.keys(errors).length === 0) {
            let userCookie = cookies.get("user");
            let userID = userCookie.id;

            let apiUrl = process.env.REACT_APP_API_URL + `/user/${userID}/change_password`;
            try {
                const response = await fetch(apiUrl, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${cookies.get("token")}`,
                    },
                    body: JSON.stringify({
                        old_password: formData.oldPassword,
                        new_password: formData.newPassword,
                    }),
                });
                if (!response.ok) {
                    const errorData = await response.json();
                    setApiError(errorData.message || 'Αποτυχία αλλαγής κωδικού: ελέγξτε τα στοιχεία σας και προσπαθήστε ξανά');
                    setSuccessMessage('');
                } else {
                    console.log('Password changed successfully!');
                    setApiError('');
                    setSuccessMessage('Ο κωδικός άλλαξε με επιτυχία.');
                }
            } catch (error) {
                console.error('Error occurred:', error);
                setApiError('Κάτι πήγε στραβά.');
                setSuccessMessage('');
            }
            setIsSubmitting(false);
        } else {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="change-password-center">
            <div className="change-password-container">
                <h2 className="change-password-title">Αλλαγή Κωδικού</h2>
                <form onSubmit={handleSubmit}>
                    <div className="change-password-form-row">
                        <label htmlFor="oldPassword">Παλιός Κωδικός:</label>
                        <input
                            type="password"
                            id="oldPassword"
                            name="oldPassword"
                            value={formData.oldPassword}
                            onChange={handleChange}
                        />
                    </div>
                    {errors.oldPassword && <div className="change-password-error-row">{errors.oldPassword}</div>}
                    <div className="change-password-form-row">
                        <label htmlFor="newPassword">Νέος Κωδικός:</label>
                        <input
                            type="password"
                            id="newPassword"
                            name="newPassword"
                            value={formData.newPassword}
                            onChange={handleChange}
                        />
                    </div>
                    {errors.newPassword && <div className="change-password-error-row">{errors.newPassword}</div>}
                    <div className="change-password-form-row">
                        <label htmlFor="newPasswordConfirm">Επιβεβαίωση Νέου Κωδικού:</label>
                        <input
                            type="password"
                            id="newPasswordConfirm"
                            name="newPasswordConfirm"
                            value={formData.newPasswordConfirm}
                            onChange={handleChange}
                        />
                    </div>
                    {errors.newPasswordConfirm &&
                        <div className="change-password-error-row">{errors.newPasswordConfirm}</div>}
                    <div className="change-password-form-row">
                        <button type="submit" className="change-password-submit" disabled={isSubmitting}>
                            Υποβολή
                        </button>
                    </div>
                    {apiError && <div className="change-password-error-row">{apiError}</div>}
                    {successMessage && <div className="change-password-success-row">{successMessage}</div>}
                </form>
            </div>
        </div>
    );
};

export default ChangePassword;
