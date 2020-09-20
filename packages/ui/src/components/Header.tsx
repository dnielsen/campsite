import React from "react";
import { Link } from "react-router-dom";
import * as g from "../styled/globalStyles";

function Header() {
  return (
    <g.HeaderWrapperContainer>
      <g.Container>
        <g.HeaderWrapper>
          <g.Logo>
            <img
              src="http://www.campsite.org/bundles/spoutlet/images/logo-campsite.png?v=dev"
              className="img-fluid"
              width="150"
              alt={"Campsite"}
            />
          </g.Logo>
          <g.NavBar>
            <ul>
              <li>
                <Link to={"/"}>Home</Link>
                <Link to={"/speakers"}>All Speakers</Link>
                <Link to={"/events/create"}>Create Event</Link>
                <Link to={"/sessions/create"}>Create Session</Link>
                <Link to={"/speakers/create"}>Create Speaker</Link>
              </li>
            </ul>
          </g.NavBar>
        </g.HeaderWrapper>
      </g.Container>
    </g.HeaderWrapperContainer>
  );
}

export default Header;
