import React from 'react';
import {Link} from "react-router-dom";

import "./Header.css"

const HeaderLoggedIn = ({onMenuToggle, sidebarOpen}) => {
    return (
        <header className="appHeader">
            <button
                type="button"
                className="menu-toggle"
                onClick={onMenuToggle}
                aria-expanded={sidebarOpen ? 'true' : 'false'}
                aria-label={sidebarOpen ? 'Κλείσιμο μενού' : 'Άνοιγμα μενού'}
            >
                <span className="menu-toggle-bar"/>
                <span className="menu-toggle-bar"/>
                <span className="menu-toggle-bar"/>
            </button>
            <nav className="nav" aria-label="Κύρια πλοήγηση">
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
    )
}

export default HeaderLoggedIn
