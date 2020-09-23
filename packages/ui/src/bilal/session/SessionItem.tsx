import React from "react";
import { Session } from "../../common/interfaces";
import SpeakerList from "../speaker/SpeakerList";
import moment from "moment";
import util from "../../common/util";

// import styled component
import * as s from "../styles/sessionStyles";
import * as g from "../styles/globalStyles";

interface Props {
  session: Session;
}

function SessionItem(props: Props) {
  return (
    <g.Container>
      <s.SessionWrapper>
        <s.FlexWrapper>
          <s.SessionListWrapper>
            <h2>Session(s)</h2>
            <ul>
              <li className="active">Intro to Python</li>
              <li>Intro to HTML/CSS/Bootstrap</li>
              <li>Intro to Python Web Framework</li>
              <li>Python Data Structure</li>
              <li>Intro to Python</li>
            </ul>
          </s.SessionListWrapper>
          <s.SessionDetailWrapper>
            <s.SpeakerSessionScheduleWrapper>
              <s.SessionName>Intro to Python</s.SessionName>
              <s.SessionDate>3:08pm - 3:08pm PDT on 09/14/2020</s.SessionDate>
              <s.SessionRegister>Registration</s.SessionRegister>
              <s.SessionWatch>Watch</s.SessionWatch>
            </s.SpeakerSessionScheduleWrapper>
            <h2>Description</h2>
            <p>
              What comes next? What are the most innovative developments in Big
              Data storage and query design? Where is the innovation, what
              should be you be trying out and looking at? In this talk I'll
              cover the latest and greatest for the Big Data world - this will
              include in-memory stores such as Aerospike, triplestores such as
              AlgebraixData and Dremel-implementations such as Google Big Query.
              Come to this talk to see these new data stores in action.
            </p>
            <h2>Speaker (s)</h2>
            <s.FlexWrapper>
              <s.SpeakerContent>
                <s.SpeakerPhoto>
                  <img
                    src="https://png.pngtree.com/thumb_back/fw800/back_our/20190622/ourmid/pngtree-minimal-city-building-banner-background-image_232405.jpg"
                    alt="speaker-img"
                    className="img-fluid"
                  />
                </s.SpeakerPhoto>
                <s.SpeakerName>Warren Smith</s.SpeakerName>
                <s.SpeakerTitle>CEO of Tesla</s.SpeakerTitle>
                <s.SpeakerSocialMedia>
                  <a href={"https://twitter.com/elonmusk"}>
                    <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                  </a>
                  <a href={"https://linkedin.com"}>
                    <i
                      className="fa fa-linkedin linkedin"
                      aria-hidden="true"
                    ></i>
                  </a>
                </s.SpeakerSocialMedia>
              </s.SpeakerContent>
              <s.SpeakerContent>
                <s.SpeakerPhoto>
                  <img
                    src="https://png.pngtree.com/thumb_back/fw800/back_our/20190622/ourmid/pngtree-minimal-city-building-banner-background-image_232405.jpg"
                    alt="speaker-img"
                    className="img-fluid"
                  />
                </s.SpeakerPhoto>
                <s.SpeakerName>Warren Smith</s.SpeakerName>
                <s.SpeakerTitle>CEO of Tesla</s.SpeakerTitle>
                <s.SpeakerSocialMedia>
                  <a href={"https://twitter.com/elonmusk"}>
                    <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                  </a>
                  <a href={"https://linkedin.com"}>
                    <i
                      className="fa fa-linkedin linkedin"
                      aria-hidden="true"
                    ></i>
                  </a>
                </s.SpeakerSocialMedia>
              </s.SpeakerContent>
              <s.SpeakerContent>
                <s.SpeakerPhoto>
                  <img
                    src="https://png.pngtree.com/thumb_back/fw800/back_our/20190622/ourmid/pngtree-minimal-city-building-banner-background-image_232405.jpg"
                    alt="speaker-img"
                    className="img-fluid"
                  />
                </s.SpeakerPhoto>
                <s.SpeakerName>Warren Smith</s.SpeakerName>
                <s.SpeakerTitle>CEO of Tesla</s.SpeakerTitle>
                <s.SpeakerSocialMedia>
                  <a href={"https://twitter.com/elonmusk"}>
                    <i className="fa fa-twitter twitter" aria-hidden="true"></i>
                  </a>
                  <a href={"https://linkedin.com"}>
                    <i
                      className="fa fa-linkedin linkedin"
                      aria-hidden="true"
                    ></i>
                  </a>
                </s.SpeakerSocialMedia>
              </s.SpeakerContent>
            </s.FlexWrapper>
          </s.SessionDetailWrapper>
        </s.FlexWrapper>
      </s.SessionWrapper>
    </g.Container>
    // <div>
    //   <div>
    //     <div>
    //       <SpeakerList speakers={props.session.speakers} />
    //     </div>
    //     <div>
    //       <div>
    //         <p>{props.session.name}</p>
    //         <p>
    //           {util.getHourRangeString(
    //             props.session.startDate,
    //             props.session.endDate,
    //           )}{" "}
    //           on {moment(props.session.startDate).format("MM/DD/YYYY")}
    //         </p>
    //         <a href={props.session.url}>{props.session.url}</a>
    //         <p>{props.session.description}</p>
    //       </div>
    //     </div>
    //   </div>
    // </div>
  );
}

export default SessionItem;
