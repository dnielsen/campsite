import React from "react";
import SpeakerItem from "./SpeakerItem";
import { Speaker } from "../common/interfaces";

function SpeakerList() {
  const speakers: Speaker[] = [
    {
      id: "adfsdfs-123-asd-123",
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
    {
      id: "adfsdfs-123-asd-123",
      bio: "I'm John and I like computers",
      headline: "Software Engineer at Google",
      name: "John Doe",
      photo:
        "https://images.unsplash.com/photo-1599748779346-5e30cd4d1635?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1000&q=80",
    },
  ];
  return (
    <div>
      {speakers.map((speaker) => (
        <SpeakerItem key={speaker.id} speaker={speaker} />
      ))}
    </div>
  );
}

export default SpeakerList;
