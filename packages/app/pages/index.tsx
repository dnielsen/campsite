import Layout from "../components/Layout";
import React from "react";
import SpeakerList from "../components/SpeakerList";

function IndexPage() {
  return (
    <Layout title="Campsite">
      <h2>Event name</h2>
      <div>
        <div>3/20/2021</div>
        <div>7:00pm - 9:00pm</div>
      </div>
      <img
        src="https://images.unsplash.com/photo-1593642634367-d91a135587b5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1650&q=80"
        alt=""
        height={200}
      />
      <p>event description</p>
      <SpeakerList />
    </Layout>
  );
}

export default IndexPage;
