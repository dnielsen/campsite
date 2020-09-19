import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import { FormConfig, SpeakerInput, SpeakerPreview } from "../common/interfaces";
import { BASE_SPEAKER_API_URL } from "../common/constants";

export default function useCreateSpeakerForm(): FormConfig<SpeakerInput> {
  const history = useHistory();

  async function onSubmit(input: SpeakerInput) {
    // Send a request to create the speaker.
    const createdSpeaker = (await fetch(BASE_SPEAKER_API_URL, {
      method: "POST",
      body: JSON.stringify(input),
    }).then((res) => res.json())) as SpeakerPreview;
    // Redirect to the created speaker page.
    history.push(`/speakers/${createdSpeaker.id}`);
  }

  const initialValues: SpeakerInput = {
    name: "",
    photo: "",
    headline: "",
    bio: "",
  };

  const validationSchema = Yup.object().shape({});

  return {
    onSubmit,
    initialValues,
    validationSchema,
  };
}
