import React, { ReactNode } from "react";
import Head from "next/head";

// import components
import Header from "./header";
import Footer from "./footer";

import * as g from "../styles/globalStyles";

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
        href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
      />
      <link
        rel="stylesheet"
        type="text/css"
        href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
      />
      <link
        href="https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500;600;700;800&display=swap"
        rel="stylesheet"
      />
    </Head>
    <g.GlobalStyle />
    <Header />
    {children}
    <Footer />
  </div>
);

export default Layout;
