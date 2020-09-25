import React from "react";
import { Session } from "../../common/interfaces";
import SessionScheduleRow from "./SessionScheduleRow";

interface Props {
  sessions: Session[];
}

function SessionSchedule(props: Props) {
  return (
    <div>
      <div>Session schedule</div>
      <table>
        <th>
          <tr>
            <td>Name</td>
            <td>Time</td>
            <td>Speakers</td>
          </tr>
        </th>
        <tbody>
          {/*temporarily we're passing a full session object */}
          {props.sessions.map((session) => (
            <SessionScheduleRow key={session.id} session={session} />
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default SessionSchedule;
