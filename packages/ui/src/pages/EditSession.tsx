import React from "react";
import { useParams } from "react-router-dom";
import SessionForm from "../components/SessionForm";
import useEditSessionForm from "../hooks/useEditSession";

function EditSession(): React.ReactElement {
  const { id } = useParams<{ id: string }>();
  const { formConfig, loading, error } = useEditSessionForm(id);

  if (loading) return <div>loading...</div>;
  if (error) return <div>error: {error.message}</div>;

  return (
    <div>
      <h3>Edit the session</h3>
      <SessionForm formConfig={formConfig} />
    </div>
  );
}

export default EditSession;
