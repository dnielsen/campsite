import React from "react";
import { SpeakerPreview } from "../../common/interfaces";

// import styled component
import * as s from "../../styles/speakerStyles";
import * as g from "../../styles/globalStyles";

interface Props {
  speaker: SpeakerPreview;
}

function SpeakerPreviewItem(props: Props) {
  return (
    <g.Container>
      <s.SpeakerProfileWrapper>
        <s.SpeakerInfoWrapper>
          <s.SpeakerPhoto>
            <img
              src={props.speaker.photo}
              alt={props.speaker.name}
              className="img-fluid"
            />
          </s.SpeakerPhoto>
          <s.SpeakerName>{props.speaker.name}</s.SpeakerName>
          <s.SpeakerTitle>{props.speaker.headline}</s.SpeakerTitle>
          <s.SpeakerSocialMedia>
            <a href={"https://twitter.com/elonmusk"}>
              <i className="fa fa-twitter twitter" aria-hidden="true"></i>
              Twitter
            </a>
            <a href={"https://linkedin.com"}>
              <i className="fa fa-linkedin linkedin" aria-hidden="true"></i>
              LinkedIn
            </a>
          </s.SpeakerSocialMedia>
          <s.SpeakerBio>
            <h3>About Me</h3>
            <p>{props.speaker.bio}</p>
          </s.SpeakerBio>
        </s.SpeakerInfoWrapper>
        <s.SpeakerSessionWrapper>
          <h2>Session</h2>
          <p>
            What comes next? What are the most innovative developments in Big
            Data storage and query design? Where is the innovation, what should
            be you be trying out and looking at? In this talk I'll cover the
            latest and greatest for the Big Data world - this will include
            in-memory stores such as Aerospike, triplestores such as
            AlgebraixData and Dremel-implementations such as Google Big Query.
            Come to this talk to see these new data stores in action.
          </p>
        </s.SpeakerSessionWrapper>
      </s.SpeakerProfileWrapper>
    </g.Container>
    /* <a href={`/speakers/${props.speaker.id}`}></a> */
  );
}

export default SpeakerPreviewItem;
