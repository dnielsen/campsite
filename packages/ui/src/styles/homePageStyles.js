// import styled component
import styled from "styled-components";

export const BannerSection = styled.div`
  padding: 70px 0px;
  p {
    color: #2faad9;
    font-size: 20px;
    font-weight: 700;
    margin: 0;
  }

  h1 {
    font-size: 72px;
    color: #282828;
    font-weight: 700;
    line-height: 1;
    margin: 0;
    margin-bottom: 30px;
  }

  h3 {
    font-size: 32px;
    margin: 0;
    color: rgba(40, 40, 40, 0.73);
  }

  @media (max-width: 767px) {
    text-align: center;
  }
`;

export const SelectSeatButton = styled.div`
  margin-bottom: 50px;
  @media (max-width: 767px) {
    text-align: center;
  }

  a {
    background: #2faad9;
    padding: 20px 50px;
    box-shadow: inset 0 0 20px 20px transparent;
    transition: box-shadow 0.2s ease;
    font-size: 20px;
    text-align: center;
    border-radius: 8px;
    color: #fff;
    text-decoration: none;
    font-weight: 600;
  }
`;

export const LearnFromBestLogos = styled.div`
  p {
    color: rgba(40, 40, 40, 0.73);
    font-size: 14px;
    font-weight: 600;

    @media (max-width: 767px) {
      text-align: center;
    }
  }
`;

export const BrandImages = styled.div`
  display: flex;

  @media (max-width: 767px) {
    justify-content: center;
  }

  img {
    display: block;
    margin: 12px 10px;
    opacity: 0.5;
    width: 100px;
  }
`;

export const SectionFutureSummit = styled.div`
  margin-top: 100px;
  margin-right: 40px;
  display: flex;
  flex-wrap: wrap;

  @media (max-width: 767px) {
    flex-direction: column;
  }
`;

export const ContentLeftWrapper = styled.div`
  flex: 1;

  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }

  p {
    padding-top: 24px;
    color: rgba(40, 40, 40, 0.73);
    line-height: 26px;
    font-weight: 500;
  }

  button {
    background: #2faad9;
    border: none;
    color: #fff;
    padding: 15px 20px;
    font-weight: 600;
    border-radius: 8px;
    font-size: 14px;
    outline: none;
  }
`;

export const ContentSubTitle = styled.p`
  color: #2faad9 !important;
  font-size: 20px;
  font-weight: 700 !important;
  margin: 0;
`;

export const VideoWrapper = styled.div`
  flex: 1;
`;

export const SectionLargeText = styled.div`
  margin-top: 120px;
  margin-bottom: 120px;
  width: 100%;
  max-width: 964px;
  margin-right: auto;
  margin-left: auto;
`;

export const TextBlack = styled.h3`
  font-size: 34px;
  color: #12141d;
  margin: 0;
  font-weight: 700;
`;

export const TextBlue = styled.div`
  font-size: 34px;
  color: #2faad9;
  margin: 0;
  font-weight: 700;
`;

export const TextGrey = styled.div`
  font-size: 34px;
  color: rgba(40, 40, 40, 0.7);
  margin: 0;
  font-weight: 700;
`;

export const SectionWhatYouWillGet = styled.div`
  margin-top: 30px;

  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }
`;

export const Boxes = styled.div`
  display: flex;
  margin-top: 40px;
  flex-wrap: wrap;

  @media (max-width: 767px) {
    flex-direction: column;
  }
`;

export const HundredPercentLive = styled.div`
  padding: 20px;
  justify-content: center;
  border-radius: 6px;
  background-color: rgba(18, 20, 29, 0.04);
  text-align: center;
  flex: 1;

  p {
    padding-left: 0;
    color: rgba(40, 40, 40, 0.73);
    font-size: 14px;
    line-height: 24px;
    font-weight: 500;
  }

  h4 {
    color: #282828;
    font-size: 24px;
  }
`;

export const GoodVibes = styled.div`
  padding: 20px;
  justify-content: center;
  border-radius: 6px;
  text-align: center;
  flex: 1;

  p {
    padding-left: 0;
    color: rgba(40, 40, 40, 0.73);
    font-size: 14px;
    line-height: 24px;
    font-weight: 500;
  }

  h4 {
    color: #282828;
    font-size: 24px;
  }
`;

export const NoBullShit = styled.div`
  padding: 20px;
  justify-content: center;
  border-radius: 6px;
  background-color: rgba(18, 20, 29, 0.04);
  text-align: center;
  flex: 1;

  p {
    padding-left: 0;
    color: rgba(40, 40, 40, 0.73);
    font-size: 14px;
    line-height: 24px;
    font-weight: 500;
  }

  h4 {
    color: #282828;
    font-size: 24px;
  }
`;

export const CircleImageWrapper = styled.div`
  width: 96px;
  height: 96px;
  border-radius: 48px;
  background: #2faad9;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
`;

export const SectionSpeakersContainer = styled.div`
  margin-top: 40px;

  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }
`;

export const SpeakersWrapper = styled.div`
  margin-top: 60px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  @media (max-width: 767px) {
    flex-direction: column;
  }
`;

export const Speaker = styled.div`
  padding: 20px;
  text-align: center;
  flex: 0 0 20%;

  img {
    width: 100px;
    height: 100px;
    border-style: 1px solid rgba(0, 0, 0, 0.08);
    object-fit: fill;
    border-radius: 60px;
  }

  h4 {
    font-size: 20px;
    line-height: 30px;
    color: #282828;
    margin: 10px 0px;
  }
`;

export const TextLead = styled.div`
  color: rgba(40, 40, 40, 0.73);
  font-size: 14px;
  font-weight: 600;
}`;

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

export const SectionAgenda = styled.div`
  margin-top: 70px;
  text-align: center;
  margin-bottom: 70px;

  h1 {
    color: #282828;
    font-size: 40px;
  }
`;

export const TabsWrapper = styled.div`
  ul {
    display: flex;
    padding: 0;
    margin: 0;
    list-style: none;
  }

  ul li {
    background: #2faad9;
    color: #fff;
    padding: 10px 15px;
    border-radius: 8px;
    font-weight: 600;
    outline: none;
    margin-right: 12px;
    margin-bottom: 12px;
    cursor: pointer;
  }
`;

export const TabContent = styled.div`
  background: #2faad9;
  text-align: left;
  padding: 20px;
  border-radius: 8px;
  color: #fff;
`;

export const EventTitle = styled.div``;

export const EventTime = styled.div``;

export const PanelDiscussion = styled.div`
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  margin-bottom: 14px;
  padding-bottom: 14px;
  border-bottom: 1px solid hsla(0, 0%, 100%, 0.1);
`;

export const topic = styled.div`
  font-weight: 500;
`;

export const PanelImages = styled.div`
  margin-top: 10px;
  img {
    width: 50px;
    height: 50px;
    border-radius: 60px;
  }
`;

export const TimeLimit = styled.div`
  font-weight: 500;
`;

export const SectionIsThisForMe = styled.div`
  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }
`;

export const FounderAndCeo = styled.div`
  flex: 1;
  padding: 0px 36px;

  h4 {
    color: #282828;
  }

  p {
    color: rgba(40, 40, 40, 0.81);
    font-size: 16px;
    line-height: 26px;
  }
`;

export const Badges = styled.div`
  display: flex;
  flex-wrap: wrap;
`;

export const Badge = styled.div`
  background: #282828;
  height: 32px;
  min-width: 32px;
  color: #fff;
  border-radius: 50px;
  padding: 2px 15px;
  font-weight: 600;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 14px;
  margin-right: 10px;
`;

export const ContentIsThisForMe = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 50px;

  @media (max-width: 767px) {
    flex-direction: column;
  }
`;

export const ButtonSaveYourSeatNow = styled.div`
  text-align: center;
  button {
    background: #2faad9;
    border: none;
    color: #fff;
    font-size: 20px;
    padding: 20px 220px;
    border-radius: 8px;
    font-weight: 600;
    outline: none;

    @media (max-width: 767px) {
      padding: 20px;
    }
  }
`;

export const AboutOrganizer = styled.div`
  flex: 1;

  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }

  p {
    color: rgba(40, 40, 40, 0.73);
    line-height: 26px;
  }
`;

export const Partners = styled.div`
  text-align: center;

  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }
`;

export const Footer = styled.div`
  h3 {
    color: #2faad9;
  }

  i.instagram {
    color: #c13584;
    font-size: 20px;
    margin-right: 7px;
  }
`;

export const FrequentlyQA = styled.div`
  text-align: center;

  h3 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
  }
`;

export const Question = styled.div`
  background: #f7f7f7;
  padding: 12px 24px;
  color: #414141;
  text-align: left;
  border-radius: 12px;
`;

export const Answer = styled.div`
  color: rgba(40, 40, 40, 0.73);
  line-height: 26px;
  margin-top: 20px;
`;

export const KeepInMindText = styled.p`
  font-weight: 600;
  margin: 0;
`;

export const OrganizerImage = styled.div`
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;

  img {
    width: 200px;
    height: 200px;
    border-radius: 50%;

    @media (max-width: 767px) {
      margin-bottom: 30px;
    }
  }
`;

export const SectionAboutOrganizer = styled.div`
  display: flex;
  flex-wrap: wrap;
  margin-top: 100px;
  margin-bottom: 100px;

  @media (max-width: 767px) {
    flex-direction: column-reverse;
  }
`;

export const PoweredBy = styled.div`
  font-size: 14px;
  margin-top: 30px;
  margin-bottom: 30px;

  a {
    color: #2faad9;
    font-weight: 600;
    text-decoration: none;
  }

  .privacy-policy {
    display: flex;
    flex-direction: column;
    font-size: 10px;
    line-height: 12px;
    font-weight: 500;
    text-decoration: underline;
  }
`;
