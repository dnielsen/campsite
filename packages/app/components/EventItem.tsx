import React from "react";
import { EventInfo } from "../common/interfaces";
import OrganizerItem from "./OrganizerItem";
import SpeakerList from "./SpeakerList";

interface Props {
  eventInfo: EventInfo;
}

function EventItem(props: Props) {
  return (
    <div>
      <h2>{props.eventInfo.name}</h2>
      <div>
        {/*For now we'll just use the startDate info*/}
        <div>When: {props.eventInfo.startDate.toLocaleString()}</div>
      </div>
      <OrganizerItem organizer={props.eventInfo.organizer} />
      <img src={props.eventInfo.photo} alt="" height={200} />
      <p>event description</p>
      <SpeakerList />
    </div>
  );
}

export default EventItem;
