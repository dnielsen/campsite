import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  EventDetails,
  FetchEventInput,
  FormEventInput,
  Option,
  UseForm,
} from "../common/interfaces";
import { BASE_EVENT_API_URL } from "../common/constants";

export default function useCreateEventForm(): UseForm<FormEventInput> {
  const history = useHistory();

  async function onSubmit(input: FormEventInput) {
    // Process the input.
    const fetchInput: FetchEventInput = {
      ...input,
      sessionIds: input.sessionOptions.map((option: Option) => option.value),
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };
    Reflect.deleteProperty(fetchInput, "sessionOptions");

    // Send a request to create the event.
    const createdEvent = (await fetch(BASE_EVENT_API_URL, {
      method: "POST",
      body: JSON.stringify(fetchInput),
    }).then((res) => res.json())) as EventDetails;
    // Redirect to the created session page.
    history.push(`/events/${createdEvent.id}`);
  }

  const initialValues: FormEventInput = {
    name: "",
    description: "",
    address: "",
    organizerName: "",
    photo: "",
    startDate: "",
    endDate: "",
    sessionOptions: [],
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = { onSubmit, validationSchema, initialValues };
  return { formConfig };
}
