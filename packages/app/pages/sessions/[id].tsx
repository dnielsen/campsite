import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { Session, SessionPreview } from "../../common/interfaces";
import SessionItem from "../../components/session/SessionItem";
import { BASE_API_URL, BASE_SPEAKER_API_URL } from "../../common/constants";
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
  const sessionPreview: SessionPreview = await res.json();
  // Fetch and join speakers on speakerIds (many to many relationship).
  const speakers = await Promise.all(
    sessionPreview.speakerIds.map((speakerId) =>
      fetch(`${BASE_SPEAKER_API_URL}/${speakerId}`).then((res: any) =>
        res.json(),
      ),
    ),
  );
  const session = { ...sessionPreview, speakers };

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
