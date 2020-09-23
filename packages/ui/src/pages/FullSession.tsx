import React from "react";
import useAPI from "../hooks/useAPI";
import { Session } from "../common/interfaces";
import { Link, useHistory, useParams } from "react-router-dom";
import util from "../common/util";
import moment from "moment";
import SpeakerList from "../components/SpeakerList";
import {
  BASE_SESSION_API_URL,
  BASE_SPEAKER_API_URL,
} from "../common/constants";
import SessionItem from "../bilal/session/SessionItem";

function FullSession() {
  const { id } = useParams<{ id: string }>();
  const history = useHistory();

  const { data: session, loading, error } = useAPI<Session>(`/sessions/${id}`);

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  return <SessionItem session={session} />;

  async function handleClick() {
    await fetch(`${BASE_SESSION_API_URL}/${id}`, { method: "DELETE" });
    // Redirect to the home page after deleting the speaker.
    history.push("/");
  }

  return (
    <div>
      <SpeakerList speakers={session.speakers} />
      <div>
        <div>
          <h3>{session.name}</h3>
          <p>
            {util.getHourRangeString(session.startDate, session.endDate)} on{" "}
            {moment(session.startDate).format("MM/DD/YYYY")}
          </p>
          <a href={session.url}>{session.url}</a>
          <p>{session.description}</p>
        </div>
      </div>
      <button type={"button"} onClick={handleClick}>
        Delete
      </button>
    </div>
  );
}

export default FullSession;
