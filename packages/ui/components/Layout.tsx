import React, { ReactNode } from "react";
import Link from "next/link";
import Head from "next/head";

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
    </Head>
    <header style={{ marginBottom: "2em" }}>
      <nav>
        <ul>
          <li>
            <Link href="/">Home</Link>
          </li>
        </ul>
      </nav>
    </header>
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
