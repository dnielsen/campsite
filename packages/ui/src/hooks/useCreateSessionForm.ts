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

interface Props {
  defaultEventIdValue: string;
}

export default function useCreateSessionForm(
  props: Props,
): UseForm<FormSessionInput> {
  const history = useHistory();

  async function onSubmit(input: FormSessionInput) {
    // Replace speakerOptions property with speakerIds.
    const fetchInput: FetchSessionInput = {
      ...input,
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };

    console.log(fetchInput);
    // Send a request to create the session.
    const createdSession = (await fetch(`${BASE_SESSION_API_URL}`, {
      method: "POST",
      body: JSON.stringify(fetchInput),
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
    eventId: props.defaultEventIdValue,
    speakerIds: [],
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = { onSubmit, validationSchema, initialValues };
  return { formConfig };
}
