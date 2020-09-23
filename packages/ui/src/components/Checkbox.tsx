import React from "react";
import { Field, FieldProps, FormikProps } from "formik";

interface Props {
  value: string;
  label: string;
  name: string;
}

function Checkbox(props: Props) {
  // TODO types
  function handleChange({ field, form }: FieldProps<string[]>) {
    // If it's already in the value array, remove it, otherwise add it.
    const newFieldValue = field.value.includes(props.value)
      ? field.value.filter((value) => value !== props.value)
      : field.value.concat(props.value);
    form.setFieldValue(props.name, newFieldValue);
  }

  return (
    <Field name={props.name}>
      {(fieldProps: FieldProps) => (
        <label>
          <input
            type="checkbox"
            {...props}
            checked={fieldProps.field.value.includes(props.value)}
            onChange={() => handleChange(fieldProps)}
          />
          {props.label}
        </label>
      )}
    </Field>
  );
}

export default Checkbox;
