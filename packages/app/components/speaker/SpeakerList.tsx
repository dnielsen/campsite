import React from "react";
import { SpeakerPreview } from "../../common/interfaces";
import { Grid } from "@material-ui/core";
import SpeakerPreviewItem from "./SpeakerPreviewItem";

interface Props {
  speakers: SpeakerPreview[];
}

function SpeakerList(props: Props) {
  return (
    <Grid container justify={"center"} spacing={8}>
      {props.speakers.map((speaker) => (
        <Grid item key={speaker.id}>
          <SpeakerPreviewItem key={speaker.id} speaker={speaker} />
        </Grid>
      ))}
    </Grid>
  );
}

export default SpeakerList;
