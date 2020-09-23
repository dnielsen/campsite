import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  EventDetails,
  FetchEventInput,
  FormEventInput,
  UseForm,
} from "../common/interfaces";
import { BASE_EVENT_API_URL } from "../common/constants";
import moment from "moment-timezone";

export default function useEventForm(): UseForm<FormEventInput> {
  const history = useHistory();

  async function onSubmit(input: FormEventInput) {
    console.log(input);
    // Process the input.
    const fetchInput: FetchEventInput = {
      ...input,
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };

    // Send a request to create the event.
    const createdEvent = (await fetch(BASE_EVENT_API_URL, {
      method: "POST",
      body: JSON.stringify(fetchInput),
    }).then((res) => res.json())) as EventDetails;
    // Redirect to the created session page.
    history.push(`/events/${createdEvent.id}`);
  }

  // For example: `06/27/2020 5:06 PM`
  const now = moment(new Date()).format("mm/DD/yyyy hh:mm a");
  const initialValues: FormEventInput = {
    name: "",
    description: "",
    address: "",
    organizerName: "",
    photo: "",
    startDate: now,
    endDate: now,
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = { onSubmit, validationSchema, initialValues };
  return { formConfig };
}
