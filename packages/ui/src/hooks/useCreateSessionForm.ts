import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  CreateSessionFetchInput,
  CreateSessionFormInput,
  Option,
  SessionPreview,
} from "../common/interfaces";
import { BASE_SESSION_API_URL } from "../common/constants";
import util from "../common/util";

export default function useCreateSessionForm() {
  const history = useHistory();

  async function handleSubmit(input: CreateSessionFormInput) {
    // Replace speakerOptions property with speakerIds.
    const fetchInput: CreateSessionFetchInput = {
      ...input,
      speakerIds: input.speakerOptions.map((option: Option) => option.value),
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };
    const createdSession = (await fetch(BASE_SESSION_API_URL, {
      method: "POST",
      body: JSON.stringify(fetchInput),
    }).then((res) => res.json())) as SessionPreview;
    // Redirect to the created session page.
    history.push(`/sessions/${createdSession.id}`);
  }

  const initialValues: CreateSessionFormInput = {
    name: "",
    description: "",
    url: "",
    startDate: "",
    endDate: "",
    speakerOptions: [],
  };

  const validationSchema = Yup.object().shape({});

  return { handleSubmit, validationSchema, initialValues };
}
