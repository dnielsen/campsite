import React from "react";
import { EventDetails } from "../common/interfaces";
import SpeakerList from "./SpeakerList";
import SessionSchedule from "./SessionSchedule";
import { createStyles, Link, Paper, Typography } from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";
import util from "../common/util";

interface Props {
  eventDetails: EventDetails;
}

const useStyles = makeStyles(() =>
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
      marginRight: "0.5em",
      marginBottom: "0.5em",
    },
    container: {
      margin: "1.5em",
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
      {/*For now we'll just use the startDate info*/}
      <div className={classes.info}>
        <Typography align={"right"}>
          When: {util.getFullDate(props.eventDetails.startDate)}
        </Typography>
        <Typography align={"right"}>
          Where: {props.eventDetails.address}
        </Typography>
        <Typography align={"right"}>
          Organizer: {props.eventDetails.organizerName}
        </Typography>
        <Typography align={"right"}>
          <Link href={"/"} variant={"h5"}>
            Register now!
          </Link>
        </Typography>
      </div>
      <img
        src={props.eventDetails.photo}
        alt={props.eventDetails.name}
        className={classes.img}
      />
      <Typography variant={"body1"} className={classes.container}>
        {props.eventDetails.description}
      </Typography>
      <div className={classes.container}>
        <SessionSchedule sessions={props.eventDetails.sessions} />
      </div>
      <Typography align={"center"} variant={"h4"}>
        Our speakers
      </Typography>
      <div className={classes.container}>
        <SpeakerList speakers={props.eventDetails.speakers} />
      </div>
    </Paper>
  );
}

export default EventItem;
