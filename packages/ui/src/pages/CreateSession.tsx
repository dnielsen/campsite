import React from "react";
import SessionForm from "./createSession/SessionForm";
import useAPI from "../hooks/useAPI";
import { EventDetails, SpeakerPreview } from "../common/interfaces";

function CreateSession() {
  // A temporary solution, later we might load just the speaker/event ids and names,
  // and do it asynchronously - after having loaded the rest of the form.
  const {
    data: speakers,
    loading: speakersLoading,
    error: speakersError,
  } = useAPI<SpeakerPreview[]>("/speakers");
  const { data: events, loading: eventsLoading, error: eventsError } = useAPI<
    EventDetails[]
  >("/events");

  if (speakersLoading || eventsLoading) return <div>loading...</div>;
  if (speakersError) return <div>error: {speakersError.message}</div>;
  if (eventsError) return <div>error: {eventsError.message}</div>;

  if (events.length === 0) return <div>Create an event first.</div>;
  if (speakers.length === 0) return <div>Create a speaker first.</div>;

  return (
    <div>
      <h3>Create a session</h3>
      <SessionForm speakers={speakers} events={events} />
    </div>
  );
}

export default CreateSession;
