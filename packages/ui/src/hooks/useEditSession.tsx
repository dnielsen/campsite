import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  UseEditForm,
  FormSessionInput,
  Session,
  FormConfig,
} from "../common/interfaces";
import { BASE_SESSION_API_URL } from "../common/constants";
import useAPI from "./useAPI";
import { useEffect, useState } from "react";
import moment from "moment-timezone";
import util from "../common/util";
import useSessionSubmit from "./useSessionSubmit";

export default function useEditSessionForm(
  id: string,
): UseEditForm<FormSessionInput> {
  const onSubmit = useSessionSubmit(id);
  const [initialValues, setInitialValues] = useState<FormSessionInput>({
    name: "",
    description: "",
    url: "",
    startDate: "",
    endDate: "",
    speakerOptions: [],
  });
  const { data: uneditedSession, loading, error } = useAPI<Session>(
    `/sessions/${id}`,
  );

  useEffect(() => {
    // Load the session values so that the session fields aren't empty.
    if (uneditedSession) {
      console.log(uneditedSession);
      setInitialValues({
        name: uneditedSession.name,
        description: uneditedSession.description,
        url: uneditedSession.url,
        startDate: util.getValueForDateField(uneditedSession.startDate),
        endDate: util.getValueForDateField(uneditedSession.endDate),
        speakerOptions: uneditedSession.speakers.map((speaker) => ({
          value: speaker.id,
          label: speaker.name,
        })),
      });
    }
  }, [uneditedSession]);

  const validationSchema = Yup.object().shape({});

  const formConfig: FormConfig<FormSessionInput> = {
    onSubmit,
    validationSchema,
    initialValues,
    enableReinitialize: true,
  };

  return { formConfig, loading, error };
}
