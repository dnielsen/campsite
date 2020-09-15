import React from "react";
import { Session, Speaker } from "../common/interfaces";
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
  speakers: Speaker[];
  // Otherwise it's a speaker schedule and we don't display the speaker name
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
          {props.sessions.map((session) => {
            session.speakers = props.speakers.filter((speaker) =>
              session.speakerIds.includes(speaker.id),
            );
            return (
              <TableRow key={session.id}>
                <TableCell>
                  <Link href={`/sessions/${session.id}`}>{session.name}</Link>
                </TableCell>
                <TableCell>
                  {util.getHourRangeString(session.startDate, session.endDate)}{" "}
                  {/*{moment(session.startDate)*/}
                  {/*  .tz("America/Los_Angeles")*/}
                  {/*  .format("ha")}*/}
                </TableCell>

                <TableCell>
                  {session.speakers.map((speaker) => (
                    <div key={speaker.id}>
                      <Link href={`/speakers/${speaker.id}`}>
                        {speaker.name}
                      </Link>
                    </div>
                  ))}
                </TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default SessionSchedule;
