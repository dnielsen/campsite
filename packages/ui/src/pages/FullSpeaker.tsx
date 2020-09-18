import React from "react";
import useAPI from "../hooks/useAPI";
import { Speaker } from "../common/interfaces";
import SpeakerSessionSchedule from "../components/SpeakerSessionSchedule";
import { Link, useParams } from "react-router-dom";
import SpeakerItem from "../components/SpeakerItem";

function FullSpeaker() {
  const { id } = useParams<{ id: string }>();

  const { data: speaker, loading, error } = useAPI<Speaker>(`/speakers/${id}`);

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  return (
    <div>
      <div>
        <div>
          <SpeakerItem speaker={speaker} />
          <a href={"https://twitter.com/elonmusk"}>Twitter</a>
          <a href={"https://linkedin.com"}>LinkedIn</a>
        </div>
        <p>{speaker.bio}</p>
      </div>
      <SpeakerSessionSchedule sessions={speaker.sessions} />
    </div>
  );
}

export default FullSpeaker;
