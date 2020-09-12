import React from "react";
import { EventDetails } from "../common/interfaces";
import SpeakerList from "./SpeakerList";
import SessionSchedule from "./SessionSchedule";

import "../styles.module.css";

interface Props {
  eventDetails: EventDetails;
}

function EventItem(props: Props) {
  return (
    <div>
      <h2 className={"comments"}>{props.eventDetails.name}</h2>
      <div>
        {/*For now we'll just use the startDate info*/}
        <div>When: {props.eventDetails.startDate.toLocaleString()}</div>
      </div>
      <h4>{props.eventDetails.organizerName}</h4>
      <img src={props.eventDetails.photo} alt="" height={200} />
      <p>{props.eventDetails.description}</p>
      <SessionSchedule sessions={props.eventDetails.sessions} />
      <SpeakerList speakers={props.eventDetails.speakers} />
    </div>
  );
}

export default EventItem;
