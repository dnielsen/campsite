import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import SelectField from "./SelectField";
import {
  FormConfig,
  FormSessionInput,
  Option,
  SpeakerPreview,
} from "../common/interfaces";
import useAPI from "../hooks/useAPI";

interface Props {
  formConfig: FormConfig<FormSessionInput>;
}

function SessionForm(props: Props) {
  const { data: speakers, loading, error } = useAPI<SpeakerPreview[]>(
    "/speakers",
  );

  if (loading) return <div>loading...</div>;
  if (error) return <div>error: {error.message}</div>;

  const options: Option[] = speakers.map((speaker) => ({
    label: speaker.name,
    value: speaker.id,
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
            <label htmlFor={"url"}>Url</label>
            <Field type={"text"} name={"url"} />
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
            <label htmlFor={"speakerOptions"}>Speakers</label>
            <SelectField
              options={options}
              name={"speakerOptions"}
              defaultValue={props.formConfig.initialValues.speakerOptions}
            />
          </section>
          <button type={"submit"} disabled={isSubmitting}>
            Submit
          </button>
        </Form>
      )}
    </Formik>
  );
}

export default SessionForm;
