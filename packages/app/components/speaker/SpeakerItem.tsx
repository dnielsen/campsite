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
          <div>
            <SpeakerPreviewItem speaker={props.speaker} />
            <a href={"https://twitter.com/elonmusk"}>Twitter</a>
            <a href={"https://linkedin.com"}>LinkedIn</a>
          </div>
          <p>{props.speaker.bio}</p>
        </div>
        <SpeakerSessionSchedule sessions={props.speaker.sessions} />
      </div>
    </div>
  );
}

export default SpeakerItem;
