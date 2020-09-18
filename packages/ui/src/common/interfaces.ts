// It's called EventDetails instead of Event because there would
// be a compatibility issue with the JavaScript APIs

export interface EventDetails extends EventResponse {
  sessions: Session[];
}

export interface EventResponse {
  id: string;
  name: string;
  description: string;
  startDate: Date;
  endDate: Date;
  photo: string;
  organizerName: string;
  // We can later define Address interface.
  // When address is null then the event is online (remote).
  address: string | null;
}

export interface Speaker extends SpeakerPreview {
  sessions?: Session[];
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

export interface CreateEventInput {
  name: string;
  description: string;
  startDate: Date;
  endDate: Date;
  photo: string;
  organizerName: string;
  address: string;
  sessions: CreateSessionInput[];
}

export interface CreateSessionInput {
  name: string;
  description: string;
  url: string;
  startDate: Date;
  endDate: Date;
  speakers: CreateSpeakerInput[];
}

export interface CreateSpeakerInput {
  name: string;
  bio: string;
  photo: string;
  headline: string;
}
