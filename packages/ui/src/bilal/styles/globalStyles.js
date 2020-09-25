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
    padding: 50px;
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

export const Logo = styled.div`
  justify-self: start;
`;
