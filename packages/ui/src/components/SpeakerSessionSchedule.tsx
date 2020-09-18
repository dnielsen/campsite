import React from "react";
import { Session } from "../common/interfaces";
import util from "../common/util";
import moment from "moment";
import { Link } from "react-router-dom";

interface Props {
  sessions: Session[];
}

function SpeakerSessionSchedule(props: Props) {
  return (
    <div>
      <p>Sessions</p>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Time</th>
            <th>Link</th>
          </tr>
        </thead>
        <tbody>
          {/*temporarily we're passing a full session object */}
          {props.sessions.map((session) => (
            <tr key={session.id}>
              <td>
                <Link to={`/sessions/${session.id}`}>{session.name}</Link>
              </td>
              <td>
                {util.getHourRangeString(session.startDate, session.endDate)} on{" "}
                {moment(session.startDate).format("MM/DD/YYYY")}
              </td>
              <td>
                <Link to={session.url}>{session.url}</Link>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default SpeakerSessionSchedule;
