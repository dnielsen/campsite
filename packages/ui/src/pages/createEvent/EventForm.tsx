import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import SelectField from "./SelectField";
import {
  FormConfig,
  FormEventInput,
  Option,
  SessionPreview,
} from "../common/interfaces";
import useAPI from "../hooks/useAPI";

interface Props {
  formConfig: FormConfig<FormEventInput>;
}

function EventForm(props: Props) {
  const { data: sessions, loading, error } = useAPI<SessionPreview[]>(
    "/sessions",
  );

  if (loading) return <div>loading...</div>;
  if (error) return <div>error: {error.message}</div>;

  const options: Option[] = sessions.map((session) => ({
    label: session.name,
    value: session.id,
  }));

  return (
    <Formik {...props.formConfig}>
      {({ isSubmitting }: FormikState<FormikValues>) => (
        <Form noValidate>
          <section>
            <label htmlFor={"name"}>Name</label>
            <Field type={"text"} name={"name"} />
          </section>
          <section>
            <label htmlFor={"description"}>Description</label>
            <Field type={"text"} name={"description"} />
          </section>
          <section>
            <label htmlFor={"photo"}>Photo</label>
            <Field type={"text"} name={"photo"} />
          </section>
          <section>
            <label htmlFor={"organizerName"}>Organizer name</label>
            <Field type={"text"} name={"organizerName"} />
          </section>
          <section>
            <label htmlFor={"address"}>Address</label>
            <Field type={"text"} name={"address"} />
          </section>
          <section>
            <label htmlFor={"startDate"}>Start date</label>
            <Field type={"date"} name={"startDate"} />
          </section>
          <section>
            <label htmlFor={"endDate"}>End date</label>
            <Field type={"date"} name={"endDate"} />
          </section>
          <section>
            <label htmlFor={"sessionOptions"}>Sessions</label>
            <SelectField options={options} name={"sessionOptions"} />
          </section>
          <button type={"submit"} disabled={isSubmitting}>
            Create
          </button>
        </Form>
      )}
    </Formik>
  );
}

export default EventForm;
