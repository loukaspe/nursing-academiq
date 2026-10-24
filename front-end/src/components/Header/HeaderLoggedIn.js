import React, {useEffect} from 'react';
import {Link, useLocation} from "react-router-dom";

import "./Header.css"

const HeaderLoggedIn = ({
    onMenuToggle,
    sidebarOpen,
    headerNavOpen,
    onHeaderNavToggle,
    onHeaderNavClose,
}) => {
    const location = useLocation();

    useEffect(() => {
        onHeaderNavClose();
    }, [location.pathname, onHeaderNavClose]);

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

            <Link to="/" className="appHeader-brand" onClick={onHeaderNavClose}>
                Nursing AcademIQ
            </Link>

            <button
                type="button"
                className="header-nav-toggle"
                onClick={onHeaderNavToggle}
                aria-expanded={headerNavOpen ? 'true' : 'false'}
                aria-controls="header-main-nav"
                aria-label={headerNavOpen ? 'Κλείσιμο μενού' : 'Άνοιγμα μενού'}
            >
                <span className="header-nav-toggle-icon" aria-hidden="true">⋮</span>
            </button>

            <nav
                id="header-main-nav"
                className={`nav nav--logged-in ${headerNavOpen ? 'is-open' : ''}`}
                aria-label="Κύρια πλοήγηση"
            >
                <ul className="ul menu">
                    <li>
                        <Link className="link" to="/my-courses" onClick={onHeaderNavClose}>
                            Διαχείριση Μαθημάτων
                        </Link>
                    </li>
                    <li>
                        <Link className="link" to="/my-quizzes" onClick={onHeaderNavClose}>
                            Διαχείριση Quiz
                        </Link>
                    </li>
                    <li>
                        <Link className="link" to="/create-tutor" onClick={onHeaderNavClose}>
                            Προσθήκη Καθηγητή
                        </Link>
                    </li>
                </ul>
                <ul className="ul myAccountMenu">
                    <li>
                        <Link className="link" to="/profile" onClick={onHeaderNavClose}>
                            Προφίλ
                        </Link>
                    </li>
                    <li>
                        <Link className="link" to="/logout" onClick={onHeaderNavClose}>
                            Αποσύνδεση
                        </Link>
                    </li>
                </ul>
            </nav>

            {headerNavOpen && (
                <button
                    type="button"
                    className="header-nav-backdrop"
                    onClick={onHeaderNavClose}
                    aria-label="Κλείσιμο μενού"
                />
            )}
        </header>
    )
}

export default HeaderLoggedIn
