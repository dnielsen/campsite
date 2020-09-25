import React from "react";
import SpeakerForm from "../components/SpeakerForm";
import useCreateSpeakerFormProps from "../hooks/useCreateSpeakerFormProps";

function CreateSpeaker() {
  const formProps = useCreateSpeakerFormProps();

  return (
    <div>
      <h3>Create a speaker</h3>
      <SpeakerForm formProps={formProps} />
    </div>
  );
}

export default CreateSpeaker;
