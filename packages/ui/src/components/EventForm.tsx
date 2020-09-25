import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import { FormProps, FormEventInput } from "../common/interfaces";
import DateTimeField from "./DateTimeField";
import ImageUploadField from "./ImageUploadField";

interface Props {
  formProps: FormProps<FormEventInput>;
}

function EventForm(props: Props) {
  return (
    <Formik {...props.formProps}>
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
            <ImageUploadField name={"photo"} />
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
            <DateTimeField name={"startDate"} />
          </section>
          <section>
            <label htmlFor={"endDate"}>End date</label>
            <DateTimeField name={"endDate"} />
          </section>
          <button type={"submit"} disabled={isSubmitting}>
            Submit
          </button>
        </Form>
      )}
    </Formik>
  );
}

export default EventForm;
