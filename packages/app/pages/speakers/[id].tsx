import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { Session, Speaker, SpeakerPreview } from "../../common/interfaces";
import { BASE_API_URL, BASE_SESSION_API_URL } from "../../common/constants";
import SpeakerItem from "../../components/speaker/SpeakerItem";

const BASE_SPEAKER_API_URL = `${BASE_API_URL}/speakers`;

interface Props {
  speaker: Speaker;
}

const StaticPropsDetail = (props: Props) => {
  return (
    <Layout title={`${props.speaker.name}`}>
      <SpeakerItem speaker={props.speaker} />
    </Layout>
  );
};

export default StaticPropsDetail;

export const getStaticProps: GetStaticProps = async ({
  params,
}): Promise<{ props: Props }> => {
  const res = await fetch(`${BASE_SPEAKER_API_URL}/${params?.id}`);
  const speakerPreview: SpeakerPreview = await res.json();
  // Fetch and join sessions on sessionIds (many to many relationship).
  const sessions: Session[] = await Promise.all(
    speakerPreview.sessionIds.map((sessionId) =>
      fetch(`${BASE_SESSION_API_URL}/${sessionId}`).then((res) => res.json()),
    ),
  );
  const speaker: Speaker = { ...speakerPreview, sessions };

  return { props: { speaker } };
};

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await fetch(BASE_SPEAKER_API_URL);
  const speakers: Speaker[] = await res.json();
  const paths = speakers.map((s) => ({ params: { id: s.id } }));
  return {
    paths,
    fallback: false,
  };
};
