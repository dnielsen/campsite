import React from "react";
import { Speaker } from "../common/interfaces";

interface Props {
  person: Speaker;
}

function Profile(props: Props) {
  return (
    <div>
      <h2>{props.person.name}</h2>
      <img
        height={200}
        src={props.person.photo}
        alt={`${props.person.name}'s photo`}
      />
      <p>{props.person.bio}</p>
    </div>
  );
}

export default Profile;
