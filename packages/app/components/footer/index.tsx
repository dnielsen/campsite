import React from "react";

// import bootstrap container
import Container from "react-bootstrap/Container";

const Footer = () => {
  return (
    <Container>
      <footer>
        <p>
          {"Copyright © "}
          <a href="http://localhost:3000">Campsite.org</a>{" "}
          {new Date().getFullYear()}
          {"."}
        </p>
      </footer>
    </Container>
  );
};

export default Footer;
