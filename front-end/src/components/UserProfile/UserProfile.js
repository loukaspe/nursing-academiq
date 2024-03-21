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

    useEffect(() => {
        fetchProfilePicture();
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
                // Refresh the profile picture after successful upload
                fetchProfilePicture();
            })
            .catch(error => {
                console.error('Error uploading file', error);
            });
    };

    const fetchProfilePicture = () => {
        let userCookie = cookies.get("user");
        let userID = userCookie.id;

        let apiUrl = process.env.REACT_APP_API_URL + `/user/${userID}/photo`

        axios.get(apiUrl, {
            headers: {
                'Authorization': `Bearer ${cookies.get("token")}`,
            },
        })
            .then(response => {
                if (response.data.path) {
                    let photoPath = process.env.REACT_APP_API_URL + response.data.path;
                    setProfilePicture(photoPath);
                }
            })
            .catch(error => {
                console.error('Error fetching profile picture', error);
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
                        Details
                    </div>
                </div>
                <div className="extendedInfo">
                    <div className="extendedSectionOne">
                        <SectionTitle title="Στοιχεία Χρήστη"/>
                    </div>
                    <div className="extendedSectionTwo">
                        <SectionTitle title="Στατιστικά Χρήστη"/>
                    </div>
                </div>
            </div>
        </>
    );
};

export default UserProfile;