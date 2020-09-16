import React from "react";
import { SpeakerPreview } from "../../common/interfaces";
import StyledAvatar from "../../styled/StyledAvatar";
import { createStyles, Link, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";

interface Props {
  speaker: SpeakerPreview;
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

function SpeakerPreviewItem(props: Props) {
  const classes = useStyles();

  return (
    <Link
      href={`/speakers/${props.speaker.id}`}
      className={classes.link}
      variant={"body2"}
      color={"primary"}
    >
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
    </Link>
  );
}

export default SpeakerPreviewItem;
