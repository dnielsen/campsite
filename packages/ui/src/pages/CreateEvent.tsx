import React from "react";
import useCreateEventForm from "../hooks/useCreateEventForm";
import EventForm from "../components/EventForm";

function CreateEvent() {
  const { formConfig } = useCreateEventForm();

  return (
    <div>
      <h3>Create an event</h3>
      <EventForm formConfig={formConfig} />
    </div>
  );
}

export default CreateEvent;
