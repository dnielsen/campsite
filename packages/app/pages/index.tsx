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
  console.log(`${BASE_EVENT_API_URL}/ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f`);
  const res = await fetch(
    "http://localhost:4444/ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
  );
  console.log(res);
  const eventDetails = await res.json();
  console.log("lololo", eventDetails);

  return {
    props: {
      eventDetails,
    },
  };
};

export default IndexPage;
