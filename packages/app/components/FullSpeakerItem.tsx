import React from "react";
import { Speaker } from "../common/interfaces";
import StyledAvatar from "../styled/StyledAvatar";
import { createStyles, Theme, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";

interface Props {
  person: Speaker;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: "100%",
      display: "flex",
      alignItems: "center",
      justifyContent: "spaceBetween",
      "& > *": {
        margin: "2em",
      },
    },
  }),
);

function FullSpeakerItem(props: Props) {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <div>
        <StyledAvatar
          src={props.person.photo}
          alt={props.person.name}
        />
        <Typography align={"center"}>{props.person.name}</Typography>
        <Typography align={"center"}>{props.person.headline}</Typography>
      </div>
      <Typography>{props.person.bio}</Typography>
    </div>
  );
}

export default FullSpeakerItem;
