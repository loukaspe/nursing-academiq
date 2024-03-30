import React, {useEffect, useState} from 'react';
import "./UserProfile.css";
import PageTitle from "../Utilities/PageTitle";
import SectionTitle from "../Utilities/SectionTitle";
import axios from 'axios';
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faUser} from "@fortawesome/free-solid-svg-icons";

import Cookies from "universal-cookie";

const cookies = new Cookies();
const UserProfile = () => {
    const [selectedFile, setSelectedFile] = useState(null);
    const [profilePicture, setProfilePicture] = useState(null);
    const [username, setUsername] = useState(null);
    const [name, setName] = useState(null);
    const [registrationNumber, setRegistrationNumber] = useState(null);
    const [birthDate, setBirthDate] = useState(null);
    const [email, setEmail] = useState(null);
    const [phone, setPhone] = useState(null);
    const [completedQuizzes, setCompletedQuizzes] = useState(null);
    const [questionsScore, setQuestionsScore] = useState(null);
    const [percentageOfCorrectAnswers, setPercentageOfCorrectAnswers] = useState(null);

    useEffect(() => {
        fetchUser();
    }, [profilePicture]);

    const handleFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
        uploadFile(event.target.files[0]);
    };

    const handleImageClick = () => {
        document.getElementById('fileInput').click();
    };

    const uploadFile = (file) => {
        let userCookie = cookies.get("user");
        let userID = userCookie.id;

        const formData = new FormData();
        formData.append('image', file);

        let apiUrl = process.env.REACT_APP_API_URL + `/user/${userID}/photo`;

        axios.post(apiUrl, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
                'Authorization': `Bearer ${cookies.get("token")}`,
            },
        })
            .then(response => {
                console.log('Upload successful');
                fetchUser();
            })
            .catch(error => {
                console.error('Error uploading file', error);
            });
    };

    const fetchUser = () => {
        let userCookie = cookies.get("user");
        let studentID = userCookie.specificID;

        let apiUrl = process.env.REACT_APP_API_URL + `/student/${studentID}/extended`

        axios.get(apiUrl, {
            headers: {
                'Authorization': `Bearer ${cookies.get("token")}`,
            },
        })
            .then(response => {
                if (response.data.student.photo) {
                    let photoPath = process.env.REACT_APP_API_URL + response.data.student.photo;
                    setProfilePicture(photoPath);
                }

                if (response.data.student.first_name && response.data.student.last_name) {
                    setName(response.data.student.first_name + ' ' + response.data.student.last_name);
                }

                if (response.data.student.username) {
                    setUsername(response.data.student.username);
                }

                if (response.data.student.registration_number) {
                    setRegistrationNumber(response.data.student.registration_number);
                }


                if (response.data.student.birth_date) {
                    setBirthDate(response.data.student.birth_date);
                }

                if (response.data.student.email) {
                    setEmail(response.data.student.email);
                }

                if (response.data.student.phone_number) {
                    setPhone(response.data.student.phone_number);
                }

                if (response.data.student.completed_quizzes) {
                    setCompletedQuizzes(response.data.student.completed_quizzes);
                }

                if (response.data.student.questions_score) {
                    setQuestionsScore(response.data.student.questions_score);
                }

                if (response.data.student.percentage_of_right_answers) {
                    setPercentageOfCorrectAnswers(response.data.student.percentage_of_right_answers);
                }
            })
            .catch(error => {
                console.error('Error fetching student data', error);
            });
    };

    return (
        <>
            <PageTitle title="Το Προφίλ Μου"/>
            <div className="profileContainer">
                <div className="basicInfo">
                    <div className="basicSectionOne">
                        {profilePicture ? (
                            <img src={profilePicture}
                                 className="profilePicture"
                                 onClick={handleImageClick}/>
                        ) : (
                            <FontAwesomeIcon icon={faUser} size="2xl" className="userIcon" onClick={handleImageClick}/>
                        )}
                        <input type="file" id="fileInput" style={{display: 'none'}} onChange={handleFileChange}/>
                        <br/>
                    </div>
                    <div className="basicSectionTwo">
                        <div className="profileDetails">
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Ονοματεπώνυμο Χρήστη: </span>{name}</div>
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Username: </span>{username}</div>
                            <button className="changePasswordButton">Αλλαγή Κωδικού</button>
                        </div>
                    </div>
                </div>
                <div className="extendedInfo">
                    <div className="extendedSectionOne">
                        <SectionTitle title="Στοιχεία Χρήστη"/>
                        <div className="profileDetails">
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Αιρθμός Μητρώου: </span>{registrationNumber}</div>
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Ημερομηνία Γέννησης: </span>{birthDate}</div>
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Email: </span>{email}</div>
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Τηλέφωνο: </span>{phone}</div>
                        </div>
                    </div>
                    <div className="extendedSectionTwo">
                        <SectionTitle title="Στατιστικά Χρήστη"/>
                        <div className="profileDetails">
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Συμπληρωμένα Quiz: </span>{completedQuizzes}</div>
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Σκορ ερωτήσεων: </span>{questionsScore}</div>
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Ποσοστό Σωστών Ερωτήσεων: </span>{percentageOfCorrectAnswers}</div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default UserProfile;