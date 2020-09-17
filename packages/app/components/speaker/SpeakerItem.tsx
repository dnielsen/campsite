import React from "react";
import { Speaker } from "../../common/interfaces";
import SpeakerPreviewItem from "./SpeakerPreviewItem";
import SpeakerSessionSchedule from "../SpeakerSessionSchedule";

interface Props {
  speaker: Speaker;
}

function SpeakerItem(props: Props) {
  return (
    <div>
      <div>
        <div>
          <SpeakerPreviewItem speaker={props.speaker} />
        </div>
      </div>
      <SpeakerSessionSchedule sessions={props.speaker.sessions} />
    </div>
  );
}

export default SpeakerItem;
