import React from "react";
import {
  Session,
  SessionPreview,
  SpeakerPreview,
} from "../../common/interfaces";
import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@material-ui/core";
import SessionScheduleRow from "./SessionScheduleRow";

interface Props {
  sessions: Session[];
}

function SessionSchedule(props: Props) {
  return (
    <TableContainer component={Paper}>
      <Typography variant={"h4"} align={"center"}>
        Session schedule
      </Typography>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell>Time</TableCell>
            <TableCell>Speakers</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {/*temporarily we're passing a full session object */}
          {props.sessions.map((session) => (
            <SessionScheduleRow key={session.id} session={session} />
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default SessionSchedule;
