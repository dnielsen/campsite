import React from "react";
import EventForm from "../components/EventForm";
import useCreateEventFormProps from "../hooks/useCreateEventFormProps";

function CreateEvent() {
  const formProps = useCreateEventFormProps();

  return (
    <div>
      <h3>Create an event</h3>
      <EventForm formProps={formProps} />
    </div>
  );
}

export default CreateEvent;
