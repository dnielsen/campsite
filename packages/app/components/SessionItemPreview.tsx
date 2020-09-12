import React from "react";
import { SessionPreview } from "../common/interfaces";
import util from "../common/util";

interface Props {
  sessionPreview: SessionPreview;
}

function SessionItemPreview(props: Props) {
  return (
    <tr>
      <td>
        <a href={`/sessions/${props.sessionPreview.id}`}>
          {props.sessionPreview.name}
        </a>
      </td>
      <td>
        {util.getHourRangeString(
          props.sessionPreview.startDate,
          props.sessionPreview.endDate,
        )}
      </td>
    </tr>
  );
}

export default SessionItemPreview;
