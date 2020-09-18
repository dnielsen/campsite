import React from "react";
import { EventDetails, SpeakerPreview } from "../common/interfaces";
import useAPI from "../hooks/useAPI";
import util from "../common/util";
import SessionSchedule from "./home/SessionSchedule";
import { Link } from "react-router-dom";
import Speakers from "../components/Speakers";

const EVENT_ID = "e3a27b7d-b37d-4cd2-b8bd-e5bd5551077c";

function Home() {
  const { data: eventDetails, loading, error } = useAPI<EventDetails>(
    `/events/${EVENT_ID}`,
  );

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  const eventSpeakers = util.getUniqueElementsFromMultidimensionalArray(
    // All of the sessions' speakers (with possible duplicates between sessions).
    eventDetails.sessions.map((session) => session.speakers),
  ) as SpeakerPreview[];

  return (
    <div>
      <p>{eventDetails.name}</p>
      <div>
        <p>
          When:{" "}
          {
            // For now we'll just use the start date but we might add
            // support for events that last a few days.
            util.getFullDate(eventDetails.startDate)
          }
        </p>
        <p>Where: {eventDetails.address}</p>
        <p>Organizer: {eventDetails.organizerName}</p>
        <p>
          <Link to={"/"}>Register now!</Link>
        </p>
      </div>
      <img src={eventDetails.photo} alt={eventDetails.name} />
      <p>{eventDetails.description}</p>
      <div>
        <SessionSchedule sessions={eventDetails.sessions} />
      </div>
      <Speakers speakers={eventSpeakers} />
    </div>
  );
}

export default Home;
