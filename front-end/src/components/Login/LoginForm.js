import React, {useState} from "react";
import "./LoginForm.css";
import InputText from "../Input/InputText";
import Cookies from "universal-cookie";
import {jwtDecode} from 'jwt-decode'
import InputPassword from "../Input/InputPassword";

const cookies = new Cookies();

const LoginForm = () => {
        const [usernameInput, setUsernameInput] = useState("");
        const [passwordInput, setPasswordInput] = useState("");

        async function login(username, password) {
            const apiUrl = process.env.REACT_APP_API_URL + "/login";
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

            if (result.access_token === undefined) {
                throw Error("unauthorized: no token");
            }

            cookies.set(
                "result",
                result,
                {
                    path: "/",
                }
            );
            cookies.set("access_token", result.access_token, {
                path: "/",
            });

            const userInfo = jwtDecode(result.access_token).UserInfo;
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

        const onSubmitHandler = (event) => {
            event.preventDefault();

            if (usernameInput.trim().length <= 0)
                return;
            setUsernameInput('')

            if (passwordInput.trim().length <= 0)
                return;
            setPasswordInput('')

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
                <h1 className="loginFormTitle">Σύνδεση Χρήστη</h1>
                <hr/>
                <form onSubmit={onSubmitHandler}>
                    <InputText
                        placeholder=""
                        label="Όνομα Χρήστη"
                        id="username_input"
                        onChangeHandler={onUsernameChangeHandler}
                        className="loginFormInput"
                    />
                    <InputPassword
                        placeholder=""
                        label="Κωδικός"
                        id="password_input"
                        onChangeHandler={onPasswordChangeHandler}
                        className="loginFormInput"
                    />
                    <button className="submitButton" type="submit">Σύνδεση</button>
                </form>
            </div>
        );
    }
;

export default LoginForm;