import React from "react";
import { useParams } from "react-router-dom";
import SpeakerForm from "../components/SpeakerForm";
import useEditSpeakerForm from "../hooks/useEditSpeakerForm";

function EditSpeaker(): React.ReactElement {
  const { id } = useParams<{ id: string }>();
  const { formConfig, loading, error } = useEditSpeakerForm(id);

  if (loading) return <div>loading...</div>;
  if (error) return <div>error: {error.message}</div>;

  return (
    <div>
      <h3>Edit the speaker</h3>
      <SpeakerForm formConfig={formConfig} />
    </div>
  );
}

export default EditSpeaker;
