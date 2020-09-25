import * as Yup from "yup";
import {
  FormProps,
  FormSpeakerInput,
  SpeakerPreview,
} from "../common/interfaces";
import { BASE_SPEAKER_API_URL } from "../common/constants";
import useAPI from "./useAPI";
import { useHistory } from "react-router-dom";

interface Props {
  id: string;
}

export default function useEditSpeakerFormProps(
  props: Props,
): FormProps<FormSpeakerInput> {
  const history = useHistory();
  const { data: speaker } = useAPI<SpeakerPreview>(`/speakers/${props.id}`);

  async function onSubmit(input: FormSpeakerInput) {
    // Send a request to edit the speaker with the input.
    await fetch(`${BASE_SPEAKER_API_URL}/${props.id}`, {
      method: "PUT",
      body: JSON.stringify(input),
    });

    // Redirect to the edited speaker page.
    history.push(`/speakers/${props.id}`);
  }

  // If event details have been fetched then put them in,
  // else leave the properties empty.
  const initialValues: FormSpeakerInput = speaker
    ? {
        name: speaker.name,
        photo: speaker.photo,
        bio: speaker.bio,
        headline: speaker.headline,
      }
    : {
        name: "",
        photo: "",
        bio: "",
        headline: "",
      };

  const validationSchema = Yup.object().shape({
    // name: Yup.string().min(2).max(50)
  });

  return {
    onSubmit,
    validationSchema,
    initialValues,
    enableReinitialize: true,
  };
}
