import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  FetchSessionInput,
  FormSessionInput,
  UseForm,
} from "../common/interfaces";
import useSessionSubmit from "./useSessionSubmit";

export default function useCreateSessionForm(): UseForm<FormSessionInput> {
  const onSubmit = useSessionSubmit();

  const initialValues: FormSessionInput = {
    name: "",
    description: "",
    url: "",
    startDate: "",
    endDate: "",
    speakerOptions: [],
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = { onSubmit, validationSchema, initialValues };
  return { formConfig };
}
