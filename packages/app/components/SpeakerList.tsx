import React from "react";
import SpeakerItem from "./SpeakerItem";
import devData from "../common/devData";
import { Speaker } from "../common/interfaces";

interface Props {
  speakers: Speaker[];
}

function SpeakerList(props: Props) {
  return (
    <div>
      {props.speakers.map((speaker) => (
        <SpeakerItem key={speaker.id} speaker={speaker} />
      ))}
    </div>
  );
}

export default SpeakerList;
