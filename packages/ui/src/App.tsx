import React from "react";
import { Route, Switch } from "react-router-dom";
import Home from "./pages/Home";
import FullSpeaker from "./pages/FullSpeaker";
import FullSession from "./pages/FullSession";
import Footer from "./components/Footer";
import Header from "./components/Header";
import Speakers from "./pages/Speakers";
import CreateEvent from "./pages/CreateEvent";
import CreateSpeaker from "./pages/CreateSpeaker";

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
