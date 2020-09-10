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
