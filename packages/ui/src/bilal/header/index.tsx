import React from "react";

// import styled component
import * as g from "../styles/globalStyles";

const Header = () => {
  return (
    <g.HeaderWrapperContainer>
      <g.Container>
        <g.HeaderWrapper>
          <g.NavBar>
            <ul>
              <li>
                <a href="#">Home</a>
              </li>
              <li>
                <a href="#">Events</a>
              </li>
              <li>
                <a href="#">Sessions</a>
              </li>
              <li>
                <a href="#">Speaker</a>
              </li>
            </ul>
          </g.NavBar>
          <g.Logo>
            <img
              src="http://www.campsite.org/bundles/spoutlet/images/logo-campsite.png?v=dev"
              className="img-fluid"
              width="150"
            />
          </g.Logo>
        </g.HeaderWrapper>
      </g.Container>
    </g.HeaderWrapperContainer>
  );
};

export default Header;
