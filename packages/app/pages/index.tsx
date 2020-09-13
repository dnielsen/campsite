import Layout from "../components/Layout";
import React from "react";
import EventItem from "../components/EventItem";
import { EventDetails } from "../common/interfaces";
import { GetStaticProps, InferGetStaticPropsType } from "next";
import { BASE_API_URL } from "../common/constants";

// Temporarily it's the root path, later we might change it to `${BASE_API_URL}/events`
const BASE_EVENT_API_URL = `${BASE_API_URL}`;

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
    // Temporarily until there's multi event support
    `${BASE_EVENT_API_URL}/ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f`,
  );
  const eventDetails = await res.json();

  return {
    props: {
      eventDetails,
    },
  };
};

export default IndexPage;
