import React from 'react';
import { Outlet } from 'react-router-dom';
import Sidebar from './Sidebar';
import TopBar from './TopBar';
import './Layout.css';

export const Layout = () => {
    const [isSidebarOpen, setIsSidebarOpen] = React.useState(true);

    const toggleSidebar = () => {
        setIsSidebarOpen(!isSidebarOpen);
    };

    return (
        <div className={`app-container ${isSidebarOpen ? '' : 'sidebar-collapsed'}`}>
            <Sidebar isOpen={isSidebarOpen} />
            <main className="main-content">
                <TopBar onToggleSidebar={toggleSidebar} />
                <div className="content-area">
                    <Outlet />
                </div>
            </main>
        </div>
    );
};
