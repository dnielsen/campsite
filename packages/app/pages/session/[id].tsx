import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import devData from "../../common/devData";
import { SerializedSession, Session } from "../../common/interfaces";
import SessionItem from "../../components/SessionItem";

interface Props {
  data: SerializedSession;
  error: Error | null;
}

function StaticPropsDetail(props: Props) {
  if (props.error)
    return <div>Something went wrong: {props.error.message}</div>;

  const deserializedSession: Session = {
    ...props.data,
    startDate: new Date(props.data.startDate),
    endDate: new Date(props.data.endDate),
  };

  return (
    <Layout title={props.data.title}>
      <SessionItem session={deserializedSession} />
    </Layout>
  );
}

export default StaticPropsDetail;

export const getStaticPaths: GetStaticPaths = async () => {
  // Get the paths we want to pre-render based on users
  const paths = devData.sessions.map((session) => ({
    params: { id: session.id },
  }));

  // We'll pre-render only these paths at build time.
  // { fallback: false } means other routes should 404.
  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const id = params?.id;
  const foundSession = devData.sessions.find((session) => session.id === id);
  if (!foundSession) throw new Error("Session not found");
  return { props: { data: foundSession } };
};
