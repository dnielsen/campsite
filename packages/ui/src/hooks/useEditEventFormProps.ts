import * as Yup from "yup";
import {
  EventDetails,
  FetchEventInput,
  FormEventInput,
  FormProps,
} from "../common/interfaces";
import { BASE_EVENT_API_URL } from "../common/constants";
import util from "../common/util";
import useAPI from "./useAPI";
import { useHistory } from "react-router-dom";

interface Props {
  id: string;
}

export default function useEditEventFormProps(
  props: Props,
): FormProps<FormEventInput> {
  const history = useHistory();
  const { data: eventDetails } = useAPI<EventDetails>(`/events/${props.id}`);

  async function onSubmit(input: FormEventInput) {
    // The dates must be of type Date for the backend, however,
    // our DateTimeField needs it in a string form, which is why
    // the form input defines those as strings.
    const fetchInput: FetchEventInput = {
      ...input,
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };

    // Send a request to edit the event with the input.
    await fetch(`${BASE_EVENT_API_URL}/${props.id}`, {
      method: "PUT",
      body: JSON.stringify(fetchInput),
    });

    // Redirect to the edited event page.
    history.push(`/events/${props.id}`);
  }

  // If event details have been fetched then put them in,
  // else leave the properties empty.
  const initialValues: FormEventInput = eventDetails
    ? {
        name: eventDetails.name,
        description: eventDetails.description,
        address: eventDetails.address,
        organizerName: eventDetails.organizerName,
        photo: eventDetails.photo,
        startDate: util.getDateFormValue(eventDetails.startDate),
        endDate: util.getDateFormValue(eventDetails.endDate),
      }
    : {
        name: "",
        description: "",
        address: "",
        organizerName: "",
        photo: "",
        startDate: "",
        endDate: "",
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
