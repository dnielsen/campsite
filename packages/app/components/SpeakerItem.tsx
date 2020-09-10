import React from "react";
import { Speaker } from "../interfaces";

interface Props {
  speaker: Speaker;
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
          <a
            href={
              // For example:
              // name: John Doe
              // href: /speaker/john-doe
              `/speaker/${props.speaker.name
                .split(" ")
                .map((namePart) => namePart.toLowerCase())
                .join("-")}`
            }
          >
            {props.speaker.name}
          </a>
        </h5>
        <span>{props.speaker.headline}</span>
        <p>{props.speaker.bio}</p>
      </div>
    </div>
  );
}

export default SpeakerItem;
