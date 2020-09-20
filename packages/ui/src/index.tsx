import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import * as g from "./styled/globalStyles";

ReactDOM.render(
  <BrowserRouter>
    <g.GlobalStyle />
    <App />
  </BrowserRouter>,
  document.getElementById("root"),
);
