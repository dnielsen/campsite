import React from "react";
import { Field, Form, Formik, FormikState, FormikValues } from "formik";
import { EventDetails, SpeakerPreview } from "../../common/interfaces";
import Checkbox from "../../components/Checkbox";
import DateTimeField from "../../components/DateTimeField";
import useSessionForm from "../../hooks/useSessionForm";

// A temporary solution, later we might load speakers and events asynchronously,
// and fetch less data.
interface Props {
  speakers: SpeakerPreview[];
  events: EventDetails[];
}

function SessionForm(props: Props) {
  const { formConfig } = useSessionForm({
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
            <DateTimeField name={"startDate"} />
          </section>
          <section>
            <label htmlFor={"endDate"}>End date</label>
            <DateTimeField name={"endDate"} />
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
