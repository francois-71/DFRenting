import React from 'react';
import './logout.css';

const Logout = () => {

    const handleLogout = () => {
        localStorage.removeItem("token");
        window.location.href = "/home";
    }

    return (
        <div>
            <button onClick={handleLogout}>Logout</button>
        </div>
    );
};

export default Logout;