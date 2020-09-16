import React from "react";
import util from "../../common/util";
import { Session } from "../../common/interfaces";

interface Props {
  session: Session;
}

function SessionScheduleRow(props: Props) {
  const { session } = props;

  return (
    <tr>
      <td>
        <a href={`/sessions/${session.id}`}>{session.name}</a>
      </td>
      <td>{util.getHourRangeString(session.startDate, session.endDate)} </td>
      <td>
        {session.speakers.map((speaker) => (
          <div key={speaker.id}>
            <a href={`/speakers/${speaker.id}`}>{speaker.name}</a>
          </div>
        ))}
      </td>
    </tr>
  );
}

export default SessionScheduleRow;
