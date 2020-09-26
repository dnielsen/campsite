// import styled component
import { createGlobalStyle } from "styled-components";
import styled from "styled-components";

export const GlobalStyle = createGlobalStyle`
  body {
    font-family: 'Montserrat', sans-serif;
    margin: 0;
    background: #ffffff;
  }
`;

export const Container = styled.div`
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 100px;

  @media (max-width: 767px) {
    padding: 0px 35px;
  }
`;

export const HeaderWrapperContainer = styled.div``;

export const Nav = styled.nav`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  padding: 15px 0px;
  align-items: center;
  border-bottom: 1px solid #ccc;

  @media (max-width: 960px) {
    display: none;

    .toggle {
      display: block;
    }
  }

  ul {
    display: flex;
    list-style: none;
    padding: 0;
    margin: 0;
    justify-content: end;
  }

  li {
    margin-right: 20px;
  }

  a {
    text-decoration: none;
    color: #777777;
    font-weight: 500;
  }
`;

export const Toggle = styled.div`
  display: none;

  @media (max-width: 960px) {
    display: block;
    text-align: right;
    font-size: 25px;
    padding-top: 20px;
  }
`;

export const NavMobile = styled.nav`
  @media (max-width: 960px) {
    display: block;
    width: 100%;
    background: #fff;
    position: relative;
    -webkit-animation: slide-down 0.3s ease-out;
    -moz-animation: slide-down 0.3s ease-out;

    @-webkit-keyframes slide-down {
      0% {
        opacity: 0;
        -webkit-transform: translateY(0);
      }
      100% {
        opacity: 1;
        -webkit-transform: translateY(0);
      }
    }
    @-moz-keyframes slide-down {
      0% {
        opacity: 0;
        -moz-transform: translateY(-100%);
      }
      100% {
        opacity: 1;
        -moz-transform: translateY(0);
      }
    }

    ul {
      list-style: none;
      padding: 0;
      margin: 0;
      text-align: center;
    }

    li {
      padding: 12px;
      border-bottom: 1px solid #ccc;
    }

    a {
      text-decoration: none;
      color: #777777;
      font-weight: 500;
    }
  }
`;

export const Logo = styled.div`
  justify-self: start;
`;
