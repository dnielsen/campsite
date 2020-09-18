import React from "react";
import { Session } from "../../common/interfaces";
import util from "../../common/util";
import { Link } from "react-router-dom";

interface Props {
  sessions: Session[];
}

function SessionSchedule(props: Props) {
  return (
    <div>
      <h3>Session schedule</h3>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Time</th>
            <th>Speakers</th>
          </tr>
        </thead>
        <tbody>
          {props.sessions.map((session) => (
            <tr key={session.id}>
              <td>
                <Link to={`/sessions/${session.id}`}>{session.name}</Link>
              </td>
              <td>
                {util.getHourRangeString(session.startDate, session.endDate)}{" "}
              </td>
              <td>
                {session.speakers.map((speaker) => (
                  <div key={speaker.id}>
                    <Link to={`/speakers/${speaker.id}`}>{speaker.name}</Link>
                  </div>
                ))}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default SessionSchedule;
