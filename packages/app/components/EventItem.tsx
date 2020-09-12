import React from "react";
import { EventDetails } from "../common/interfaces";
import SpeakerList from "./SpeakerList";
import SessionSchedule from "./SessionSchedule";
import { createStyles, Paper, Theme, Typography } from "@material-ui/core";
import moment from "moment";
import { makeStyles } from "@material-ui/styles";

interface Props {
  eventDetails: EventDetails;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    img: {
      width: "100%",
      height: "300px",
      objectFit: "cover",
    },
    paper: {
      padding: "2rem 0",
    },
    info: {
      margin: "100px",
    },
  }),
);

function EventItem(props: Props) {
  const classes = useStyles();
  return (
    <Paper className={classes.paper}>
      <Typography align={"center"} variant={"h3"}>
        {props.eventDetails.name}
      </Typography>
      {/*<Paper>*/}
      {/*For now we'll just use the startDate info*/}
      <div className={"info"}>
        <Typography align={"right"}>
          When: {moment(props.eventDetails.startDate).format("MM/DD/YYYY")}
        </Typography>
        <Typography align={"right"}>
          Organizer: {props.eventDetails.organizerName}
        </Typography>
      </div>
      <img
        src={props.eventDetails.photo}
        alt={props.eventDetails.name}
        className={classes.img}
      />
      <p>{props.eventDetails.description}</p>
      <SessionSchedule sessions={props.eventDetails.sessions} />
      <SpeakerList speakers={props.eventDetails.speakers} />
    </Paper>
  );
}

export default EventItem;
