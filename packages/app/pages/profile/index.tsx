import { GetStaticProps } from "next";
import Link from "next/link";

import Layout from "../../components/Layout";
import { Person } from "../../common/interfaces";
import { people } from "../../common/devData";

interface Props {
  people: Person[];
}

function WithStaticProps(props: Props) {
  return (
    <Layout title="Users List">
      <h1>Users List</h1>
      <p>
        Example fetching data from inside <code>getStaticProps()</code>.
      </p>
      <p>You are currently on: /users</p>
      <div>{props.people}</div>
      <p>
        <Link href="/">
          <a>Go home</a>
        </Link>
      </p>
    </Layout>
  );
}

export const getStaticProps: GetStaticProps = async () => {
  // Example for including static props in a Next.js function component page.
  // Don't forget to include the respective types for any props passed into
  // the component.
  return { props: { people } };
};

export default WithStaticProps;
