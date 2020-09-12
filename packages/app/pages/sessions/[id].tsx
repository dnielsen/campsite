import { GetStaticProps, GetStaticPaths, InferGetStaticPropsType } from "next";
import Layout from "../../components/Layout";
import React from "react";
import devData from "../../common/devData";
import { Session } from "../../common/interfaces";
import SessionItem from "../../components/SessionItem";
import { useRouter } from "next/router";

interface Props {
  data: Session;
}

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
  const res = await fetch(`http://localhost:4444/sessions/${params?.id}`);
  const session: Session = await res.json();
  return { props: { data: session } };
};

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await fetch("http://localhost:4444/sessions");
  const sessions: Session[] = await res.json();
  const paths = sessions.map((s) => ({ params: { id: s.id } }));
  return {
    paths,
    fallback: false,
  };
};
