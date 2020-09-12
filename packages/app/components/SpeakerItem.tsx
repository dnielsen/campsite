import React from "react";
import { Speaker } from "../common/interfaces";
import { createStyles, Link, Paper, Typography } from "@material-ui/core";
import StyledAvatar from "../styled/StyledAvatar";
import { makeStyles } from "@material-ui/styles";
import { display } from "@material-ui/system";
import SpeakerPreviewItem from "./SpeakerPreviewItem";

interface Props {
  speaker: Speaker;
}

const useStyles = makeStyles(() =>
  createStyles({
    link: {
      width: "100%",
      display: "flex",
      alignItems: "center",
      flexDirection: "column",
      justifySelf: "center",
    },
  }),
);

function SpeakerItem(props: Props) {
  const classes = useStyles();
  return (
    <Link
      href={`/speakers/${props.speaker.id}`}
      className={classes.link}
      variant={"body2"}
      color={"primary"}
    >
      {/*<StyledAvatar alt={props.speaker.name} src={props.speaker.photo} />*/}
      {/*<div>*/}
      {/*  <Typography align={"center"}>{props.speaker.name}</Typography>*/}
      {/*  <Typography align={"center"}>{props.speaker.headline}</Typography>*/}
      {/*</div>*/}
      <SpeakerPreviewItem speaker={props.speaker} />
    </Link>
  );
}

export default SpeakerItem;
