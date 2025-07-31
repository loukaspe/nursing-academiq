import React from 'react';
import {Link} from "react-router-dom";

import "./Header.css"

const HeaderLoggedOut = () => {
    return (
        <>
            <header className="appHeader">
                <div className="appTitle">
                    <Link className="link" to="/">Nursing Academiq</Link>
                </div>
                <nav className="nav">
                    <ul className="ul menu">
                        <li>
                            <Link className="link" to="/courses">Μαθήματα</Link>
                        </li>
                        <li>
                            <Link className="link" to="/quizzes">Quizzes</Link>
                        </li>
                    </ul>
                    <ul className="ul myAccountMenu">
                        <li>
                            <Link className="link" to="/login">Σύνδεση</Link>
                        </li>
                    </ul>
                </nav>
            </header>
        </>
    )
}

export default HeaderLoggedOut