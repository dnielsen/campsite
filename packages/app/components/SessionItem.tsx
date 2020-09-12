import React from "react";
import { Session } from "../common/interfaces";
import { Grid, Paper, Typography } from "@material-ui/core";
import SpeakerList from "./SpeakerList";

interface Props {
  session: Session;
}

function SessionItem(props: Props) {
  return (
    <Paper style={{ padding: "1em" }}>
      <Grid container spacing={1}>
        <Grid item xs={4}>
          <SpeakerList speakers={props.session.speakers} />
        </Grid>
        <Grid item xs={8}>
          <div>
            <Typography variant={"h2"} gutterBottom>
              {props.session.name}
            </Typography>
            <Typography variant={"subtitle2"} gutterBottom>
              Starting at {new Date(props.session.startDate).toLocaleString()}
            </Typography>
            <Typography gutterBottom>{props.session.description}</Typography>
          </div>
        </Grid>
      </Grid>
    </Paper>
  );
}

export default SessionItem;
