import React from 'react';
import {Link} from "react-router-dom";

import "./Header.css"

const HeaderLoggedIn = () => {
    return (
        <>
            <header className="appHeader">
                {/*<Logo/>*/}
                <nav className="nav">
                    <ul className="ul menu">
                        <li>
                            <Link className="link" to="/my-courses">Διαχείριση Μαθημάτων</Link>
                        </li>
                        <li>
                            <Link className="link" to="/my-quizzes">Διαχείριση Quiz</Link>
                        </li>
                        <li>
                            <Link className="link" to="/create-tutor">Προσθήκη Καθηγητή</Link>
                        </li>
                    </ul>
                    <ul className="ul myAccountMenu">
                        <li>
                            <Link className="link" to="/profile">Προφίλ</Link>
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