import Layout from "../components/Layout";
import React from "react";
import EventItem from "../components/EventItem";
import devData from "../common/devData";

function IndexPage() {
  return (
    <Layout title="Campsite">
      <EventItem eventInfo={devData.eventInfo} />
    </Layout>
  );
}

export default IndexPage;
