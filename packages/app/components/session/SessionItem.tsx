import React from "react";
import { Session } from "../../common/interfaces";
import SpeakerList from "../speaker/SpeakerList";
import moment from "moment";
import util from "../../common/util";

interface Props {
  session: Session;
}

function SessionItem(props: Props) {
  return (
    <div>
      <div>
        <div>
          <SpeakerList speakers={props.session.speakers} />
        </div>
        <div>
          <div>
            <p>{props.session.name}</p>
            <p>
              {util.getHourRangeString(
                props.session.startDate,
                props.session.endDate,
              )}{" "}
              on {moment(props.session.startDate).format("MM/DD/YYYY")}
            </p>
            <a href={props.session.url}>{props.session.url}</a>
            <p>{props.session.description}</p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default SessionItem;
