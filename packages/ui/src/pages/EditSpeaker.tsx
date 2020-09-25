import React from "react";
import { useParams } from "react-router-dom";
import useEditSpeakerFormProps from "../hooks/useEditSpeakerFormProps";
import SpeakerForm from "../components/SpeakerForm";

function EditSpeaker() {
  const { id } = useParams<{ id: string }>();
  const formProps = useEditSpeakerFormProps({ id });
  return (
    <div>
      <h3>Edit speaker</h3>
      <SpeakerForm formProps={formProps} />
    </div>
  );
}

export default EditSpeaker;
