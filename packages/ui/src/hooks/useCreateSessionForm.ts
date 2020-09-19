import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import { CreateSessionInput, SessionPreview } from "../common/interfaces";
import { BASE_SESSION_API_URL } from "../common/constants";

export default function useCreateSessionForm() {
  const history = useHistory();

  async function handleSubmit(input: CreateSessionInput) {
    console.log(input);
    const properInput = {
      ...input,
      speakerIds: input.speakerOptions.map((option) => option.value),
    };
    Reflect.deleteProperty(properInput, "speakerOptions");
    console.log(properInput);
    // Send a request to create the session.
    const createdSession = (await fetch(BASE_SESSION_API_URL, {
      method: "POST",
      body: JSON.stringify(properInput),
    }).then((res) => res.json())) as SessionPreview;
    // Redirect to the created session page.
    history.push(`/sessions/${createdSession.id}`);
  }

  const initialValues: CreateSessionInput = {
    name: "new session",
    description: "the description",
    url: "https://apple.com",
    startDate: new Date(),
    endDate: new Date(),
    speakerOptions: [],
  };

  const validationSchema = Yup.object().shape({});

  return { handleSubmit, validationSchema, initialValues };
}
