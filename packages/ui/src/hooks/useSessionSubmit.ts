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

  const fetchOptions: FetchOptions = {
    url: BASE_SESSION_API_URL,
    method: "POST",
  };
  // If `id` is specified then we're updating a session instead of creating one.
  if (id) {
    fetchOptions.method = "PUT";
    fetchOptions.url = `${BASE_SESSION_API_URL}/${id}`;
  }

  async function onSubmit(input: FormSessionInput) {
    // Replace speakerOptions property with speakerIds.
    const fetchInput: FetchSessionInput = {
      ...input,
      speakerIds: input.speakerOptions.map((option: Option) => option.value),
      startDate: new Date(input.startDate),
      endDate: new Date(input.endDate),
    };
    // Send a request to create/update the session.
    const newSession = (await fetch(fetchOptions.url, {
      method: fetchOptions.method,
      body: JSON.stringify(fetchInput),
    }).then((res) => res.json())) as SessionPreview;
    // Redirect to the created session page.
    history.push(`/sessions/${newSession.id}`);
  }

  return onSubmit;
}
