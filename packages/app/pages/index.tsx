import Layout from "../components/Layout";
import React from "react";
import EventItem from "../components/event/EventItem";
import { EventDetails, EventResponse } from "../common/interfaces";
import { GetStaticProps, InferGetStaticPropsType } from "next";
import {
  BASE_EVENT_API_URL,
  BASE_SESSION_API_URL,
  BASE_SPEAKER_API_URL,
} from "../common/constants";
import { zipkinFetch as fetch } from "../common/fetch";

function IndexPage(props: InferGetStaticPropsType<typeof getStaticProps>) {
  return (
    <Layout title="Campsite">
      <EventItem eventDetails={props.eventDetails} />
    </Layout>
  );
}

export const getStaticProps: GetStaticProps = async ({
  params,
}): Promise<{ props: { eventDetails: EventDetails } }> => {
  const res = await fetch(
    `${BASE_EVENT_API_URL}/ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f`,
  );
  const event = (await res.json()) as EventResponse;

  const sessionsPromise = Promise.all(
    event.sessionIds.map((sessionId) =>
      fetch(`${BASE_SESSION_API_URL}/${sessionId}`).then((res: any) =>
        res.json(),
      ),
    ),
  );
  const speakersPromise = Promise.all(
    event.speakerIds.map((speakerId) =>
      fetch(`${BASE_SPEAKER_API_URL}/${speakerId}`).then((res: any) =>
        res.json(),
      ),
    ),
  );

  const [sessionPreviews, speakerPreviews] = await Promise.all([
    sessionsPromise,
    speakersPromise,
  ]);

  // Add speakers to the sessions
  const sessions = sessionPreviews.map((session) => ({
    ...session,
    speakers: speakerPreviews.filter((speaker) =>
      speaker.sessionIds.includes(session.id),
    ),
  }));

  const eventDetails = {
    ...event,
    sessions: sessions,
    speakers: speakerPreviews,
  };

  return {
    props: {
      eventDetails,
    },
  };
};

export default IndexPage;
