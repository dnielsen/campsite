import React from "react";
import { Link, TableCell, TableRow } from "@material-ui/core";
import util from "../../common/util";
import { Session } from "../../common/interfaces";

interface Props {
  session: Session;
}

function SessionScheduleRow(props: Props) {
  const { session } = props;

  return (
    <TableRow key={session.id}>
      <TableCell>
        <Link href={`/sessions/${session.id}`}>{session.name}</Link>
      </TableCell>
      <TableCell>
        {util.getHourRangeString(session.startDate, session.endDate)}{" "}
      </TableCell>
      <TableCell>
        {session.speakers.map((speaker) => (
          <div key={speaker.id}>
            <Link href={`/speakers/${speaker.id}`}>{speaker.name}</Link>
          </div>
        ))}
      </TableCell>
    </TableRow>
  );
}

export default SessionScheduleRow;
