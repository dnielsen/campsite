import React from "react";
import { SpeakerPreview } from "../../common/interfaces";

// import styled component
import * as s from "../../styles/speakerStyles";
import * as g from "../../styles/globalStyles";

interface Props {
  speaker: any;
}

function SpeakerPreviewItem(props: Props) {
  return (
    <g.Container>
      <s.SpeakerProfileWrapper>
        <s.SpeakerBanner>
          <img
            src="https://png.pngtree.com/thumb_back/fw800/back_our/20190622/ourmid/pngtree-minimal-city-building-banner-background-image_232405.jpg"
            alt="banner-image"
          />
        </s.SpeakerBanner>
        <s.FlexWrapper>
          <s.SpeakerInfoWrapper>
            <s.SpeakerPhoto>
              <img
                src={props.speaker.photo}
                alt={props.speaker.name}
                className="img-fluid"
              />
            </s.SpeakerPhoto>
            <s.SpeakerContent>
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
            </s.SpeakerContent>
          </s.SpeakerInfoWrapper>
          <s.SpeakerSessionWrapper>
            <h2>Session(s)</h2>
            <s.SpeakerSessionScheduleWrapper>
              <s.SessionName>Intro to Python</s.SessionName>
              <s.SessionDate>3:08pm - 3:08pm PDT on 09/14/2020</s.SessionDate>
              <s.SessionRegister>Registration</s.SessionRegister>
              <s.SessionWatch>Watch</s.SessionWatch>
            </s.SpeakerSessionScheduleWrapper>
            <p>
              What comes next? What are the most innovative developments in Big
              Data storage and query design? Where is the innovation, what
              should be you be trying out and looking at? In this talk I'll
              cover the latest and greatest for the Big Data world - this will
              include in-memory stores such as Aerospike, triplestores such as
              AlgebraixData and Dremel-implementations such as Google Big Query.
              Come to this talk to see these new data stores in action.
            </p>
            <s.SpeakerSessionScheduleWrapper>
              <s.SessionName>Intro to Python</s.SessionName>
              <s.SessionDate>3:08pm - 3:08pm PDT on 09/14/2020</s.SessionDate>
              <s.SessionRegister>Registration</s.SessionRegister>
              <s.SessionWatch>Watch</s.SessionWatch>
            </s.SpeakerSessionScheduleWrapper>
            <p>
              What comes next? What are the most innovative developments in Big
              Data storage and query design? Where is the innovation, what
              should be you be trying out and looking at? In this talk I'll
              cover the latest and greatest for the Big Data world - this will
              include in-memory stores such as Aerospike, triplestores such as
              AlgebraixData and Dremel-implementations such as Google Big Query.
              Come to this talk to see these new data stores in action.
            </p>
            <s.SpeakerSessionScheduleWrapper>
              <s.SessionName>Intro to Python</s.SessionName>
              <s.SessionDate>3:08pm - 3:08pm PDT on 09/14/2020</s.SessionDate>
              <s.SessionRegister>Registration</s.SessionRegister>
              <s.SessionWatch>Watch</s.SessionWatch>
            </s.SpeakerSessionScheduleWrapper>
            <p>
              What comes next? What are the most innovative developments in Big
              Data storage and query design? Where is the innovation, what
              should be you be trying out and looking at? In this talk I'll
              cover the latest and greatest for the Big Data world - this will
              include in-memory stores such as Aerospike, triplestores such as
              AlgebraixData and Dremel-implementations such as Google Big Query.
              Come to this talk to see these new data stores in action.
            </p>
          </s.SpeakerSessionWrapper>
        </s.FlexWrapper>
      </s.SpeakerProfileWrapper>
    </g.Container>
    /* <a href={`/speakers/${props.speaker.id}`}></a> */
  );
}

export default SpeakerPreviewItem;
