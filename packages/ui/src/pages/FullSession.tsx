import React from "react";
import useAPI from "../hooks/useAPI";
import { Session } from "../common/interfaces";
import { useParams } from "react-router-dom";
import util from "../common/util";
import moment from "moment";
import Speakers from "../components/Speakers";

function FullSession() {
  const { id } = useParams<{ id: string }>();

  const { data: session, loading, error } = useAPI<Session>(`/sessions/${id}`);

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  return (
    <div>
      <Speakers speakers={session.speakers} />
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
    </div>
  );
}

export default FullSession;
