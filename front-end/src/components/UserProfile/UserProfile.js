import React, {useEffect, useState} from 'react';
import "./UserProfile.css";
import PageTitle from "../Utilities/PageTitle";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faUser} from "@fortawesome/free-solid-svg-icons";

import Cookies from "universal-cookie";
import {Link} from "react-router-dom";
import api from "../Utilities/APICaller";

const cookies = new Cookies();
const UserProfile = () => {
    const [username, setUsername] = useState(null);
    const [name, setName] = useState(null);
    const [academicRank, setAcademicRank] = useState(null);
    const [email, setEmail] = useState(null);

    useEffect(() => {
        fetchUser();
    }, []);

    const fetchUser = async () => {
        let userCookie = cookies.get("user");
        let tutorID = userCookie.specificID;

        let apiUrl = `/tutor/${tutorID}`

        try {
            const response = await api.get(apiUrl);
            // TODO if 401 show unauthorized
            // TODO if 500 show server error

            if (response.status === 500) {
                throw Error(response.data.message);
            }

            if (response.status === 401) {
                throw Error("unauthorized: 401");
            }

            if (response.data.tutor.Tutor === undefined) {
                throw Error("error getting user for tutor");
            }

            if (response.data.tutor.Tutor.first_name && response.data.tutor.Tutor.last_name) {
                setName(response.data.tutor.Tutor.first_name + ' ' + response.data.tutor.Tutor.last_name);
            }

            if (response.data.tutor.Tutor.username) {
                setUsername(response.data.tutor.Tutor.username);
            }

            if (response.data.tutor.Tutor.academic_rank) {
                setAcademicRank(response.data.tutor.Tutor.academic_rank);
            }

            if (response.data.tutor.Tutor.email) {
                setEmail(response.data.tutor.Tutor.email);
            }
        } catch (error) {
            console.error('Error fetching tutor data', error);
        }
    };

    return (
        <>
            <PageTitle title="Το Προφίλ Μου"/>
            <div className="profileContainer">
            <div className="profileContainerBox">
                <div className="basicInfo">
                    <div className="basicSectionOne">
                        <FontAwesomeIcon icon={faUser} size="2xl" className="userIcon"/>
                    </div>
                    <div className="basicSectionTwo">
                        <div className="profileDetails">
                            <div className="profileDetailsText"><span className="profileDetailsTextTitle">Ονοματεπώνυμο Χρήστη: </span>{name}
                            </div>
                            <div className="profileDetailsText"><span
                                className="profileDetailsTextTitle">Username: </span>{username}</div>
                            <div className="profileDetailsText"><span
                                className="profileDetailsTextTitle">Βαθμίδα: </span>{academicRank}
                            </div>
                            <div className="profileDetailsText"><span
                                className="profileDetailsTextTitle">Email: </span>{email}</div>
                            <Link className="changePasswordButton" to="/change-password">Αλλαγή Κωδικού</Link>
                        </div>
                    </div>
                </div>
                </div>
            </div>
        </>
    );
};

export default UserProfile;