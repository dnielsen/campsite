import React, { useState } from "react";
// import {
//   CreateEventInput,
//   CreateSessionInput,
//   EventDetails,
// } from "../common/interfaces";
// import { BASE_EVENT_API_URL } from "../common/constants";
// import SessionForm from "./createEvent/SessionForm";
// import DatePicker from "react-datepicker";
// import SessionInput from "./createEvent/SessionInput";
//
function CreateEvent() {
  //   const [name, setName] = useState("");
  //   const [description, setDescription] = useState("");
  //   const [startDate, setStartDate] = useState(new Date());
  //   const [endDate, setEndDate] = useState(new Date());
  //   const [address, setAddress] = useState("");
  //   const [organizerName, setOrganizerName] = useState("");
  //   const [photo, setPhoto] = useState("");
  //   const [sessions, setSessions] = useState<CreateSessionInput[]>([]);
  //
  //   async function handleSubmit(e: React.FormEvent) {
  //     e.preventDefault();
  //     console.log(e);
  //     const input: CreateEventInput = {
  //       name,
  //       description,
  //       startDate,
  //       endDate,
  //       organizerName,
  //       address,
  //       photo,
  //       sessions,
  //     };
  //
  //     const createdEvent = (await fetch(BASE_EVENT_API_URL, {
  //       method: "POST",
  //       body: JSON.stringify(input),
  //     }).then((res) => res.json())) as EventDetails;
  //     console.log(createdEvent);
  //   }

  return <div></div>;
  //   return (
  //     <div>
  //       <h2>Create event</h2>
  //       <form onSubmit={handleSubmit}>
  //         <section>
  //           <label htmlFor="eventName">Name</label>
  //           <input
  //             type="text"
  //             name="eventName"
  //             id="eventName"
  //             value={name}
  //             onChange={(event) => setName(event.target.value)}
  //           />
  //         </section>
  //         <section>
  //           <label htmlFor="eventDescription">Description</label>
  //           <input
  //             type="text"
  //             name="eventDescription"
  //             id="eventDescription"
  //             value={description}
  //             onChange={(event) => setDescription(event.target.value)}
  //           />
  //         </section>
  //         <section>
  //           <label htmlFor="eventAddress">Address</label>
  //           <input
  //             type="text"
  //             name="eventAddress"
  //             id="eventAddress"
  //             value={address}
  //             onChange={(event) => setAddress(event.target.value)}
  //           />
  //         </section>
  //         <section>
  //           <label htmlFor="eventOrganizerName">Organizer name</label>
  //           <input
  //             type="text"
  //             name="eventOrganizerName"
  //             id="eventOrganizerName"
  //             value={organizerName}
  //             onChange={(event) => setOrganizerName(event.target.value)}
  //           />
  //         </section>
  //         <section>
  //           <label htmlFor="eventPhoto">Photo</label>
  //           <input
  //             type="text"
  //             name="eventPhoto"
  //             id="eventPhoto"
  //             value={photo}
  //             onChange={(event) => setPhoto(event.target.value)}
  //           />
  //         </section>
  //         <section>
  //           <label htmlFor="eventStartDate">Start date</label>
  //           <DatePicker
  //             id={"eventStartDate"}
  //             selected={startDate}
  //             onChange={(date: Date) => setStartDate(date)}
  //           />
  //         </section>
  //         <section>
  //           <label htmlFor="eventEndDate">End date</label>
  //           <DatePicker
  //             id={"eventEndDate"}
  //             selected={endDate}
  //             onChange={(date: Date) => setEndDate(date)}
  //           />
  //         </section>
  //         <div>
  //           <h3>Add sessions</h3>
  //           {/*<SessionInput />*/}
  //           {/*<SessionForm sessions={sessions} setSessions={setSessions} />*/}
  //         </div>
  //         <button type={"submit"}>Create</button>
  //       </form>
  //     </div>
  //   );
}

export default CreateEvent;
