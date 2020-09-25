import React from "react";
import { useParams } from "react-router-dom";
import EventForm from "../components/EventForm";
import useEditEventFormProps from "../hooks/useEditEventFormProps";

function EditEvent() {
  const { id } = useParams<{ id: string }>();
  const formProps = useEditEventFormProps({ id });

  return (
    <div>
      <h3>Edit Event</h3>
      <EventForm formProps={formProps} />
    </div>
  );
}

export default EditEvent;
