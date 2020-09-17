import React from "react";
import { EventDetails } from "../../common/interfaces";
import SpeakerList from "../speaker/SpeakerList";
import SessionSchedule from "../session/SessionSchedule";
import util from "../../common/util";

// import bootstrap container
import Container from "react-bootstrap/Container";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";

// import styled component
import * as s from "../../styles/eventStyles";

interface Props {
  eventDetails: EventDetails;
}

function EventItem(props: Props) {
  return (
    <Container>
      <s.EventWrapper>
        <s.EventMainTitle>
          <h1>Events</h1>
        </s.EventMainTitle>
        <Row>
          <Col lg={4} md={6} sm={12}>
            <s.Event>
              {/*For now we'll just use the startDate info*/}
              <img
                src={props.eventDetails.photo}
                alt={props.eventDetails.name}
                className="img-fluid"
              />
              <s.EventContent>
                <s.EventHeading>{props.eventDetails.name}</s.EventHeading>
                <s.EventTime>
                  <i className="fa fa-calendar mr-2" aria-hidden="true"></i>
                  {util.getFullDate(props.eventDetails.startDate)}
                </s.EventTime>
                <s.EventLocation>
                  <i className="fa fa-map-marker mr-2" aria-hidden="true"></i>
                  {props.eventDetails.address}
                </s.EventLocation>
                <s.EventOrganizer>
                  <i className="fa fa-user mr-2" aria-hidden="true"></i>
                  {props.eventDetails.organizerName}
                </s.EventOrganizer>
                <s.EventDescription>
                  {props.eventDetails.description}
                </s.EventDescription>
                <s.EventRegister>
                  <a href={"/"}>Register Now</a>
                </s.EventRegister>{" "}
              </s.EventContent>
            </s.Event>
            {/* <div>
                <SessionSchedule sessions={props.eventDetails.sessions} />
              </div> */}
            {/* <p>Our speakers</p>
              <div>
                <SpeakerList speakers={props.eventDetails.speakers} />
              </div> */}
          </Col>
        </Row>
      </s.EventWrapper>
    </Container>
  );
}

export default EventItem;
