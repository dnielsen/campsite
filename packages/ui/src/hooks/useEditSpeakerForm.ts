import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  FormConfig,
  Speaker,
  SpeakerInput,
  SpeakerPreview,
} from "../common/interfaces";
import { BASE_SPEAKER_API_URL } from "../common/constants";
import useAPI from "./useAPI";
import { useEffect, useState } from "react";

interface UseEditSpeakerForm {
  formConfig: FormConfig<SpeakerInput>;
  loading: boolean;
  error: Error | null;
}

export default function useEditSpeakerForm(id: string): UseEditSpeakerForm {
  const history = useHistory();
  const [initialValues, setInitialValues] = useState<SpeakerInput>({
    name: "",
    photo: "",
    headline: "",
    bio: "",
  });
  const { data: uneditedSpeaker, loading, error } = useAPI<Speaker>(
    `/speakers/${id}`,
  );

  useEffect(() => {
    // Load the speaker values so that we don't have to start writing
    // speaker fields from zero.
    if (uneditedSpeaker) {
      setInitialValues({
        bio: uneditedSpeaker.bio,
        headline: uneditedSpeaker.headline,
        name: uneditedSpeaker.name,
        photo: uneditedSpeaker.photo,
      });
    }
  }, [uneditedSpeaker]);

  async function onSubmit(input: SpeakerInput) {
    // Send a request to create the speaker.
    const editedSpeaker = (await fetch(`${BASE_SPEAKER_API_URL}/${id}`, {
      method: "PUT",
      body: JSON.stringify(input),
    }).then((res) => res.json())) as SpeakerPreview;
    // Redirect to the created speaker page.
    history.push(`/speakers/${editedSpeaker.id}`);
  }

  const validationSchema = Yup.object().shape({});

  const formConfig = {
    onSubmit,
    validationSchema,
    initialValues,
  };

  return { formConfig, loading, error };
}
