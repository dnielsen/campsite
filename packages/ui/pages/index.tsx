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
    `${BASE_EVENT_API_URL}/e3a27b7d-b37d-4cd2-b8bd-e5bd5551077c`,
  );
  const eventDetails: EventDetails = (await res.json())

  return {
    props: {
      eventDetails,
    },
  };
};

export default IndexPage;
