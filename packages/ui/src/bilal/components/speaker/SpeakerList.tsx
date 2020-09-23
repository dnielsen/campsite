import React from "react";
import { SpeakerPreview } from "../../../common/interfaces";
import SpeakerPreviewItem from "./SpeakerPreviewItem";

interface Props {
  speakers: SpeakerPreview[];
}

function SpeakerList(props: Props) {
  return (
    <div>
      {props.speakers.map((speaker) => (
        <div key={speaker.id}>
          <SpeakerPreviewItem key={speaker.id} speaker={speaker} />
        </div>
      ))}
    </div>
  );
}

export default SpeakerList;
