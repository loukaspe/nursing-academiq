import React from 'react';

import "./Header.css"
import HeaderLoggedOut from "./HeaderLoggedOut";
import Cookies from "universal-cookie";
import HeaderLoggedIn from "./HeaderLoggedIn";

const cookies = new Cookies();

const Header = ({onMenuToggle, sidebarOpen}) => {
    const token = cookies.get("access_token");
    return token
        ? <HeaderLoggedIn onMenuToggle={onMenuToggle} sidebarOpen={sidebarOpen}/>
        : <HeaderLoggedOut onMenuToggle={onMenuToggle} sidebarOpen={sidebarOpen}/>;
}

export default Header
