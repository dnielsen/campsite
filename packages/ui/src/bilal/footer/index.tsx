import React from "react";

// import styled component
import * as g from "../styles/globalStyles";

const Footer = () => {
  return (
    <g.Container>
      <footer>
        <p>
          {"Copyright © "}
          <a href="http://localhost:3000">Campsite.org</a>{" "}
          {new Date().getFullYear()}
          {"."}
        </p>
      </footer>
    </g.Container>
  );
};

export default Footer;
