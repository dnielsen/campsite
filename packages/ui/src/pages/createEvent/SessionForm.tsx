import React, { useState } from "react";
import {
  CreateSessionInput,
  CreateSpeakerInput,
} from "../../common/interfaces";
import DatePicker from "react-datepicker";

interface Props {
  sessions: CreateSessionInput[];
  setSessions: (sessions: CreateSessionInput[]) => void;
}

function SessionForm(props: Props) {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [url, setUrl] = useState("");
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [speakers, setSpeakers] = useState<CreateSpeakerInput[]>([]);

  function handleSubmit() {
    const session: CreateSessionInput = {
      name,
      description,
      url,
      startDate,
      endDate,
      speakers,
    };
    props.setSessions(props.sessions.concat(session));
  }
  return (
    <form onSubmit={handleSubmit}>
      <section>
        <label htmlFor="sessionName">Name</label>
        <input
          type="text"
          name="sessionName"
          id="sessionName"
          value={name}
          onChange={(event) => setName(event.target.value)}
        />
      </section>
      <section>
        <label htmlFor="sessionDescription">Description</label>
        <input
          type="text"
          name="sessionDescription"
          id="sessionDescription"
          value={description}
          onChange={(event) => setDescription(event.target.value)}
        />
      </section>
      <section>
        <label htmlFor="sessionUrl">Address</label>
        <input
          type="text"
          name="sessionUrl"
          id="sessionUrl"
          value={url}
          onChange={(event) => setUrl(event.target.value)}
        />
      </section>
      <section>
        <label htmlFor="sessionStartDate">Start date</label>
        <DatePicker
          id={"sessionStartDate"}
          selected={startDate}
          onChange={(date: Date) => setStartDate(date)}
        />
      </section>
      <section>
        <label htmlFor="sessionEndDate">End date</label>
        <DatePicker
          id={"sessionEndDate"}
          selected={endDate}
          onChange={(date: Date) => setEndDate(date)}
        />
      </section>
    </form>
  );
}

export default SessionForm;
