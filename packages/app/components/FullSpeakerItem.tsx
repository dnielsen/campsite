import React from "react";
import { Speaker } from "../common/interfaces";
import StyledAvatar from "../styled/StyledAvatar";
import {
  createStyles,
  Link,
  Paper,
  Theme,
  Typography,
} from "@material-ui/core";
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
      "& > *": {
        margin: "2em",
      },
    },
    link: {
      display: "block",
    },
  }),
);

function FullSpeakerItem(props: Props) {
  const classes = useStyles();
  console.log(props.speaker);
  return (
    <Paper style={{ padding: "1em" }} className={classes.root}>
      <div>
        <SpeakerPreviewItem speaker={props.speaker} />
        <Link
          variant={"subtitle2"}
          href={"https://twitter.com/elonmusk"}
          className={classes.link}
        >
          Twitter
        </Link>
        <Link
          href={"https://linkedin.com"}
          className={classes.link}
          variant={"subtitle2"}
        >
          LinkedIn
        </Link>
      </div>
      <Typography>{props.speaker.bio}</Typography>
      {/*{props.speaker.sessions.map((session) => (*/}
      {/*  <div key={session.id}>hello {session.name}</div>*/}
      {/*))}*/}
    </Paper>
  );
}

export default FullSpeakerItem;
