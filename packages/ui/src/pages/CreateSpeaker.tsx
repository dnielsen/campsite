import React from "react";
import { CreateSpeakerInput } from "../common/interfaces";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import useCreateSpeakerForm from "../hooks/useCreateSpeakerForm";

const INPUT: CreateSpeakerInput = {
  bio: "The bio",
  headline: "The headline",
  name: "The name",
  photo:
    "https://www.biography.com/.image/t_share/MTY2NzA3ODE3OTgwMzcyMjYw/jeff-bezos-andrew-harrer_bloomberg-via-getty-images.jpg",
};

function CreateSpeaker() {
  const {
    initialValues,
    handleSubmit,
    validationSchema,
  } = useCreateSpeakerForm();

  return (
    <div>
      <h3>Create a speaker</h3>
      <Formik
        initialValues={initialValues}
        onSubmit={handleSubmit}
        validationSchema={validationSchema}
      >
        {({ isSubmitting }: FormikState<FormikValues>) => (
          <Form noValidate>
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
              <Field type={"text"} name={"photo"} />
            </section>
            <button type={"submit"} disabled={isSubmitting}>
              Create
            </button>
          </Form>
        )}
      </Formik>
    </div>
  );
}

export default CreateSpeaker;
