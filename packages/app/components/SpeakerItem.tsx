import React from "react";
import { Speaker } from "../common/interfaces";
import { Avatar, createStyles, Theme } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";

interface Props {
  speaker: Speaker;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    avatar: {
      width: "8em",
      height: "8em",
    },
  }),
);

function SpeakerItem(props: Props) {
  const classes = useStyles();
  return (
    <div>
      <Avatar
        alt={props.speaker.name}
        src={props.speaker.photo}
        className={classes.avatar}
      />
      <div>
        <h5>
          <a href={`/profile/${props.speaker.id}`}>{props.speaker.name}</a>
        </h5>
        <span>{props.speaker.headline}</span>
        <p>{props.speaker.bio}</p>
      </div>
    </div>
  );
}

export default SpeakerItem;
