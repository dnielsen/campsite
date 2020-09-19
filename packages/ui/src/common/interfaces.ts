import * as Yup from "yup";

// It's called EventDetails instead of Event because there would
// be a compatibility issue with the JavaScript APIs.
export interface EventDetails {
  id: string;
  name: string;
  description: string;
  startDate: Date;
  endDate: Date;
  photo: string;
  organizerName: string;
  // We might add null address support later - that is, when an event takes place online
  // instead of in-person.
  address: string;
  sessions: Session[];
}

export interface Speaker extends SpeakerPreview {
  sessions: Session[];
}

export interface SpeakerPreview {
  id: string;
  name: string;
  photo: string;
  headline: string;
  bio: string;
}

export interface Session extends SessionPreview {
  speakers: SpeakerPreview[];
}

export interface SessionPreview {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;
  description: string;
  url: string;
}

export interface BaseCreateEventInput {
  name: string;
  description: string;
  photo: string;
  organizerName: string;
  address: string;
}

export interface CreateEventFormInput extends BaseCreateEventInput {
  sessionOptions: Option[];
  startDate: string;
  endDate: string;
}

export interface CreateEventFetchInput extends BaseCreateEventInput {
  sessionIds: string[];
  startDate: Date;
  endDate: Date;
}

export interface BaseCreateSessionInput {
  name: string;
  description: string;
  url: string;
}

export interface CreateSessionFetchInput extends BaseCreateSessionInput {
  speakerIds: string[];
  startDate: Date;
  endDate: Date;
}

export interface CreateSessionFormInput extends BaseCreateSessionInput {
  speakerOptions: Option[];
  startDate: string;
  endDate: string;
}

export interface SpeakerInput {
  name: string;
  bio: string;
  photo: string;
  headline: string;
}

export interface Option {
  label: string;
  value: string;
}

export interface FormConfig<T> {
  onSubmit: (input: T) => void;
  initialValues: T;
  validationSchema: Yup.ObjectSchema;
  enableReinitialize?: boolean;
}
