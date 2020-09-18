import React from "react";
import { SpeakerPreview } from "../common/interfaces";
import SpeakerItem from "./SpeakerItem";

interface Props {
  speakers: SpeakerPreview[];
}

function Speakers(props: Props) {
  return (
    <div>
      <h4>Our speakers</h4>
      <ul>
        {props.speakers.map((speaker) => (
          <li key={speaker.id}>
            <SpeakerItem speaker={speaker} />
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Speakers;
