import React from "react";
import useCreateSpeakerForm from "../hooks/useCreateSpeakerForm";
import SpeakerForm from "../components/SpeakerForm";

function CreateSpeaker() {
  const formConfig = useCreateSpeakerForm();

  return (
    <div>
      <h3>Create a speaker</h3>
      <SpeakerForm formConfig={formConfig} />
    </div>
  );
}

export default CreateSpeaker;
