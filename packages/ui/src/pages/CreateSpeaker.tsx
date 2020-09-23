import React from "react";
import SpeakerForm from "./createSpeaker/SpeakerForm";
import useSpeakerForm from "../hooks/useSpeakerForm";

function CreateSpeaker() {
  const { formConfig } = useSpeakerForm();

  return (
    <div>
      <h3>Create a speaker</h3>
      <SpeakerForm formConfig={formConfig} />
    </div>
  );
}

export default CreateSpeaker;
