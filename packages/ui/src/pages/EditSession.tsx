import React from "react";
import { useParams } from "react-router-dom";
import SessionForm from "../components/SessionForm";
import useEditSessionFormProps from "../hooks/useEditSessionFormProps";
import useAPI from "../hooks/useAPI";
import { EventDetails, SpeakerPreview } from "../common/interfaces";

function EditSession() {
  const { id } = useParams<{ id: string }>();
  const formProps = useEditSessionFormProps({ id });
  // A temporary solution, later we might load just the speaker/event ids and names,
  // and do it asynchronously, that is after having loaded the rest of the form.
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
  return (
    <div>
      <h3>Edit Session</h3>
      <SessionForm speakers={speakers} events={events} formProps={formProps} />
    </div>
  );
}

export default EditSession;
