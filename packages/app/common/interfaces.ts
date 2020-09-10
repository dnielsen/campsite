export interface Person {
  id: string;
  name: string;
  photo: string;
  headline: string;
  bio: string;
}

export interface EventInfo {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;
  photo: string;
  organizer: Person;
  // We can later define Address interface.
  // When address is null then the event is online (remote).
  address: string | null;
}

export interface SessionPreview {
  id: string;
  title: string;
  startDate: Date;
  endDate: Date;
}

export interface Session extends SessionPreview {
  id: string;
  description: string;
  speakers: Person[];
}

// Converts startDate and endDate into strings from Date's
export type SerializedSession = Omit<Session, "startDate"> &
  Omit<Session, "endDate"> & { startDate: string; endDate: string };
