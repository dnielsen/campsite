import React, { ReactNode } from "react";
import Link from "next/link";
import Head from "next/head";
import {
  AppBar,
  Container,
  CssBaseline,
  IconButton,
  Typography,
} from "@material-ui/core";
import HomeIcon from "@material-ui/icons/Home";

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
    <CssBaseline />
    <header style={{ marginBottom: "2em" }}>
      <AppBar position={"static"}>
        <Link href={"/"}>
          <IconButton edge={"start"} color={"inherit"}>
            <HomeIcon fontSize={"large"} />
          </IconButton>
        </Link>
      </AppBar>
    </header>
    <Container maxWidth="md">{children}</Container>
    <footer style={{ marginTop: "2em" }}>
      <Typography variant="body2" color="textSecondary" align="center">
        {"Copyright © "}
        <Link href="http://localhost:3000">Campsite.org</Link>{" "}
        {new Date().getFullYear()}
        {"."}
      </Typography>
    </footer>
  </div>
);

export default Layout;
