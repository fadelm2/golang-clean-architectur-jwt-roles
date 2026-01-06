import React from 'react';
import { Bell, Trophy, Maximize, MoreHorizontal, ChevronDown, Menu } from 'lucide-react';
import './TopBar.css';

const TopBar = ({ onToggleSidebar }) => {
    return (
        <div className="topbar">
            <div className="greeting-section" style={{ display: 'flex', alignItems: 'center', gap: '15px' }}>
                <button className="menu-btn" onClick={onToggleSidebar}>
                    <Menu size={24} />
                </button>
                <div>
                    <h1>Welcome back, Jason!</h1>
                    <p>Fleet Monitoring Dashboard</p>
                </div>
            </div>

            <div className="actions-section">
                <div className="action-button">
                    <Bell size={18} />
                    <span className="badge">v2.1.1</span>
                </div>

                <div className="action-button">
                    <Trophy size={18} />
                    <div className="notification-dot"></div>
                </div>

                <div className="action-button">
                    <Maximize size={18} />
                </div>

                <div className="action-button">
                    <MoreHorizontal size={18} />
                </div>
            </div>
        </div>
    );
};

export default TopBar;
