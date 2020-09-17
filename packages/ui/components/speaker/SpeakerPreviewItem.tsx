import React from "react";
import { SpeakerPreview } from "../../common/interfaces";

interface Props {
  speaker: SpeakerPreview;
}

function SpeakerPreviewItem(props: Props) {
  return (
    <a href={`/speakers/${props.speaker.id}`}>
      <div>
        <img src={props.speaker.photo} alt={props.speaker.name} />
        <p>{props.speaker.name}</p>
        <p>{props.speaker.headline}</p>
      </div>
    </a>
  );
}

export default SpeakerPreviewItem;
