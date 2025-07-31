import React from 'react';
import {Outlet} from "react-router-dom";

import "./Layout.css"
import Header from "../Header/Header";
import Sidebar from "../Sidebar/Sidebar";

const Layout = () => {
    return (
        <>
            <div className="layout-container">
                <Sidebar />
                <main className="layout-main">
                    <Header/>
                    <Outlet/>
                </main>
            </div>
        </>

    );
}

export default Layout