import React, {useState} from "react";
import "./LoginPage.css";
import LoginForm from "./LoginForm";
import Logo from "../Logo/Logo";


const LoginPage = (props) => {
        return (
            <>
                <div className="loginPageContainer">
                    <div className="formContainer">
                        <LoginForm/>
                    </div>
                    <div className="logoContainer">
                        <Logo/>
                    </div>
                </div>
                <div style={{clear: 'both'}}></div>
            </>
        );
    }
;

export default LoginPage;