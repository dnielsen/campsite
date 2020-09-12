import React from "react";
import { Session } from "../common/interfaces";
import { Grid, Typography } from "@material-ui/core";
import SpeakerList from "./SpeakerList";

interface Props {
  session: Session;
}

function SessionItem(props: Props) {
  return (
    <Grid container spacing={1}>
      <Grid item xs={4}>
        <SpeakerList speakers={props.session.speakers} />
      </Grid>
      <Grid item xs={8}>
        <div>
          <Typography variant={"h2"}>{props.session.name}</Typography>
          <Typography>
            {new Date(props.session.startDate).toLocaleString()}
          </Typography>
          <Typography>{props.session.description}</Typography>
        </div>
      </Grid>
    </Grid>
  );
}

export default SessionItem;
