import React from "react";
import { Link } from "react-router-dom";
import * as g from "../styled/globalStyles";

function Footer() {
  return (
    <g.Container>
      <footer>
        <p>
          {"Copyright © "}
          <Link to="/">Campsite.org</Link> {new Date().getFullYear()}
          {"."}
        </p>
      </footer>
    </g.Container>
  );
}

export default Footer;
