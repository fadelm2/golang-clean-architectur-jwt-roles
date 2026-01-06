import React from 'react';
import { NavLink } from 'react-router-dom';
import {
    LayoutDashboard,
    Car,
    User,
    Briefcase,
    Users,
    Shield,
    Settings
} from 'lucide-react';
import './Sidebar.css';

const Sidebar = ({ isOpen }) => {
    return (
        <div className={`sidebar ${isOpen ? '' : 'collapsed'}`}>
            <div className="brand">
                <div className="logo-icon">
                    <div className="grid-icon"></div>
                </div>
                <span className="brand-name">FLASH</span>
            </div>

            <nav className="nav-menu">
                <NavLink to="/" className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}>
                    <LayoutDashboard size={20} />
                    <span>Dashboard</span>
                </NavLink>
                <NavLink to="/monitoring-car" className="nav-item">
                    <Car size={20} />
                    <span>Monitoring Car</span>
                </NavLink>
                <NavLink to="/driver" className="nav-item">
                    <User size={20} />
                    <span>Driver</span>
                </NavLink>
                <NavLink to="/karyawan" className="nav-item">
                    <Briefcase size={20} />
                    <span>Karyawan</span>
                </NavLink>
                <NavLink to="/user-management" className="nav-item">
                    <Users size={20} />
                    <span>User Management</span>
                </NavLink>
                <NavLink to="/administration" className="nav-item">
                    <Shield size={20} />
                    <span>Administration</span>
                </NavLink>
                <NavLink to="/settings" className="nav-item settings-item">
                    <Settings size={20} />
                    <span>Settings</span>
                </NavLink>
            </nav>

            <div className="user-profile-bottom">
                <img src="https://ui-avatars.com/api/?name=Jason+Smith&background=random" alt="User" />
                <div className="user-info">
                    <div className="name">Jason Smith</div>
                    <div className="role">Admin</div>
                </div>
            </div>
        </div>
    );
};

export default Sidebar;
