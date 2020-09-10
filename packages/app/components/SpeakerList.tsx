import React from "react";
import SpeakerItem from "./SpeakerItem";
import devData from "../common/devData";

function SpeakerList() {
  return (
    <div>
      {devData.people.map((person) => (
        <SpeakerItem key={person.id} speaker={person} />
      ))}
    </div>
  );
}

export default SpeakerList;
