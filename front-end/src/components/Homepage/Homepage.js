import React from "react";
import Cookies from "universal-cookie";
import HomepageLoggedIn from "./HomepageLoggedIn";
import HomepageLoggedOut from "./HomepageLoggedOut";

const cookies = new Cookies();

const Homepage = () => {
    const token = cookies.get("access_token");
    return token ? <HomepageLoggedIn/> : <HomepageLoggedOut/>;
};

export default Homepage;