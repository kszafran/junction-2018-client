import React from "react";
import avatarPlaceholder from "../../assets/images/avatar.png";
import graphImage from "../../assets/images/graph.png";
import appLogo from "../../assets/images/logo_top.png";
import { Link } from "react-router-dom";

import "./index.css";

const Header = () => {
  return (
    <div className="header">
      <div />
      <Link to="/">
        <img src={appLogo} alt="app logo" />
      </Link>
      <div className="right-links">
        <a href="/network">
          <img src={graphImage} alt="network graph" />
        </a>
        <img src={avatarPlaceholder} alt="avatar placeholder" />
      </div>
    </div>
  );
};

export default Header;
