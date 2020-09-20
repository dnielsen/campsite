import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import { EventDetails, SpeakerPreview } from "../../common/interfaces";
import Checkbox from "../../components/Checkbox";
import useCreateSessionForm from "../../hooks/useCreateSessionForm";

interface Props {
  speakers: SpeakerPreview[];
  events: EventDetails[];
}

function SessionForm(props: Props) {
  const { formConfig } = useCreateSessionForm({
    defaultEventIdValue: props.events[0].id,
  });

  return (
    <Formik {...formConfig}>
      {({ isSubmitting }: FormikState<FormikValues>) => (
        <Form noValidate>
          <section>
            <label htmlFor={"name"}>Name</label>
            <Field type={"text"} name={"name"} />
          </section>
          <section>
            <label htmlFor={"description"}>Description</label>
            <Field type={"text"} name={"description"} />
          </section>
          <section>
            <label htmlFor={"url"}>Url</label>
            <Field type={"text"} name={"url"} />
          </section>
          <section>
            <label htmlFor={"startDate"}>Start date</label>
            <Field type={"date"} name={"startDate"} />
          </section>
          <section>
            <label htmlFor={"endDate"}>End date</label>
            <Field type={"date"} name={"endDate"} />
          </section>
          <section>
            <label htmlFor="eventId">Event</label>
            <Field name={"eventId"} as={"select"}>
              {props.events.map((event) => (
                <option key={event.id} value={event.id}>
                  {event.name}
                </option>
              ))}
            </Field>
          </section>
          <section>
            {props.speakers.map((speaker) => (
              <Checkbox
                key={speaker.id}
                name={"speakerIds"}
                value={speaker.id}
                label={speaker.name}
              />
            ))}
            {/*<SelectField*/}
            {/*  options={options}*/}
            {/*  name={"speakerOptions"}*/}
            {/*  defaultValue={props.formConfig.initialValues.speakerOptions}*/}
            {/*/>*/}
          </section>
          <button type={"submit"} disabled={isSubmitting}>
            Submit
          </button>
        </Form>
      )}
    </Formik>
  );
}

export default SessionForm;
