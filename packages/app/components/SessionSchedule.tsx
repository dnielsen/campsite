import React from "react";
import devData from "../common/devData";
import SessionPreview from "./SessionItemPreview";
import { Session } from "../common/interfaces";

interface Props {
  sessions: Session[];
}

function SessionSchedule(props: Props) {
  console.log(props.sessions);
  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>Session schedule</th>
          </tr>
        </thead>
        <tbody>
          {/*temporarily we're passing a full session object */}
          {props.sessions.map((session) => (
            <SessionPreview key={session.id} sessionPreview={session} />
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default SessionSchedule;
