// import styled component
import styled from "styled-components";

export const SpeakerBanner = styled.div`
  flex: 0 0 100%;
  max-width: 100%;

  img {
    max-width: 100%;
    height: auto;
    width: 100%;
    border-radius: 12px;
  }
`;

export const SpeakerProfileWrapper = styled.div`
  background: #f5f5f5;
  padding: 30px 20px;
  border-radius: 8px;
  border: 1px solid #ccc;
  margin: 50px 0px;
`;

export const SpeakerSessionWrapper = styled.div`
  flex: 1;
  position: relative;
  background: #fff;
  padding: 15px;
  border-radius: 8px;
  border: 1px solid #e3e3e3;
  box-shadow: 0px 5px 5px 0 rgba(0, 0, 0, 0.18);

  @media (max-width: 767px) {
    margin: 15px;
  }

  h2 {
    color: #414141;
  }

  p {
    font-size: 14px;
    color: #414141;
  }
`;

export const FlexWrapper = styled.div`
  display: flex;
  flex-wrap: wrap;

  @media (max-width: 767px) {
    flex-direction: column;
  }
`;

export const SpeakerInfoWrapper = styled.div`
  background: #fff;
  padding: 15px;
  border-bottom-left-radius: 8px;
  border-bottom-right-radius: 8px;
  box-shadow: 0px 5px 5px 0 rgba(0, 0, 0, 0.18);
  position: relative;
  margin-right: 15px;
  border: 1px solid #e3e3e3;
  border-radius: 8px;
  flex: 0 0 20%;

  @media (max-width: 767px) {
    margin: 15px;
    display: flex;
  }
`;

export const SpeakerContent = styled.div`
  margin-top: 50px;
  text-align: center;

  @media (max-width: 767px) {
    text-align: left;
    margin: 0px 20px;
  }
`;

export const SpeakerPhoto = styled.div`
  position: absolute;
  top: -100px;
  border: 5px solid #fff;
  border-radius: 50%;
  left: 40px;

  @media (max-width: 767px) {
    position: unset;
  }

  img {
    width: 150px;
    height: 150px;
    border-radius: 50%;
  }
`;

export const SpeakerName = styled.p`
  color: #2faad9;
  font-size: 16px;
  font-weight: 600;
  margin: 5px 0px;
`;

export const SpeakerTitle = styled.p`
  color: #414141;
  font-size: 14px;
  font-weight: 600;
  margin: 5px 0px;
`;

export const SpeakerSocialMedia = styled.div`
  margin: 20px 0px;

  a {
    text-decoration: none;
    font-size: 14px;
    margin-bottom: 10px;
    color: #414141;
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
  color: rgba(26, 26, 26, 0.6);
  font-size: 12px;
`;

export const SpeakerSessionScheduleWrapper = styled.div`
  background: #f5f5f5;
  display: flex;
  justify-content: space-between;
  padding: 15px;
  color: #414141;
  font-size: 14px;
  border: 1px solid #e3e3e3;
  font-weight: 600;
  border-radius: 8px;
`;

export const SessionName = styled.div``;
export const SessionDate = styled.div``;
export const SessionRegister = styled.div``;
export const SessionWatch = styled.div``;
