// import styled component
import styled from "styled-components";

export const EventWrapper = styled.div`
  margin: 70px 0px;
`;

export const EventMainTitle = styled.div`
  text-align: center;
  padding-bottom: 30px;

  h1 {
    font-weight: bold;
    text-transform: uppercase;
  }
`;

export const Event = styled.div`
  box-shadow: 4px 4px 14px 0 rgba(0, 0, 0, 0.18);

  img {
    width: 350px;
    height: 160px;
  }
`;

export const EventContent = styled.div`
  padding: 20px;
`;

export const EventHeading = styled.h1`
  color: #414141;
  font-size: 22px;
  font-weight: 600;
`;

export const EventTime = styled.p`
  font-size: 12px;
  color: #777777;
  font-weight: 500;
  margin-bottom: 10px;

  i {
    font-size: 18px;
  }
`;

export const EventOrganizer = styled.p`
  font-size: 12px;
  color: #777777;
  font-weight: 500;
  margin-bottom: 10px;

  i {
    font-size: 18px;
  }
`;

export const EventLocation = styled.p`
  font-size: 12px;
  color: #777777;
  font-weight: 500;
  margin-bottom: 10px;

  i {
    font-size: 18px;
  }
`;

export const EventDescription = styled.p`
  font-size: 14px;
  color: #777777;
  font-weight: 500;
`;

export const EventRegister = styled.div`
  a {
    background: #0d0e0e;
    color: #fff;
    text-transform: uppercase;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 600;
    box-shadow: 4px 4px 14px 0 rgba(0, 0, 0, 0.2);
    text-decoration: none;
  }
`;
