import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { useUser } from "../contexts/UserContext";

function Header() {
  const { user, logout } = useUser();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <header>
      <nav
        className="navbar navbar-expand-lg fixed-top navbar-dark"
        style={{
          backgroundColor: "#342f2f",
          paddingLeft: "1rem",
          paddingRight: "1rem",
        }}
      >
        <div className="container">
          <Link className="navbar-brand" to="/">
            <img
              // src="https://assets-lighthouse.s3.amazonaws.com/uploads/image/file/6227/restaurant-list-logo.png"
              width="30"
              height="30"
              className="d-inline-block align-top"
              alt=""
            />
          </Link>

          {user && (
            <>
              <button
                className="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarCollapse"
                aria-controls="navbarCollapse"
                aria-expanded="false"
                aria-label="Toggle navigation"
              >
                <span className="navbar-toggler-icon"></span>
              </button>
              <div className="collapse navbar-collapse" id="navbarCollapse">
                <ul className="navbar-nav ms-auto mb-2 mb-md-0">
                  {user.user_type === "ADMIN" && (
                    <li className="nav-item">
                      <Link to="/users/admin">
                        <button className="btn btn-outline-secondary my-2 mx-1">
                          Admin 後台
                        </button>
                      </Link>
                    </li>
                  )}
                  <li className="nav-item d-inline-flex align-items-center">
                    <span className="text-white nav-link">
                      Hi, {user.email}
                    </span>
                  </li>
                  <li className="nav-item">
                    <Link to="/user/profile">
                      <button className="btn btn-outline-warning my-2 mx-1">
                        Profile
                      </button>
                    </Link>
                  </li>
                  {/* <li className="nav-item">
                    <Link to="/restaurants/create">
                      <button className="btn btn-outline-info my-2 mx-1">
                        新增餐廳
                      </button>
                    </Link>
                  </li> */}
                  <li className="nav-item">
                    <button
                      onClick={handleLogout}
                      className="btn btn-outline-danger my-2 mx-1"
                    >
                      Logout
                    </button>
                  </li>
                </ul>
              </div>
            </>
          )}
        </div>
      </nav>
      <div className="banner"></div>
    </header>
  );
}

export default Header;
