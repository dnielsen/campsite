import React from "react";
import { Session } from "../../common/interfaces";
import { Grid, Link, Paper, Typography } from "@material-ui/core";
import SpeakerList from "../speaker/SpeakerList";
import moment from "moment";
import util from "../../common/util";

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
            <Typography variant={"h2"}>{props.session.name}</Typography>
            <Typography variant={"subtitle2"}>
              {util.getHourRangeString(
                props.session.startDate,
                props.session.endDate,
              )}{" "}
              on {moment(props.session.startDate).format("MM/DD/YYYY")}
            </Typography>
            <Typography
              component={Link}
              href={props.session.url}
              variant={"subtitle1"}
            >
              {props.session.url}
            </Typography>
            <Typography gutterBottom>{props.session.description}</Typography>
          </div>
        </Grid>
      </Grid>
    </Paper>
  );
}

export default SessionItem;
