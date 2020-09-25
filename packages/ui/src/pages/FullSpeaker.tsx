import React from "react";
import useAPI from "../hooks/useAPI";
import { Speaker } from "../common/interfaces";
import { Link, useHistory, useParams } from "react-router-dom";
import { BASE_SPEAKER_API_URL } from "../common/constants";
import * as g from "../styled/globalStyles";
import * as s from "../styled/speakerStyles";
import util from "../common/util";
import moment from "moment";

function FullSpeaker() {
  const { id } = useParams<{ id: string }>();
  const history = useHistory();

  const { data: speaker, loading, error } = useAPI<Speaker>(`/speakers/${id}`);

  if (loading) return <div>loading...</div>;
  if (error) return <div>something went wrong: {error.message}</div>;

  async function handleClick() {
    await fetch(`${BASE_SPEAKER_API_URL}/${id}`, { method: "DELETE" });
    // Redirect to the home page after deleting the speaker.
    history.push("/");
  }

  return (
    <g.Container>
      <s.SpeakerProfileWrapper>
        <s.SpeakerInfoWrapper>
          <s.SpeakerPhoto>
            <img src={speaker.photo} alt={speaker.name} className="img-fluid" />
          </s.SpeakerPhoto>
          <s.SpeakerName>{speaker.name}</s.SpeakerName>
          <s.SpeakerTitle>{speaker.headline}</s.SpeakerTitle>
          <s.SpeakerSocialMedia>
            <a href={"https://twitter.com/elonmusk"}>
              <i className="fa fa-twitter twitter" aria-hidden />
              Twitter
            </a>
            <a href={"https://linkedin.com"}>
              <i className="fa fa-linkedin linkedin" aria-hidden />
              LinkedIn
            </a>
          </s.SpeakerSocialMedia>
          <s.SpeakerBio>
            <h3>About Me</h3>
            <p>{speaker.bio}</p>
          </s.SpeakerBio>
          <Link to={`/speakers/${id}/edit`}>Edit</Link>
          <button onClick={handleClick}>Delete</button>
        </s.SpeakerInfoWrapper>
        <s.SpeakerSessionWrapper>
          <s.SpeakerSessionScheduleWrapper>
            <h2>Session Schedule</h2>
            <table>
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Time & Date</th>
                  <th>Link</th>
                </tr>
              </thead>
              <tbody>
                {speaker.sessions?.map((session) => (
                  <tr key={session.id}>
                    <td>
                      <Link to={`/sessions/${session.id}`}>{session.name}</Link>
                    </td>
                    <td>
                      {util.getHourRangeString(
                        session.startDate,
                        session.endDate,
                      )}{" "}
                      on {moment(session.startDate).format("MM/DD/YYYY")}
                    </td>
                    <td>
                      <a href={session.url}>View</a>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </s.SpeakerSessionScheduleWrapper>
        </s.SpeakerSessionWrapper>
      </s.SpeakerProfileWrapper>
    </g.Container>
  );
}

export default FullSpeaker;
