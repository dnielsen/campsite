import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { people } from "../../common/devData";
import { Person } from "../../common/interfaces";
import Profile from "../../components/Profile";

type Props = {
  data: Person;
  error: Error | null;
};

const StaticPropsDetail = (props: Props) => {
  if (props.error)
    return <div>Something went wrong: {props.error.message}</div>;

  return (
    <Layout title={`${props.data.name}`}>
      <Profile person={props.data} />
    </Layout>
  );
};

export default StaticPropsDetail;

export const getStaticPaths: GetStaticPaths = async () => {
  // Get the paths we want to pre-render based on users
  const paths = people.map((person) => ({
    params: { id: person.id },
  }));

  // We'll pre-render only these paths at build time.
  // { fallback: false } means other routes should 404.
  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const id = params?.id;
  const foundPerson = people.find((person) => person.id === id);
  if (!foundPerson) throw new Error("Person not found");
  return { props: { data: foundPerson } };
};
