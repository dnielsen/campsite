// It's called EventDetails instead of Event because there would
// be a compatibility issue with the JavaScript APIs

export interface EventDetails extends EventResponse {
  sessions: Session[];
}

export interface EventResponse {
  id: string;
  name: string;
  description: string;
  startDate: string;
  endDate: string;
  photo: string;
  organizerName: string;
  // We can later define Address interface.
  // When address is null then the event is online (remote).
  address: string | null;
  speakerIds: string[];
  sessionIds: string[];
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
  sessionIds: string[];
}

export interface Session extends SessionPreview {
  speakers: SpeakerPreview[];
}

export interface SessionPreview {
  id: string;
  name: string;
  startDate: string;
  endDate: string;
  speakerIds: string[];
  description: string;
  url: string;
}

export interface CreateSpeakerInput {
  name: string;
  bio: string;
  headline: string;
  photo: string;
}
