import React from "react";
import { Session } from "../common/interfaces";
import {
  Link,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@material-ui/core";
import util from "../common/util";
import moment from "moment";

interface Props {
  sessions: Session[];
}

function SpeakerSessionSchedule(props: Props) {
  return (
    <TableContainer component={Paper}>
      <Typography variant={"h4"} align={"center"}>
        Sessions
      </Typography>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell>Time</TableCell>
            <TableCell>Link</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {/*temporarily we're passing a full session object */}
          {props.sessions.map((session) => (
            <TableRow key={session.id}>
              <TableCell>
                <Link href={`/sessions/${session.id}`}>{session.name}</Link>
              </TableCell>
              <TableCell>
                {util.getHourRangeString(session.startDate, session.endDate)} on{" "}
                {moment(session.startDate).format("MM/DD/YYYY")}
              </TableCell>
              <TableCell>
                <Link href={session.url}>{session.url}</Link>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default SpeakerSessionSchedule;
