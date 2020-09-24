import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import { FormProps, FormSpeakerInput } from "../common/interfaces";
import ImageUploadField from "./ImageUploadField";

interface Props {
  formProps: FormProps<FormSpeakerInput>;
}

function SpeakerForm(props: Props) {
  return (
    <Formik {...props.formProps}>
      {({ isSubmitting }: FormikState<FormikValues>) => (
        <Form>
          <section>
            <label htmlFor="name">Name</label>
            <Field type={"text"} name={"name"} />
          </section>
          <section>
            <label htmlFor="bio">Bio</label>
            <Field type={"text"} name={"bio"} />
          </section>
          <section>
            <label htmlFor="headline">Headline</label>
            <Field type={"text"} name={"headline"} />
          </section>
          <section>
            {/*For now it's just a url, later we might add a photo upload*/}
            <label htmlFor="photo">Photo</label>
            <ImageUploadField name={"photo"} />
          </section>
          <button type={"submit"} disabled={isSubmitting}>
            Submit
          </button>
        </Form>
      )}
    </Formik>
  );
}

export default SpeakerForm;
