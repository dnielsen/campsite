import React from "react";
import { Speaker } from "../../common/interfaces";
import { createStyles, Link, Paper, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";
import SpeakerPreviewItem from "./SpeakerPreviewItem";
import SpeakerSessionSchedule from "../SpeakerSessionSchedule";

interface Props {
  speaker: Speaker;
}

const useStyles = makeStyles(() =>
  createStyles({
    root: {
      width: "100%",
      alignItems: "center",
      "& > *": {
        margin: "2em",
      },
      padding: "1em",
    },
    link: {
      display: "block",
    },
    speakerInfoWrapper: {
      display: "flex",
      marginBottom: "2em",
    },
    sessions: {
      display: "flex",
      justifyContent: "spaceBetween",
    },
  }),
);

function SpeakerItem(props: Props) {
  const classes = useStyles();
  return (
    <Paper className={classes.root}>
      <div>
        <div className={classes.speakerInfoWrapper}>
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
          <Typography align={"center"}>{props.speaker.bio}</Typography>
        </div>
        <SpeakerSessionSchedule sessions={props.speaker.sessions} />
      </div>
    </Paper>
  );
}

export default SpeakerItem;
