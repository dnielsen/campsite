import React from "react";
import { SpeakerPreview } from "../common/interfaces";
import { Link } from "react-router-dom";

interface Props {
  speaker: SpeakerPreview;
}

function SpeakerItem(props: Props) {
  return (
    <Link to={`/speakers/${props.speaker.id}`}>
      <img src={props.speaker.photo} alt={props.speaker.name} />
      <h5>{props.speaker.name}</h5>
      <h6>{props.speaker.headline}</h6>
    </Link>
  );
}

export default SpeakerItem;
