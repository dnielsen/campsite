import React, { Fragment } from "react";
import { EventDetails } from "../../common/interfaces";
import util from "../../common/util";

// import styled component
import * as s from "../styles/eventStyles";
import * as g from "../styles/globalStyles";

interface Props {
  eventDetails: EventDetails;
}

function EventItem(props: Props) {
  return (
    <Fragment>
      <g.Container>
        <s.EventWrapper>
          <s.EventMainTitle>
            <h1>Events</h1>
          </s.EventMainTitle>
          <s.FlexWrapper>
            <s.Event>
              <img
                src={props.eventDetails.photo}
                alt={props.eventDetails.name}
                className="img-fluid"
              />
              <s.EventContent>
                <s.EventHeading>{props.eventDetails.name}</s.EventHeading>
                <s.EventTime>
                  <i className="fa fa-calendar mr-2" aria-hidden="true"></i>
                  {util.getFullDateString(props.eventDetails.startDate)}
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
            <s.Event>
              <img
                src={props.eventDetails.photo}
                alt={props.eventDetails.name}
                className="img-fluid"
              />
              <s.EventContent>
                <s.EventHeading>{props.eventDetails.name}</s.EventHeading>
                <s.EventTime>
                  <i className="fa fa-calendar mr-2" aria-hidden="true"></i>
                  {util.getFullDateString(props.eventDetails.startDate)}
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
            <s.Event>
              <img
                src={props.eventDetails.photo}
                alt={props.eventDetails.name}
                className="img-fluid"
              />
              <s.EventContent>
                <s.EventHeading>{props.eventDetails.name}</s.EventHeading>
                <s.EventTime>
                  <i className="fa fa-calendar mr-2" aria-hidden="true"></i>
                  {util.getFullDateString(props.eventDetails.startDate)}
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
          </s.FlexWrapper>
        </s.EventWrapper>
      </g.Container>
    </Fragment>
  );
}

export default EventItem;
