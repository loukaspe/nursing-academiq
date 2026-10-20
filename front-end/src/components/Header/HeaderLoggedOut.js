import React from 'react';
import {Link} from "react-router-dom";

import "./Header.css"

const HeaderLoggedOut = ({onMenuToggle, sidebarOpen}) => {
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
    )
}

export default HeaderLoggedOut
