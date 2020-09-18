import { CreateSpeakerInput } from "../common/interfaces";
import * as Yup from "yup";
import { BASE_SPEAKER_API_URL } from "../common/constants";
import { FormikProps } from "formik";

export default function useCreateSpeakerForm() {
  // TODO add reset form
  async function handleSubmit(input: CreateSpeakerInput) {
    const createdSpeaker = await fetch(BASE_SPEAKER_API_URL, {
      method: "POST",
      body: JSON.stringify(input),
    }).then((res) => res.json());
    console.log(createdSpeaker);
  }

  const initialValues: CreateSpeakerInput = {
    name: "",
    headline: "",
    bio: "",
    photo: "",
  };

  const validationSchema = Yup.object().shape({});

  return { handleSubmit, validationSchema, initialValues };
}
