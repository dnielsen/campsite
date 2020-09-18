import React from "react";
import { Link } from "react-router-dom";

function Header() {
  return (
    <header>
      <nav>
        <ul>
          <li>
            <Link to={"/"}>Home</Link>
            <Link to={"/speakers"}>All Speakers</Link>
            <Link to={"/speakers/create"}>Create Speaker</Link>
            <Link to={"/create"}>Create Event</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
}

export default Header;
