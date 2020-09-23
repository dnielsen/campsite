import React from "react";
import { useHistory, useParams } from "react-router-dom";
import useAPI from "../hooks/useAPI";
import { EventDetails } from "../common/interfaces";
import EventItem from "../bilal/event/EventItem";
import HomePage from "../bilal/home";

function Events() {
  return <HomePage />;
}

export default Events;
