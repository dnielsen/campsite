export interface Speaker {
  id: string;
  name: string;
  photo: string;
  headline: string;
  bio: string;
}

export interface Organizer {
  id: string;
  name: string;
  photo: string;
}

export interface EventInfo {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;
  photo: string;
  organizer: Organizer;
  // We can later define Address interface.
  // When address is null then the event is online (remote).
  address: string | null;
}
