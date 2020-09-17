import React, { ReactNode } from "react";
import Link from "next/link";
import Head from "next/head";

// import styled component
import * as g from "../styles/globalStyles";

// import components
import Header from "./header";

type Props = {
  children?: any;
  title?: string;
};

const Layout = ({ children, title = "This is the default title" }: Props) => (
  <div>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      <link
        rel="stylesheet"
        type="text/css"
        href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
      />
    </Head>
    <Header />
    {children}
    <footer>
      <p>
        {"Copyright © "}
        <a href="http://localhost:3000">Campsite.org</a>{" "}
        {new Date().getFullYear()}
        {"."}
      </p>
    </footer>
  </div>
);

export default Layout;
