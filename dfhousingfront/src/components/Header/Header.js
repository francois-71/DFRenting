import React from "react";
import "./header.css";
import { Link } from "react-router-dom";
const Header = (props) => {
  const { isLoggedIn } = props;
  return (
    <div>
      <nav className="navbar navbar-dark mb-5">
        <div className="container">
          <Link className="navbar-brand" to="/home">
            Home
          </Link>

          {isLoggedIn ? (
            <div>
              <Link className="navbar-brand" to="/profile">
                Profile
              </Link>
              <Link className="navbar-brand" to="/logout">
                Logout
              </Link>
            </div>
          ) : (
            <div>
              <Link className="navbar-brand" to="/login">
                Login
              </Link>
              <Link className="navbar-brand" to="/register">
                Register
              </Link>
            </div>
          )}
        </div>
      </nav>
    </div>
  );
};

export default Header;
