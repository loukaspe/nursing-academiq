import React, {useState} from "react";
import "./LoginForm.css";
import InputText from "../Input/InputText";
import Cookies from "universal-cookie";
import {jwtDecode} from 'jwt-decode'

const cookies = new Cookies();

const LoginForm = (props) => {
        // const [credentials, setCredentials] = useState({username: '', password: ''})
        // const [jwtToken, setJwtToken] = useState("");

        const [usernameInput, setUsernameInput] = useState("");
        const [passwordInput, setPasswordInput] = useState("");

        async function login(username, password) {
            let requestData = {username: username, password: password};

            const response = await fetch(apiUrl, {
                method: 'POST',
                body: JSON.stringify(requestData),
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${process.env.REACT_APP_API_KEY}`,
                },
                credentials: 'include',
            });
            const result = await response.json();
            // TODO if 401 show unauthorized
            // TODO if 500 show server error
            if (response.status === 500) {
                throw Error(result.message);
            }

            if (response.status === 401) {
                throw Error("unauthorized: 401");
            }

            if (result.token === undefined) {
                throw Error("unauthorized: no token");
            }


            cookies.set(
                "result",
                result,
                {
                    path: "/",
                }
            );
            cookies.set("token", result.token, {
                path: "/",
            });

            const userInfo = jwtDecode(result.token).UserInfo;
            console.log(userInfo)
            cookies.set(
                "user",
                {
                    id: userInfo.UserID,
                    type: userInfo.User.UserType,
                    specificID: userInfo.User.SpecificID
                },
                {
                    path: "/",
                }
            );

            window.location.href = "/";
        }

        const apiUrl = process.env.REACT_APP_API_URL + "/login";

        const onSubmitHandler = (event) => {
            event.preventDefault();

            // for empty submit
            if (usernameInput.trim().length <= 0)
                return;
            setUsernameInput('')

            if (passwordInput.trim().length <= 0)
                return;
            setPasswordInput('')

            // setCredentials({username: usernameInput, password: passwordInput})

            login(usernameInput, passwordInput)
                .catch(error => console.log(error))
        };

        const onUsernameChangeHandler = (event) => {
            setUsernameInput(event.target.value)
        };

        const onPasswordChangeHandler = (event) => {
            setPasswordInput(event.target.value)
        };

        return (
            <div className="loginForm">
                <h1>Σύνδεση Χρήστη</h1>
                <hr/>
                <form onSubmit={onSubmitHandler}>
                    <InputText
                        placeholder=""
                        label="Όνομα Χρήστη"
                        id="username_input"
                        onChangeHandler={onUsernameChangeHandler}
                    />
                    <InputText
                        placeholder=""
                        label="Κωδικός"
                        id="password_input"
                        onChangeHandler={onPasswordChangeHandler}
                    />
                    <button className="submitButton" type="submit">Σύνδεση</button>
                </form>
            </div>
        );
    }
;

export default LoginForm;