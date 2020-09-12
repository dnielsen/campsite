import Layout from "../components/Layout";
import React from "react";
import EventItem from "../components/EventItem";
import { EventDetails } from "../common/interfaces";
import { GetStaticProps, InferGetStaticPropsType } from "next";

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
    "http://localhost:4444/ad29d4f9-b0dd-4ea3-9e96-5ff193b50d6f",
  );
  const eventDetails = await res.json();

  return {
    props: {
      eventDetails,
    },
  };
};

// export const getStaticPaths: GetStaticPaths = async () => {
//   // Call an external API endpoint to get posts
//   const res = await fetch("http://localhost:4444/doesnt-matter-now/sessions");
//   const posts = await res.json();
//
//   // Get the paths we want to pre-render based on posts
//   const paths = posts.map((post) => ({
//     params: { id: post.id },
//   }));
//
//   // We'll pre-render only these paths at build time.
//   // { fallback: false } means other routes should 404.
//   return { paths, fallback: false };
// };

export default IndexPage;
