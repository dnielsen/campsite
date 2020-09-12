import React from "react";
import SpeakerItem from "./SpeakerItem";
import { Speaker } from "../common/interfaces";
import { Grid } from "@material-ui/core";

interface Props {
  speakers: Speaker[];
}

function SpeakerList(props: Props) {
  const speakers = [...props.speakers, ...props.speakers];
  return (
    <Grid container justify={"center"} spacing={8}>
      {speakers.map((speaker) => (
        <Grid item key={speaker.id}>
          <SpeakerItem key={speaker.id} speaker={speaker} />
        </Grid>
      ))}
    </Grid>
  );
}

export default SpeakerList;
