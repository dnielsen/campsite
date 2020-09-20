import * as Yup from "yup";
import {
  FetchSessionInput,
  FormSessionInput,
  Option,
  SessionPreview,
  UseForm,
} from "../common/interfaces";
import { BASE_SESSION_API_URL } from "../common/constants";
import { useHistory } from "react-router-dom";

export default function useCreateSessionForm(): UseForm<FormSessionInput> {
  const history = useHistory();

  async function onSubmit(input: FormSessionInput) {
    // Replace speakerOptions property with speakerIds.
    const fetchInput: FetchSessionInput = {
      ...input,
      speakerIds: input.speakerOptions.map((option: Option) => option.value),
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };
    Reflect.deleteProperty(fetchInput, "speakerOptions");

    // Send a request to create the session.
    const createdSession = (await fetch(`${BASE_SESSION_API_URL}`, {
      method: "POST",
      body: JSON.stringify(input),
    }).then((res) => res.json())) as SessionPreview;
    // Redirect to the created session page.
    history.push(`/sessions/${createdSession.id}`);
  }

  const initialValues: FormSessionInput = {
    name: "",
    description: "",
    url: "",
    startDate: "",
    endDate: "",
    speakerOptions: [],
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = { onSubmit, validationSchema, initialValues };
  return { formConfig };
}
