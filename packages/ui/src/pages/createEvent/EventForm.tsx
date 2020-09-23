import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import { FormConfig, FormEventInput } from "../../common/interfaces";
import DateTimeField from "../../components/DateTimeField";

interface Props {
  formConfig: FormConfig<FormEventInput>;
}

function EventForm(props: Props) {
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
