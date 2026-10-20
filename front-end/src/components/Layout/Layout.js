import React, {useCallback, useState} from 'react';
import {Outlet} from "react-router-dom";

import "./Layout.css"
import Header from "../Header/Header";
import Sidebar from "../Sidebar/Sidebar";

const Layout = () => {
    const [sidebarOpen, setSidebarOpen] = useState(false);

    const closeSidebar = useCallback(() => setSidebarOpen(false), []);
    const toggleSidebar = useCallback(() => setSidebarOpen((open) => !open), []);

    return (
        <div className="layout-container">
            {sidebarOpen && (
                <button
                    type="button"
                    className="sidebar-backdrop"
                    onClick={closeSidebar}
                    aria-label="Κλείσιμο μενού"
                />
            )}
            <Sidebar isOpen={sidebarOpen} onNavigate={closeSidebar}/>
            <main className="layout-main">
                <Header onMenuToggle={toggleSidebar} sidebarOpen={sidebarOpen}/>
                <div className="layout-content">
                    <Outlet/>
                </div>
            </main>
        </div>
    );
}

export default Layout
