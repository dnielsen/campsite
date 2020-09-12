import React from "react";
import { Speaker } from "../common/interfaces";
import StyledAvatar from "../styled/StyledAvatar";
import { createStyles, Paper, Theme, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";
import SpeakerPreviewItem from "./SpeakerPreviewItem";

interface Props {
  speaker: Speaker;
}

const useStyles = makeStyles(() =>
  createStyles({
    root: {
      width: "100%",
      display: "flex",
      alignItems: "center",
      justifyContent: "spaceBetween",
      // flexDirection: "column",
      "& > *": {
        margin: "2em",
      },
    },
  }),
);

function FullSpeakerItem(props: Props) {
  const classes = useStyles();
  return (
    <Paper style={{ padding: "1em" }} className={classes.root}>
      <SpeakerPreviewItem speaker={props.speaker} />
      <Typography>{props.speaker.bio}</Typography>
    </Paper>
  );
}

export default FullSpeakerItem;
