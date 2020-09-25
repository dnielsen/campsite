import React from "react";
import { Speaker } from "../../common/interfaces";
import SpeakerPreviewItem from "./SpeakerPreviewItem";
import SpeakerSessionSchedule from "../SpeakerSessionSchedule";

interface Props {
  speaker: any;
}

function SpeakerItem(props: Props) {
  return (
    <div>
      <SpeakerPreviewItem speaker={props.speaker} />
      {/* <SpeakerSessionSchedule sessions={props.speaker.sessions} /> */}
    </div>
  );
}

export default SpeakerItem;
