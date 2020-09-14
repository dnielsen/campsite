import React from "react";
import { Speaker } from "../common/interfaces";
import { createStyles, Link } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";
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
      <SpeakerPreviewItem speaker={props.speaker} />
    </Link>
  );
}

export default SpeakerItem;
