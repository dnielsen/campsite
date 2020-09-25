import React, { useState } from "react";
import DateTime from "react-datetime";
import "react-datetime/css/react-datetime.css";
import moment, { Moment } from "moment-timezone";
import { Field, FormikProps, FormikValues, useField } from "formik";

interface Props {
  name: string;
}

function DateTimeField(props: Props) {
  return (
    <Field name={props.name}>
      {/*TODO types*/}
      {({ field, form }: any) => (
        <DateTime
          name={props.name}
          {...field}
          onChange={(date) =>
            form.setFieldValue(props.name, moment(date).toDate())
          }
        />
      )}
    </Field>
  );
}

export default DateTimeField;
