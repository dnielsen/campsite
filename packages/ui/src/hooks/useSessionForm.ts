import * as Yup from "yup";
import {
  FetchSessionInput,
  FormSessionInput,
  SessionPreview,
  UseForm,
} from "../common/interfaces";
import { BASE_SESSION_API_URL } from "../common/constants";
import { useHistory } from "react-router-dom";
import moment from "moment-timezone";

interface Props {
  defaultEventIdValue: string;
}

export default function useSessionForm(
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

    // Send a request to create the session.
    const createdSession = (await fetch(`${BASE_SESSION_API_URL}`, {
      method: "POST",
      body: JSON.stringify(fetchInput),
    }).then((res) => res.json())) as SessionPreview;
    // Redirect to the created session page.
    history.push(`/sessions/${createdSession.id}`);
  }

  // For example: `06/27/2020 5:06 PM`
  const now = moment(new Date()).format("mm/DD/yyyy hh:mm a");
  const initialValues: FormSessionInput = {
    name: "",
    description: "",
    url: "",
    startDate: now,
    endDate: now,
    eventId: props.defaultEventIdValue,
    speakerIds: [],
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = { onSubmit, validationSchema, initialValues };
  return { formConfig };
}
