import React from "react";
import { Speaker } from "../common/interfaces";
import StyledAvatar from "../styled/StyledAvatar";
import { Typography } from "@material-ui/core";

interface Props {
  speaker: Speaker;
}

function SpeakerPreviewItem(props: Props) {
  return (
    <div>
      <StyledAvatar src={props.speaker.photo} alt={props.speaker.name} />
      <Typography align={"center"} variant={"h6"}>
        {props.speaker.name}
      </Typography>
      <Typography
        align={"center"}
        variant={"caption"}
        component={"h6"}
        gutterBottom
      >
        {props.speaker.headline}
      </Typography>
    </div>
  );
}

export default SpeakerPreviewItem;
