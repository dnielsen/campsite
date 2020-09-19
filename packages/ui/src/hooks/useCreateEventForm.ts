import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  CreateEventFetchInput,
  CreateEventFormInput,
  EventDetails,
  Option,
} from "../common/interfaces";
import { BASE_EVENT_API_URL } from "../common/constants";
import util from "../common/util";

export default function useCreateEventForm() {
  const history = useHistory();

  async function handleSubmit(input: CreateEventFormInput) {
    // Process the input.
    const fetchInput: CreateEventFetchInput = {
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
    history.push(`/${createdEvent.id}`);
  }

  const initialValues: CreateEventFormInput = {
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

  return { handleSubmit, validationSchema, initialValues };
}
