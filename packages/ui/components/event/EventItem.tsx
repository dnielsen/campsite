import React from "react";
import { EventDetails } from "../../common/interfaces";
import SpeakerList from "../speaker/SpeakerList";
import SessionSchedule from "../session/SessionSchedule";
import util from "../../common/util";

interface Props {
  eventDetails: EventDetails;
}

function EventItem(props: Props) {
  return (
    <div>
      <p>{props.eventDetails.name}</p>
      {/*For now we'll just use the startDate info*/}
      <div>
        <p>When: {util.getFullDate(props.eventDetails.startDate)}</p>
        <p>Where: {props.eventDetails.address}</p>
        <p>Organizer: {props.eventDetails.organizerName}</p>
        <p>
          <a href={"/"}>Register now!</a>
        </p>
      </div>
      <img src={props.eventDetails.photo} alt={props.eventDetails.name} />
      <p>{props.eventDetails.description}</p>
      <div>
        <SessionSchedule sessions={props.eventDetails.sessions} />
      </div>
      <p>Our speakers</p>
      <div>
        <SpeakerList speakers={props.eventDetails.speakers} />
      </div>
    </div>
  );
}

export default EventItem;
