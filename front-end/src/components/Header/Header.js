import React, {useCallback, useState} from 'react';

import "./Header.css"
import HeaderLoggedOut from "./HeaderLoggedOut";
import Cookies from "universal-cookie";
import HeaderLoggedIn from "./HeaderLoggedIn";

const cookies = new Cookies();

const Header = ({onMenuToggle, sidebarOpen}) => {
    const [headerNavOpen, setHeaderNavOpen] = useState(false);
    const token = cookies.get("access_token");

    const closeHeaderNav = useCallback(() => setHeaderNavOpen(false), []);
    const toggleHeaderNav = useCallback(() => setHeaderNavOpen((open) => !open), []);

    const handleSidebarToggle = useCallback(() => {
        setHeaderNavOpen(false);
        onMenuToggle();
    }, [onMenuToggle]);

    return token
        ? (
            <HeaderLoggedIn
                onMenuToggle={handleSidebarToggle}
                sidebarOpen={sidebarOpen}
                headerNavOpen={headerNavOpen}
                onHeaderNavToggle={toggleHeaderNav}
                onHeaderNavClose={closeHeaderNav}
            />
        )
        : (
            <HeaderLoggedOut
                onMenuToggle={handleSidebarToggle}
                sidebarOpen={sidebarOpen}
            />
        );
}

export default Header
