import React from "react";
import { Redirect } from "react-router-dom";

const EVENT_ID = "80e211fd-d33a-4149-8d8b-47ff5e4c8f0f";

// Temporarily the `/` route will redirect to the event which of id is hardcoded above.
function Home() {
  return <Redirect to={`/events/${EVENT_ID}`} />;
}

export default Home;
