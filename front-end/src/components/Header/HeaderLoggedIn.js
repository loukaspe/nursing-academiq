import React from 'react';
import {Link} from "react-router-dom";

import "./Header.css"
import Logo from "../Logo/Logo";

const HeaderLoggedIn = () => {
    return (
        <>
            <header className="appHeader">
                <div className="appTitle">
                    <Link className="link" to="/">Nursing Academiq</Link>
                </div>
                {/*<Logo/>*/}
                <nav className="nav">
                    <ul className="ul menu">
                        <li>
                            <Link className="link" to="/courses">Τα Μαθήματά Μου</Link>
                        </li>
                        <li>
                            <Link className="link" to="/courses">Κατάλογος Μαθημάτων</Link>
                        </li>
                        <li>
                            <Link className="link" to="/questions">Διαθέσιμα Quiz</Link>
                        </li>
                        <li>
                            <Link className="link" to="/questions">Ιστορικό Quiz</Link>
                        </li>
                    </ul>
                    <ul className="ul myAccountMenu">
                        <li>
                            <Link className="link" to="/questions">Προφίλ</Link>
                        </li>
                        <li>
                            <Link className="link" to="/logout">Αποσύνδεση</Link>
                        </li>
                    </ul>
                </nav>
            </header>
        </>
    )
}

export default HeaderLoggedIn