import React from "react";
import SpeakerItem from "./SpeakerItem";
import { people } from "../common/devData";

function SpeakerList() {
  return (
    <div>
      {people.map((speaker) => (
        <SpeakerItem key={speaker.id} speaker={speaker} />
      ))}
    </div>
  );
}

export default SpeakerList;
