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
                aria-label={sidebarOpen ? 'Κλείσιμο μενού μαθημάτων' : 'Άνοιγμα μενού μαθημάτων'}
            >
                <span className="menu-toggle-bar"/>
                <span className="menu-toggle-bar"/>
                <span className="menu-toggle-bar"/>
            </button>

            <Link to="/" className="appHeader-brand">
                Nursing AcademIQ
            </Link>

            <nav className="nav nav--logged-out" aria-label="Κύρια πλοήγηση">
                <ul className="ul menu menu--desktop">
                    <li>
                        <Link className="link" to="/quizzes">Quizzes</Link>
                    </li>
                </ul>
                <ul className="ul myAccountMenu">
                    <li>
                        <Link className="link link--accent" to="/login">Σύνδεση</Link>
                    </li>
                </ul>
            </nav>
        </header>
    )
}

export default HeaderLoggedOut
