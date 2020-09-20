import React from "react";
import { Field } from "formik";

interface Props {
  value: string;
  label: string;
  name: string;
}

function Checkbox(props: Props) {
  // TODO types
  function handleChange({ field, form }: any) {
    // If it's already in the value array, remove it, otherwise add it.
    const newFieldValue = field.value.includes(props.value)
      ? field.value.filter((value: any) => value !== props.value)
      : field.value.concat(props.value);
    form.setFieldValue(props.name, newFieldValue);
  }

  return (
    <Field name={props.name}>
      {/*TODO types*/}
      {({ field, form }: any) => (
        <label>
          <input
            type="checkbox"
            {...props}
            checked={field.value.includes(props.value)}
            onChange={() => handleChange({ field, form })}
          />
          {props.label}
        </label>
      )}
    </Field>
  );
}

export default Checkbox;
