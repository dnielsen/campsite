import React from "react";
import { Session } from "../common/interfaces";
import util from "../common/util";
import moment from "moment";

interface Props {
  sessions: Session[];
}

function SpeakerSessionSchedule(props: Props) {
  return (
    <div>
      <p>Sessions</p>
      <table>
        <th>
          <tr>
            <td>Name</td>
            <td>Time</td>
            <td>Link</td>
          </tr>
        </th>
        <tbody>
          {/*temporarily we're passing a full session object */}
          {props.sessions.map((session) => (
            <tr key={session.id}>
              <td>
                <a href={`/sessions/${session.id}`}>{session.name}</a>
              </td>
              <td>
                {util.getHourRangeString(session.startDate, session.endDate)} on{" "}
                {moment(session.startDate).format("MM/DD/YYYY")}
              </td>
              <td>
                <a href={session.url}>{session.url}</a>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default SpeakerSessionSchedule;
