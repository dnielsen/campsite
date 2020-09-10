import Layout from "../components/Layout";
import React from "react";
import EventItem from "../components/EventItem";
import { eventInfo } from "../common/devData";

function IndexPage() {
  return (
    <Layout title="Campsite">
      <EventItem eventInfo={eventInfo} />
    </Layout>
  );
}

export default IndexPage;
