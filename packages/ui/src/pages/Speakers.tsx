import React from "react";
import useAPI from "../hooks/useAPI";
import SpeakerList from "../components/SpeakerList";
import { Speaker } from "../common/interfaces";

function Speakers() {
  const { data: speakers, loading, error } = useAPI<Speaker[]>("/speakers");

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  return (
    <div>
      <SpeakerList speakers={speakers} />
    </div>
  );
}

export default Speakers;
