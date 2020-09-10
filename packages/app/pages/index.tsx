import Layout from "../components/Layout";
import React from "react";
import { EventInfo } from "../common/interfaces";
import EventItem from "../components/EventItem";

function IndexPage() {
  const eventInfo: EventInfo = {
    id: "asd123-das-asd",
    name: "The Big Data Event",
    photo:
      "https://images.unsplash.com/photo-1593642634367-d91a135587b5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1650&q=80",
    startDate: new Date(),
    endDate: new Date("12-12-2021"),
    organizer: {
      id: "324k-dsf",
      name: "John Smith",
      photo:
        "https://images.unsplash.com/photo-1599701834133-9ae09fcfa601?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80",
    },
    address: null,
  };

  return (
    <Layout title="Campsite">
      <EventItem eventInfo={eventInfo} />
    </Layout>
  );
}

export default IndexPage;
