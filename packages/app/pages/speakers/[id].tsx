import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { Speaker } from "../../common/interfaces";
import FullSpeakerItem from "../../components/FullSpeakerItem";

const BASE_SPEAKERS_API_URL = "http://localhost:4444/speakers";

interface Props {
  data: Speaker;
}

const StaticPropsDetail = (props: Props) => {
  return (
    <Layout title={`${props.data.name}`}>
      <FullSpeakerItem person={props.data} />
    </Layout>
  );
};

export default StaticPropsDetail;

export const getStaticProps: GetStaticProps = async ({
  params,
}): Promise<{ props: Props }> => {
  const res = await fetch(`${BASE_SPEAKERS_API_URL}/${params?.id}`);
  const speaker: Speaker = await res.json();
  return { props: { data: speaker } };
};

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await fetch(BASE_SPEAKERS_API_URL);
  const speakers: Speaker[] = await res.json();
  const paths = speakers.map((s) => ({ params: { id: s.id } }));
  return {
    paths,
    fallback: false,
  };
};
