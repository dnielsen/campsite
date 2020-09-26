import React, { Component, Fragment } from "react";

// import styled component
import * as g from "../../styles/globalStyles";

import MenuItem from "./menuItem";

class Header extends Component {
  state = {
    mobileMenu: false,
  };

  handleToggleChange = () => {
    this.setState({
      mobileMenu: !this.state.mobileMenu,
    });
  };

  render() {
    const { mobileMenu } = this.state;
    return (
      <g.Container>
        <g.Nav>
          <g.Logo>
            <img
              src="http://www.campsite.org/bundles/spoutlet/images/logo-campsite.png?v=dev"
              className="img-fluid"
              width="150"
            />
          </g.Logo>
          <ul>
            {MenuItem.map((item, index) => {
              return (
                <li key={index}>
                  <a href={item.url}>{item.title}</a>
                </li>
              );
            })}
          </ul>
        </g.Nav>
        <g.Toggle onClick={this.handleToggleChange}>
          {" "}
          <i
            className={mobileMenu ? "fa fa-times" : "fa fa-bars"}
            aria-hidden="true"
          ></i>
        </g.Toggle>
        {mobileMenu && (
          <g.NavMobile>
            <g.Logo>
              <img
                src="http://www.campsite.org/bundles/spoutlet/images/logo-campsite.png?v=dev"
                className="img-fluid"
                width="150"
              />
            </g.Logo>
            <ul>
              {MenuItem.map((item, index) => {
                return (
                  <li key={index}>
                    <a href={item.url}>{item.title}</a>
                  </li>
                );
              })}
            </ul>
          </g.NavMobile>
        )}
      </g.Container>
    );
  }
}

export default Header;
