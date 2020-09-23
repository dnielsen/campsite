// import styled component
import styled from "styled-components";

export const FlexWrapper = styled.div`
  display: flex;

  @media (max-width: 767px) {
    flex-direction: column;
  }
`;

export const EventWrapper = styled.div`
  margin: 70px 0px;
`;

export const EventMainTitle = styled.div`
  text-align: center;
  padding-bottom: 30px;

  h1 {
    color: #282828;
    line-height: 42px;
    font-weight: 700;
    font-size: 32px;
    margin: 0;
  }
`;

export const Event = styled.div`
  flex: 1;
  margin: 0px 15px;
  border: 1px solid #ccc;

  img {
    max-width: 100%;
    width: 100%;
    height: 300px;
  }

  @media (max-width: 767px) {
    margin-bottom: 30px;
  }
`;

export const EventContent = styled.div`
  padding: 20px 10px;
`;

export const EventHeading = styled.h1`
  color: #2faad9;
  font-size: 22px;
  font-weight: 600;
  margin: 0;
`;

export const EventTime = styled.p`
  font-size: 12px;
  color: #777777;
  font-weight: 500;

  i {
    font-size: 18px;
    margin-right: 15px;
  }
`;

export const EventOrganizer = styled.p`
  font-size: 12px;
  color: #777777;
  font-weight: 500;
  margin-bottom: 10px;

  i {
    font-size: 18px;
    margin-right: 15px;
  }
`;

export const EventLocation = styled.p`
  font-size: 12px;
  color: #777777;
  font-weight: 500;
  margin-bottom: 10px;

  i {
    font-size: 18px;
    margin-right: 15px;
  }
`;

export const EventDescription = styled.p`
  font-size: 14px;
  color: #777777;
  font-weight: 500;
  padding-bottom: 20px;
`;

export const EventRegister = styled.div`
  margin: 15px 0px;
  a {
    background: #2faad9;
    color: #fff;
    text-transform: uppercase;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 600;
    box-shadow: 4px 4px 14px 0 rgba(0, 0, 0, 0.2);
    text-decoration: none;
  }
`;
