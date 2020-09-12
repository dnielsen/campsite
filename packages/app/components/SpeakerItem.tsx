import React from "react";
import { Speaker } from "../common/interfaces";
import { Link, Typography } from "@material-ui/core";
import StyledAvatar from "../styled/StyledAvatar";

interface Props {
  speaker: Speaker;
}

function SpeakerItem(props: Props) {
  return (
    <div>
      <Link href={`/speakers/${props.speaker.id}`}>
        <StyledAvatar alt={props.speaker.name} src={props.speaker.photo} />
        <div>
          <Typography align={"center"}>{props.speaker.name}</Typography>
          <Typography align={"center"}>{props.speaker.headline}</Typography>
        </div>
      </Link>
    </div>
  );
}

export default SpeakerItem;
