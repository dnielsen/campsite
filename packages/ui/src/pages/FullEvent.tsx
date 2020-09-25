import React from "react";
import { EventDetails } from "../common/interfaces";
import useAPI from "../hooks/useAPI";
import util from "../common/util";
import SessionSchedule from "./fullEvent/SessionSchedule";
import { Link, useHistory, useParams } from "react-router-dom";
import SpeakerList from "../components/SpeakerList";
import { BASE_EVENT_API_URL } from "../common/constants";
import * as s from "../styled/eventStyles";
import { Container, Col, Row } from "react-bootstrap";

function FullEvent() {
  const { id } = useParams<{ id: string }>();
  const history = useHistory();
  const { data: eventDetails, loading, error } = useAPI<EventDetails>(
    `/events/${id}`,
  );

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  async function handleDelete() {
    // Send a request to delete the event.
    await fetch(`${BASE_EVENT_API_URL}/${id}`, { method: "DELETE" });
    // Redirect to the home page after deleting the speaker.
    history.push("/");
  }

  if (eventDetails.sessions) {
    const eventSpeakersWithDuplicates = eventDetails.sessions
      .map((session) => session.speakers)
      .flat();

    eventDetails.speakers = util.getUniqueSpeakers(eventSpeakersWithDuplicates);
  }

  return (
    <Container>
      <s.EventWrapper>
        <s.EventMainTitle>
          <h1>Events</h1>
        </s.EventMainTitle>
        <Row>
          <Col lg={4} md={6} sm={12}>
            <s.Event>
              <img
                src={eventDetails.photo}
                alt={eventDetails.name}
                className={"img-fluid"}
              />
              <s.EventContent>
                <s.EventHeading>{eventDetails.name}</s.EventHeading>
                <s.EventTime>
                  <i className={"fa fa-calendar mr-2"} aria-hidden />
                  {util.getFullDateString(eventDetails.startDate)}
                </s.EventTime>
                <s.EventLocation>
                  <i className={"fa fa-map-marker mr-2"} aria-hidden />
                  {/*// For now we'll just use the start date but we might add*/}
                  {/*// support for events that last a few days.*/}
                  {eventDetails.address}
                </s.EventLocation>
                <s.EventOrganizer>
                  <i className="fa fa-user mr-2" aria-hidden />
                  {eventDetails.organizerName}
                </s.EventOrganizer>
                <s.EventDescription>
                  {eventDetails.description}
                </s.EventDescription>
                <s.EventRegister>
                  <a href={"/"}>Register Now</a>
                </s.EventRegister>
              </s.EventContent>
            </s.Event>
            <Link to={`/events/${id}/edit`}>Edit</Link>
            <button type={"button"} onClick={handleDelete}>
              Delete
            </button>
            {eventDetails.sessions && (
              <SessionSchedule sessions={eventDetails.sessions} />
            )}
            {eventDetails.speakers && (
              <SpeakerList speakers={eventDetails.speakers} />
            )}
          </Col>
        </Row>
      </s.EventWrapper>
    </Container>
  );
}

export default FullEvent;
