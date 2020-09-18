import React from "react";
import { Route, Switch } from "react-router-dom";
import Home from "./pages/Home";
import FullSpeaker from "./pages/FullSpeaker";
import FullSession from "./pages/FullSession";
import Footer from "./components/Footer";
import Header from "./components/Header";
import CreateSpeaker from "./pages/CreateSpeaker";

function App() {
  return (
    <div>
      <Header />
      <main>
        <Switch>
          <Route path="/speakers/:id">
            <FullSpeaker />
          </Route>
          <Route path="/speakers/create">
            <CreateSpeaker />
          </Route>
          <Route path="/sessions/:id">
            <FullSession />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </main>
      <Footer />
    </div>
  );
}

export default App;
