import React from 'react';
import {Outlet} from "react-router-dom";

import "./Layout.css"
import Header from "../Header/Header";

const Layout = () => {
    return (
        <>
            <Header/>
            <Outlet/>
            <div>
                <h2>Footer</h2>
            </div>
        </>

    );
}

export default Layout