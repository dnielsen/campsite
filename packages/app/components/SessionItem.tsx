import React from "react";
import { Session } from "../common/interfaces";

interface Props {
  session: Session;
}

function SessionItem(props: Props) {
  return <div>hello {props.session.title}</div>;
}

export default SessionItem;
