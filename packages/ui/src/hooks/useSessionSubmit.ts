import {
  FetchSessionInput,
  FormSessionInput,
  Option,
  SessionPreview,
} from "../common/interfaces";
import { BASE_SESSION_API_URL } from "../common/constants";
import { useHistory } from "react-router-dom";

interface FetchOptions {
  url: string;
  method: "PUT" | "POST";
}

export default function useSessionSubmit(
  id?: string,
): (input: FormSessionInput) => void {
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

    if (id) {
      await edit(fetchInput);
    } else {
      await create(fetchInput);
    }
  }

  async function create(input: FetchSessionInput) {
    // Send a request to create the session.
    const createdSession = (await fetch(`${BASE_SESSION_API_URL}`, {
      method: "POST",
      body: JSON.stringify(input),
    }).then((res) => res.json())) as SessionPreview;
    // Redirect to the created session page.
    history.push(`/sessions/${createdSession.id}`);
  }

  async function edit(input: FetchSessionInput) {
    // Send a request to edit the session.
    await fetch(`${BASE_SESSION_API_URL}/${id}`, {
      method: "PUT",
      body: JSON.stringify(input),
    });
    // Redirect to the edited session page.
    history.push(`/sessions/${id}`);
  }
  return onSubmit;
}
