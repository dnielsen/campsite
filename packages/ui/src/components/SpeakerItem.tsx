import React, { useState } from "react";
import { SpeakerPreview } from "../common/interfaces";
import { Link } from "react-router-dom";

interface Props {
  speaker: SpeakerPreview;
}

function SpeakerItem(props: Props) {
  return (
    <Link to={`/speakers/${props.speaker.id}`}>
      <img
        src={props.speaker.photo}
        alt={props.speaker.name}
        // We're setting the height temporarily until there's css
        height={200}
      />
      <h5>{props.speaker.name}</h5>
      <h6>{props.speaker.headline}</h6>
    </Link>
  );
}

export default SpeakerItem;
