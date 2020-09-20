import React from "react";
import { Redirect } from "react-router-dom";
import useAPI from "../hooks/useAPI";
import { EventDetails } from "../common/interfaces";

function Home() {
  // For dev we're just grabbing whatever event id from the database since there's
  // no <EventList /> for now and redirecting to this event. It's hardcoded in the backend
  // that when it's run, it's creating a sample event.
  // ------
  const { data: events, loading, error } = useAPI<EventDetails[]>("/events");
  if (loading) return <div>loading...</div>;
  if (error) return <div>error: {error.message}</div>;
  if (events.length === 0)
    return (
      <div>
        No events in the database. Restart the server to create a sample event.
      </div>
    );
  // ----
  return <Redirect to={`/events/${events[0].id}`} />;
}

export default Home;
