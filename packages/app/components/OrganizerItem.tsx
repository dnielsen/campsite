import React from "react";
import { Organizer } from "../common/interfaces";

interface Props {
  organizer: Organizer;
}

function OrganizerItem(props: Props) {
  return (
    <div>
      <div>{props.organizer.name}</div>
      <img
        src={props.organizer.photo}
        alt={`${props.organizer.photo}'s photo`}
        height={200}
      />
    </div>
  );
}

export default OrganizerItem;
