import React from "react";

// import bootstrap container
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";

const Header = () => {
  return (
    <Container>
      <Navbar expand="md">
        <Navbar.Brand href="/">LOGO HERE </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="ml-auto">
            <Nav.Link href="/">Home</Nav.Link>
            <Nav.Link href="/">Events</Nav.Link>
            <Nav.Link href="/">Sessions</Nav.Link>
            <Nav.Link href="/">Speakers</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    </Container>
  );
};

export default Header;
