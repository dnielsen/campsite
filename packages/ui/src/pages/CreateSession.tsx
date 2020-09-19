import React from "react";
import useCreateSessionForm from "../hooks/useCreateSessionForm";
import SessionForm from "../components/SessionForm";

function CreateSession() {
  const { formConfig } = useCreateSessionForm();

  return (
    <div>
      <h3>Create a session</h3>
      <SessionForm formConfig={formConfig} />
    </div>
  );
}

export default CreateSession;
