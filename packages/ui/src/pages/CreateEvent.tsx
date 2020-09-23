import React from "react";
import useEventForm from "../hooks/useEventForm";
import EventForm from "./createEvent/EventForm";

function CreateEvent() {
  const { formConfig } = useEventForm();

  return (
    <div>
      <h3>Create an event</h3>
      <EventForm formConfig={formConfig} />
    </div>
  );
}

export default CreateEvent;
