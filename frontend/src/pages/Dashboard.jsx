import React from 'react';
import StatCard from '../components/Dashboard/StatCard';
import NavCard from '../components/Dashboard/NavCard';
import DeploymentCard from '../components/Dashboard/DeploymentCard';
import {
    LayoutDashboard,
    Car,
    User,
    Users,
    Briefcase,
    Shield,
    Monitor,
    Cloud,
    Database,
    AlertCircle
} from 'lucide-react';
import './Dashboard.css';

export const Dashboard = () => {
    return (
        <div className="dashboard-container">
            <StatCard />

            <div className="nav-grid">
                <NavCard
                    icon={LayoutDashboard}
                    title="Dashboard"
                    variant="purple"
                />
                <NavCard
                    icon={Car}
                    title="Monitoring Car"
                    subtitle="LIVE TRACKING • VEHICLE STATUS • HISTORY"
                    variant="orange"
                    path="/monitoring-car"
                />
                <NavCard
                    icon={User}
                    title="Driver"
                    variant="green"
                    path="/driver"
                />
                <NavCard
                    icon={Users}
                    title="User Management"
                    variant="purple"
                    path="/user-management"
                />
                <NavCard
                    icon={Briefcase}
                    title="Karyawan"
                    variant="blue"
                    path="/karyawan"
                />
                <NavCard
                    icon={Briefcase}
                    title="Karyawan"
                    variant="blue"
                    path="/karyawan"
                />
                <NavCard
                    icon={Shield}
                    title="Administration"
                    variant="red"
                    path="/administration"
                />
            </div>

            <div className="deployment-section">
                <h3>Deployment Summary</h3>
                <div className="deployment-grid">
                    <DeploymentCard title="File Services" icon={Monitor} status="success" />
                    <DeploymentCard title="Cloud Services" icon={Monitor} status="success" />
                    <DeploymentCard title="Data Protection" icon={Monitor} status="success" />
                    <DeploymentCard title="Sudit Sog" icon={Monitor} status="warning" />
                </div>
            </div>
        </div>
    );
};
