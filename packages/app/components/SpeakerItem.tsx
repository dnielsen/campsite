import React from "react";
import { Person } from "../common/interfaces";

interface Props {
  speaker: Person;
}

function SpeakerItem(props: Props) {
  return (
    <div>
      <img
        height={200}
        src={props.speaker.photo}
        alt={`${props.speaker.name}'s photo`}
      />
      <div>
        <h5>
          <a href={`/profile/${props.speaker.id}`}>{props.speaker.name}</a>
        </h5>
        <span>{props.speaker.headline}</span>
        <p>{props.speaker.bio}</p>
      </div>
    </div>
  );
}

export default SpeakerItem;
