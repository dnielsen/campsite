import React from "react";
import { Route, Switch } from "react-router-dom";
import Home from "./pages/Home";
import FullSpeaker from "./pages/FullSpeaker";
import FullSession from "./pages/FullSession";
import Footer from "./components/Footer";
import Header from "./components/Header";
import Speakers from "./pages/Speakers";
import CreateSpeaker from "./pages/CreateSpeaker";
import CreateSession from "./pages/CreateSession";

function App() {
  return (
    <div>
      <Header />
      <main>
        <Switch>
          <Route exact path="/speakers">
            <Speakers />
          </Route>
          <Route exact path="/speakers/create">
            <CreateSpeaker />
          </Route>
          <Route path="/speakers/:id">
            <FullSpeaker />
          </Route>
          <Route exact path="/sessions/create">
            <CreateSession />
          </Route>
          <Route path="/sessions/:id">
            <FullSession />
          </Route>
          {/*<Route exact path="/create">*/}
          {/*  <CreateEvent />*/}
          {/*</Route>*/}
          <Route exact path="/">
            <Home />
          </Route>
        </Switch>
      </main>
      <Footer />
    </div>
  );
}

export default App;
