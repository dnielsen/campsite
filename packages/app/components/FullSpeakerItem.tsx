import React from "react";
import { Speaker } from "../common/interfaces";
import {
  createStyles,
  Link,
  List,
  ListItem,
  ListItemText,
  Paper,
  Typography,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";
import SpeakerPreviewItem from "./SpeakerPreviewItem";
import moment from "moment";

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
    },
    link: {
      display: "block",
    },
    flex: {
      display: "flex",
    },
  }),
);

function FullSpeakerItem(props: Props) {
  const classes = useStyles();
  console.log();
  return (
    <Paper style={{ padding: "1em" }} className={classes.root}>
      <div className={classes.flex}>
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
      <List>
        {props.speaker.sessions.map((session) => (
          <Paper key={session.id}>
            <ListItem>
              <ListItemText
                primary={
                  <Link href={`/sessions/${session.id}`}>{session.name}</Link>
                }
                secondary={
                  <Typography variant={"subtitle2"}>
                    {moment(session.startDate).format("MM/DD/YYYY")}
                  </Typography>
                }
              />
            </ListItem>
          </Paper>
        ))}
      </List>
    </Paper>
  );
}

export default FullSpeakerItem;
