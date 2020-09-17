import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { Session } from "../../common/interfaces";
import SessionItem from "../../components/session/SessionItem";
import { BASE_API_URL } from "../../common/constants";
import { zipkinFetch as fetch } from "../../common/fetch";

interface Props {
  data: Session;
}

const BASE_SESSION_API_URL = `${BASE_API_URL}/sessions`;

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
  const res = await fetch(`${BASE_SESSION_API_URL}/${params?.id}`);
  const session: Session = await res.json();

  return { props: { data: session } };
};

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await fetch(BASE_SESSION_API_URL);
  const sessions: Session[] = await res.json();
  const paths = sessions.map((s) => ({ params: { id: s.id } }));
  return {
    paths,
    fallback: false,
  };
};
