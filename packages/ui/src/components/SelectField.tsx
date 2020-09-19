import React from "react";
import { Option } from "../common/interfaces";
import Select, { ValueType } from "react-select";
import { useField } from "formik";

interface Props {
  options: Option[];
  name: string;
  defaultValue?: Option[];
}

function SelectField(props: Props) {
  const [field, meta, helpers] = useField({
    name: props.name,
    type: "select",
  });

  // `options` is of type Option[] but somehow types don't match, so we'll
  // keep it at that for now.
  function handleChange(options: ValueType<Option>) {
    console.log(options);
    helpers.setValue(options);
  }

  return (
    <Select
      isMulti
      options={props.options}
      onChange={handleChange}
      defaultValue={props.defaultValue}
    />
  );
}

export default SelectField;
