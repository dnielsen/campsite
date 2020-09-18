import React from "react";
import { Link } from "react-router-dom";

function Footer() {
  return (
    <footer>
      <p>
        {"Copyright © "}
        <Link to="/">Campsite.org</Link> {new Date().getFullYear()}
        {"."}
      </p>
    </footer>
  );
}

export default Footer;
