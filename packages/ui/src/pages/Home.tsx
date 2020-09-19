import React from "react";
import { Redirect } from "react-router-dom";

const EVENT_ID = "e3a27b7d-b37d-4cd2-b8bd-e5bd5551077c";

// Temporarily the `/` route will redirect to the event which of id is hardcoded above.
function Home() {
  return <Redirect to={`/events/${EVENT_ID}`} />;
}

export default Home;
