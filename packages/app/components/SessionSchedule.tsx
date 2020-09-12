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
                {util.getHourRangeString(session.startDate, session.endDate)}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default SessionSchedule;
