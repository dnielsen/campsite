import React from "react";
import devData from "../common/devData";
import SessionPreview from "./SessionItemPreview";

function SessionSchedule() {
  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>Session schedule</th>
          </tr>
        </thead>
        <tbody>
          {devData.sessionPreviews.map((preview) => (
            <SessionPreview key={preview.id} sessionPreview={preview} />
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default SessionSchedule;
