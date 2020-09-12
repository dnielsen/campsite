import { GetStaticProps, GetStaticPaths, InferGetStaticPropsType } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { Session } from "../../common/interfaces";
import SessionItem from "../../components/SessionItem";

interface Props {
  data: Session;
}

const BASE_SESSIONS_API_URL = "http://localhost:4444/sessions";

function StaticPropsDetail(props: Props) {
  return (
    <Layout title={props.data.name}>
      <SessionItem session={props.data} />
    </Layout>
  );
}

export default StaticPropsDetail;

export const getStaticProps: GetStaticProps = async ({
  params,
}): Promise<{ props: Props }> => {
  const res = await fetch(`${BASE_SESSIONS_API_URL}/${params?.id}`);
  const session: Session = await res.json();
  console.log(session);
  return { props: { data: session } };
};

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await fetch(BASE_SESSIONS_API_URL);
  const sessions: Session[] = await res.json();
  const paths = sessions.map((s) => ({ params: { id: s.id } }));
  return {
    paths,
    fallback: false,
  };
};
