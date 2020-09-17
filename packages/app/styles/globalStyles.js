// import styled component
import { createGlobalStyle } from "styled-components";
import styled from "styled-components";

export const GlobalStyle = createGlobalStyle`
  body {
    font-family: 'Montserrat', sans-serif;
    margin: 0;
  }
`;

export const Container = styled.div`
  width: 100%;
  max-width: 1140px;
  margin: 0 auto;
`;

export const HeaderWrapper = styled.header`
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  padding: 15px 0px;
  align-items: center;
`;

export const Logo = styled.div``;

export const NavBar = styled.div`
  ul {
    display: flex;
    list-style: none;
    padding: 0;
    margin: 0;
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
