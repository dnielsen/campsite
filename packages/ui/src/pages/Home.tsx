import React from "react";
import { EventDetails, SpeakerPreview } from "../common/interfaces";
import useAPI from "../hooks/useAPI";
import util from "../common/util";
import SessionSchedule from "./home/SessionSchedule";
import { Link } from "react-router-dom";
import SpeakerList from "../components/SpeakerList";

const EVENT_ID = "80e211fd-d33a-4149-8d8b-47ff5e4c8f0f";

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
          {/*// For now we'll just use the start date but we might add*/}
          {/*// support for events that last a few days.*/}
          When: {util.getFullDateString(eventDetails.startDate)}
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
      <SpeakerList speakers={eventSpeakers} />
    </div>
  );
}

export default Home;
