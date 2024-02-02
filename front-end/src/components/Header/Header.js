import React from 'react';
import {Link} from "react-router-dom";

import "./Header.css"
import Logo from "../Logo/Logo";
import HeaderLoggedOut from "./HeaderLoggedOut";
import Cookies from "universal-cookie";
import HeaderLoggedIn from "./HeaderLoggedIn";

const cookies = new Cookies();

const Header = () => {
    const token = cookies.get("token");
    return token ? <HeaderLoggedIn/> : <HeaderLoggedOut/>;
}

export default Header