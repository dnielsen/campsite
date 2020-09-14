export interface Speaker {
  id: string;
  name: string;
  photo: string;
  headline: string;
  bio: string;
  sessions: Session[];
}

export interface EventDetails {
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
  speakers: Speaker[];
  sessions: Session[];
}

export interface SessionPreview {
  id: string;
  name: string;
  startDate: string;
  endDate: string;
}

export interface Session extends SessionPreview {
  description: string;
  url: string;
  speakers: Speaker[];
}
