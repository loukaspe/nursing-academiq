import React, {useState} from 'react';
import "./CreateTutor.css";

import api from "../Utilities/APICaller";

const CreateTutor = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [birthDate, setBirthDate] = useState('');
    const [phoneNumber, setPhoneNumber] = useState('');
    const [academicRank, setAcademicRank] = useState('');
    const [error, setError] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const formatBirthDate = (date) => {
        const [year, month, day] = date.split("-");
        return `${month}-${day}-${year}`;
    };
    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsSubmitting(true);

        if (
            username.trim() === '' ||
            firstName.trim() === '' ||
            lastName.trim() === '' ||
            password.trim() === '' ||
            email.trim() === '' ||
            phoneNumber.trim() === '' ||
            academicRank.trim() === '' ||
            birthDate.trim() === ''
        ) {
            setError('Παρακαλώ συμπληρώστε όλα τα πεδία.');
            return;
        }

        if (password.length < 6) {
            setError('Ο Κωδικός πρέπει να είναι τουλάχιστον 6 χαρακτήρες');
            return;
        }

        try {
            let apiUrl = `/tutor`

            await api.post(apiUrl, {
                username: username,
                password: password,
                first_name: firstName,
                last_name: lastName,
                email: email,
                birth_date: formatBirthDate(birthDate),
                phone_number: phoneNumber,
                academic_rank: academicRank,
            }).then((response) => {
                console.log(response.data);
            }).then(() => {
                window.location.href = `/`;
            });
        } catch (error) {
            console.error('Error creating the tutor', error);
            setError('Υπήρξε πρόβλημα κατά την δημιουργία του Εκπαιδευτή. Παρακαλώ δοκιμάστε ξανά.');
        }
        alert("Ο νέος εκπαιδευτής δημιουργήθηκε με επιτυχία.");
        setIsSubmitting(false);
    };

    return (
        <div className="create-tutor-center">
            <div className="create-tutor-container">
                <h2 className="create-tutor-title">Δημιουργία Νέου Εκπαιδευτή</h2>
                <form onSubmit={handleSubmit}>
                    <div className="create-tutor-form-row">
                        <label htmlFor="username">Όνομα Χρήστη:</label>
                        <input
                            type="text"
                            id="username"
                            name="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="password">Κωδικός:</label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <p>     </p>
                        <div className="create-tutor-note">
                            *Ο κωδικός μπορεί να αλλάξει αργότερα απο το Προφίλ του νέου χρήστη
                        </div>
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="firstName">Όνομα:</label>
                        <input
                            type="text"
                            id="firstName"
                            name="firstName"
                            value={firstName}
                            onChange={(e) => setFirstName(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="lastName">Επώνυμο:</label>
                        <input
                            type="text"
                            id="lastName"
                            name="lastName"
                            value={lastName}
                            onChange={(e) => setLastName(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="email">Email:</label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="phoneNumber">Κινητό Τηλέφωνο:</label>
                        <input
                            type="text"
                            id="phoneNumber"
                            name="phoneNumber"
                            value={phoneNumber}
                            onChange={(e) => setPhoneNumber(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="academicRank">Ακαδημαϊκή Βαθμίδα:</label>
                        <input
                            type="text"
                            id="academicRank"
                            name="academicRank"
                            value={academicRank}
                            onChange={(e) => setAcademicRank(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <label htmlFor="birthDate">Ημερομηνία Γέννησης:</label>
                        <input
                            type="date"
                            id="birthDate"
                            name="birthDate"
                            value={birthDate}
                            onChange={(e) => setBirthDate(e.target.value)}
                        />
                    </div>
                    <div className="create-tutor-form-row">
                        <button type="submit" className="create-tutor-submit" disabled={isSubmitting}>
                            Υποβολή
                        </button>
                    </div>
                    {error && <div className="create-tutor-error-row">{error}</div>}
                </form>
            </div>
        </div>
    )
        ;
};

export default CreateTutor;
