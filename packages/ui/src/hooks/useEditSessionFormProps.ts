import * as Yup from "yup";
import {
  EventDetails,
  FetchEventInput,
  FetchSessionInput,
  FormEventInput,
  FormProps,
  FormSessionInput,
  Session,
} from "../common/interfaces";
import { BASE_EVENT_API_URL, BASE_SESSION_API_URL } from "../common/constants";
import util from "../common/util";
import useAPI from "./useAPI";
import { useHistory } from "react-router-dom";

interface Props {
  id: string;
}

export default function useEditSessionFormProps(
  props: Props,
): FormProps<FormSessionInput> {
  const history = useHistory();
  const { data: session } = useAPI<Session>(`/sessions/${props.id}`);

  async function onSubmit(input: FormSessionInput) {
    // The dates must be of type Date for the backend, however,
    // our DateTimeField needs it in a string form, which is why
    // the form input defines those as strings.
    const fetchInput: FetchSessionInput = {
      ...input,
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };

    // Send a request to edit the event with the input.
    await fetch(`${BASE_SESSION_API_URL}/${props.id}`, {
      method: "PUT",
      body: JSON.stringify(fetchInput),
    });

    // Redirect to the edited event page.
    history.push(`/sessions/${props.id}`);
  }

  // If event details have been fetched then put them in,
  // else leave the properties empty.
  const initialValues: FormSessionInput = session
    ? {
        name: session.name,
        description: session.description,
        url: session.url,
        startDate: util.getDateFormValue(session.startDate),
        endDate: util.getDateFormValue(session.endDate),
        speakerIds: session.speakers.map((speaker) => speaker.id),
        eventId: session.eventId,
      }
    : {
        name: "",
        description: "",
        url: "",
        startDate: "",
        endDate: "",
        eventId: "",
        speakerIds: [],
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
