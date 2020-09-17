// import styled component
import styled from "styled-components";

export const SpeakerProfileWrapper = styled.div`
  background: #f5f6f8;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #ccc;
  display: flex;
  flex-wrap: wrap;
`;

export const SpeakerSessionWrapper = styled.div`
  flex: 0 0 75%;
  max-width: 70%;
  position: relative;
  width: 100%;
  background: #fff;
  padding: 15px;
  border-radius: 8px;
  border: 1px solid #e3e3e3;
  box-shadow: 0px 5px 5px 0 rgba(0, 0, 0, 0.18);

  h2 {
    color: #414141;
  }

  p {
    font-size: 14px;
    color: #414141;
  }
`;

export const SpeakerInfoWrapper = styled.div`
  background: #fff;
  padding: 15px;
  border-radius: 8px;
  border: 1px solid #e3e3e3;
  box-shadow: 0px 5px 5px 0 rgba(0, 0, 0, 0.18);
  flex: 0 0 20%;
  max-width: 20%;
  position: relative;
  width: 100%;
  margin-right: 15px;
`;

export const SpeakerPhoto = styled.div`
  text-align: center;

  img {
    width: 150px;
    height: 150px;
    border-radius: 50%;
  }
`;

export const SpeakerName = styled.p`
  text-align: center;

  color: #2faad9;
  font-size: 16px;
  font-weight: 600;
  margin: 5px 0px;
`;

export const SpeakerTitle = styled.p`
  color: #414141;
  font-size: 14px;
  font-weight: 600;
  text-align: center;
  margin: 5px 0px;
`;

export const SpeakerSocialMedia = styled.div`
  text-align: center;
  margin: 20px 0px;

  a {
    text-decoration: none;
    font-size: 14px;
    margin-bottom: 10px;
    color: #1a2857;
    font-weight: 500;
    margin-right: 15px;
  }

  i.twitter {
    color: #34c4f2;
    font-size: 20px;
    margin-right: 7px;
  }

  i.linkedin {
    color: #0274b3;
    font-size: 20px;
    margin-right: 7px;
  }
`;

export const SpeakerBio = styled.div`
  color: #6d737b;
  font-size: 12px;
`;
