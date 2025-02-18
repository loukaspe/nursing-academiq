import React from 'react';

import "./Header.css"
import HeaderLoggedOut from "./HeaderLoggedOut";
import Cookies from "universal-cookie";
import HeaderLoggedIn from "./HeaderLoggedIn";

const cookies = new Cookies();

const Header = () => {
    const token = cookies.get("access_token");
    return token ? <HeaderLoggedIn/> : <HeaderLoggedOut/>;
}

export default Header