import { EventDetails, Speaker, Session, SessionPreview } from "./interfaces";

const eventInfo: EventDetails = {
  id: "asd123-das-asd",
  name: "The Big Data Event",
  photo:
    "https://images.unsplash.com/photo-1593642634367-d91a135587b5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1650&q=80",
  startDate: new Date("11-11-2019"),
  endDate: new Date("12-12-2021"),
  organizer: {
    id: "324k-dsf",
    bio: "bio",
    headline: "headline",
    name: "John Smith",
    photo:
      "https://images.unsplash.com/photo-1599701834133-9ae09fcfa601?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80",
  },
  address: null,
};

const people: Speaker[] = [
  {
    id: "123",
    bio: "I'm John and I like computers",
    headline: "Software Engineer at Google",
    name: "John Doe",
    photo:
      "https://images.unsplash.com/photo-1599748779346-5e30cd4d1635?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80",
  },
  {
    id: "adfsdfgfs-123-asd-123",
    bio: "I'm John and I like computers",
    headline: "Software Engineer at Google",
    name: "John Doe",
    photo:
      "https://images.unsplash.com/photo-1599748779346-5e30cd4d1635?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80",
  },
  {
    id: "adfsdfs-123-asd-123",
    bio: "I'm John and I like computers",
    headline: "Software Engineer at Google",
    name: "John Doe",
    photo:
      "https://images.unsplash.com/photo-1599748779346-5e30cd4d1635?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80",
  },
];

const sessions: Session[] = [
  {
    id: "23423-dfksk",
    title: "Session Title",
    description: "Session description.",
    startDate: new Date("11-11-2015"),
    endDate: new Date("12-12-2017"),
    speakers: people.slice(2),
  },
];

const sessionPreviews: SessionPreview[] = sessions.map((s) => ({
  id: s.id,
  title: s.title,
  startDate: s.startDate,
  endDate: s.endDate,
}));

export default { eventInfo, people, sessions, sessionPreviews };
