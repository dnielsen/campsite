import React from "react";
import { EventDetails } from "../common/interfaces";
import useAPI from "../hooks/useAPI";
import { useHistory, useParams } from "react-router-dom";
import { BASE_EVENT_API_URL } from "../common/constants";
import EventItem from "../bilal/components/event/EventItem";

function FullEvent() {
  const { id } = useParams<{ id: string }>();
  const history = useHistory();
  const { data: eventDetails, loading, error } = useAPI<EventDetails>(
    `/events/${id}`,
  );

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  return <EventItem eventDetails={eventDetails} />;

  async function handleClick() {
    await fetch(`${BASE_EVENT_API_URL}/${id}`, { method: "DELETE" });
    // Redirect to the home page after deleting the speaker.
    history.push("/");
  }
  //
  // if (eventDetails.sessions) {
  //   const eventSpeakersWithDuplicates = eventDetails.sessions
  //     .map((session) => session.speakers)
  //     .flat();
  //
  //   eventDetails.speakers = util.getUniqueSpeakers(eventSpeakersWithDuplicates);
  // }
  //
  // return (
  //   <Container>
  //     <s.EventWrapper>
  //       <s.EventMainTitle>
  //         <h1>Events</h1>
  //       </s.EventMainTitle>
  //       <Row>
  //         <Col lg={4} md={6} sm={12}>
  //           <s.Event>
  //             <img
  //               src={eventDetails.photo}
  //               alt={eventDetails.name}
  //               className={"img-fluid"}
  //             />
  //             <s.EventContent>
  //               <s.EventHeading>{eventDetails.name}</s.EventHeading>
  //               <s.EventTime>
  //                 <i className={"fa fa-calendar mr-2"} aria-hidden />
  //                 {util.getFullDateString(eventDetails.startDate)}
  //               </s.EventTime>
  //               <s.EventLocation>
  //                 <i className={"fa fa-map-marker mr-2"} aria-hidden />
  //                 {/*// For now we'll just use the start date but we might add*/}
  //                 {/*// support for events that last a few days.*/}
  //                 {eventDetails.address}
  //               </s.EventLocation>
  //               <s.EventOrganizer>
  //                 <i className="fa fa-user mr-2" aria-hidden />
  //                 {eventDetails.organizerName}
  //               </s.EventOrganizer>
  //               <s.EventDescription>
  //                 {eventDetails.description}
  //               </s.EventDescription>
  //               <s.EventRegister>
  //                 <a href={"/"}>Register Now</a>
  //               </s.EventRegister>
  //             </s.EventContent>
  //           </s.Event>
  //           <button type={"button"} onClick={handleClick}>
  //             Delete
  //           </button>
  //           {eventDetails.sessions && (
  //             <SessionSchedule sessions={eventDetails.sessions} />
  //           )}
  //           {eventDetails.speakers && (
  //             <SpeakerList speakers={eventDetails.speakers} />
  //           )}
  //         </Col>
  //       </Row>
  //     </s.EventWrapper>
  //   </Container>
  // );
}

export default FullEvent;
