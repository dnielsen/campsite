import React from "react";
import { Route, Switch } from "react-router-dom";
import FullSpeaker from "./pages/FullSpeaker";
import FullSession from "./pages/FullSession";
import Speakers from "./pages/Speakers";
import CreateSpeaker from "./pages/CreateSpeaker";
import CreateSession from "./pages/CreateSession";
import CreateEvent from "./pages/CreateEvent";
import Home from "./pages/Home";
import FullEvent from "./pages/FullEvent";
import Header from "./bilal/components/header";
import Footer from "./bilal/components/footer";
import Events from "./pages/Events";

function App() {
  return (
    <div>
      <Header />
      {/*<Header />*/}
      <main>
        <Switch>
          <Route exact path="/speakers">
            <Speakers />
          </Route>
          <Route exact path="/events">
            <Events />
          </Route>
          <Route exact path="/speakers/create">
            <CreateSpeaker />
          </Route>
          <Route exact path="/speakers/:id">
            <FullSpeaker />
          </Route>
          <Route exact path="/sessions/create">
            <CreateSession />
          </Route>
          <Route exact path="/sessions/:id">
            <FullSession />
          </Route>
          <Route exact path="/events/create">
            <CreateEvent />
          </Route>
          <Route exact path="/events/:id">
            <FullEvent />
          </Route>
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
