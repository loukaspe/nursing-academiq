import React, {useState} from "react";
import "./LoginPage.css";
import LoginForm from "./LoginForm";
import Logo from "../Logo/Logo";

const LoginPage = () => {
        return (
            <>
                <div className="loginPageContainer">
                    <div className="logoContainer">
                        <Logo/>
                    </div>
                    <div className="formContainer">
                        <LoginForm/>
                    </div>
                </div>
                <div style={{clear: 'both'}}></div>
            </>
        );
    }
;

export default LoginPage;