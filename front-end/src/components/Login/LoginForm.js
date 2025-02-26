import React, {useState} from "react";
import "./LoginForm.css";
import InputText from "../Input/InputText";
import Cookies from "universal-cookie";
import {jwtDecode} from "jwt-decode";
import InputPassword from "../Input/InputPassword";

const cookies = new Cookies();

const LoginForm = () => {
    const [usernameInput, setUsernameInput] = useState("");
    const [passwordInput, setPasswordInput] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const login = async (username, password) => {
        setError("");
        setLoading(true);

        try {
            const apiUrl = `${process.env.REACT_APP_API_URL}/login`;
            const requestData = {username, password};

            const response = await fetch(apiUrl, {
                method: "POST",
                body: JSON.stringify(requestData),
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                },
                credentials: "include",
            });

            const result = await response.json();

            if (response.status === 401) {
                setPasswordInput("");
                throw new Error("Λάθος όνομα χρήστη ή κωδικός.");
            }

            if (!response.ok) {
                throw new Error(result.error || `Σφάλμα: ${response.status}`);
            }

            if (!result.access_token) {
                throw new Error("Unauthorized: No token received");
            }

            cookies.set("result", result, {path: "/"});
            cookies.set("access_token", result.access_token, {path: "/"});

            const userInfo = jwtDecode(result.access_token).UserInfo;
            cookies.set(
                "user",
                {
                    id: userInfo.UserID,
                    type: userInfo.User.UserType,
                    specificID: userInfo.User.SpecificID,
                },
                {path: "/"}
            );

            window.location.href = "/";
        } catch (error) {
            setError(error.message);
        } finally {
            setLoading(false);
        }
    };

    const onSubmitHandler = (event) => {
        event.preventDefault();

        if (!usernameInput.trim()) {
            setError("Το όνομα χρήστη είναι υποχρεωτικό.");
            return;
        }

        if (!passwordInput.trim()) {
            setError("Ο κωδικός είναι υποχρεωτικός.");
            return;
        }

        login(usernameInput, passwordInput);
    };

    return (
        <div className="loginForm">
            <h1 className="loginFormTitle">Σύνδεση Χρήστη</h1>
            <hr/>
            {error && <p className="loginErrorMessage">{error}</p>}
            <form onSubmit={onSubmitHandler}>
                <InputText
                    label="Όνομα Χρήστη"
                    id="username_input"
                    onChangeHandler={(e) => setUsernameInput(e.target.value)}
                    className="loginFormInput"
                    value={usernameInput}
                />
                <InputPassword
                    label="Κωδικός"
                    id="password_input"
                    onChangeHandler={(e) => {
                        setPasswordInput(e.target.value);
                        setError("")
                    }}
                    className="loginFormInput"
                    value={passwordInput}
                />
                <button className="submitButton" type="submit" disabled={loading}>
                    {loading ? "Σύνδεση..." : "Σύνδεση"}
                </button>
            </form>
        </div>
    );
};

export default LoginForm;
