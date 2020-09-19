import React, { useState } from "react";
import { SessionPreview } from "../../common/interfaces";
import useAPI from "../../hooks/useAPI";
import Select from "react-select";

interface Option {
  value: string;
  label: string;
}

function SessionInput() {
  // For now we'll fetch all sessions, later we might fetch only those
  // that haven't been assigned to any event id.
  // const [selectedSessions, setSelectedSessions] = useState<string[]>([]);
  // const { data: sessions, loading, error } = useAPI<SessionPreview[]>(
  //   "/sessions",
  // );
  //
  // if (loading) return <div>loading...</div>;
  // if (error) return <div>error: {error.message}</div>;
  //
  // const options: Option[] = sessions.map((session) => ({
  //   label: session.name,
  //   value: session.id,
  // }));

  return (
    <div>{/*<Select isMulti name={"sessions"} options={options} />*/}</div>
  );
}

export default SessionInput;
