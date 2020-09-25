import moment from "moment-timezone";
import { Option, SpeakerPreview } from "./interfaces";

const TIMEZONE = "America/Los_Angeles";

function getHourRangeString(startDate: string, endDate: string) {
  // for example: 12:55pm
  const startTime = moment(startDate).tz(TIMEZONE).format("h:mma");
  // For example: 12:55pm PDT
  const endTime = moment(endDate).tz(TIMEZONE).format("h:mma z");

  return `${startTime} - ${endTime}`;
}

function getFullDateString(date: string) {
  return moment(date).tz(TIMEZONE).format("MM/DD/YYYY h:mma z");
}

function getUniqueElementsFromMultidimensionalArray(arr: any[]) {
  return [...new Set(arr.flat())];
}

function getUniqueSpeakers(speakers: SpeakerPreview[]) {
  return speakers.reduce((uniqueSpeakers, currSpeaker: SpeakerPreview) => {
    if (
      uniqueSpeakers
        .map((speaker: SpeakerPreview) => speaker.id)
        .includes(currSpeaker.id)
    ) {
      return uniqueSpeakers;
    } else {
      return uniqueSpeakers.concat(currSpeaker);
    }
  }, [] as SpeakerPreview[]);
}

// For example: `06/27/2020 5:06 PM`. We need this function
// because `react-datetime` library requires the date formatted this way.
function getDateFormValue(date: string | Date) {
  return moment(date).format("MM/DD/yyyy hh:mm a");
}

export default {
  getHourRangeString,
  getFullDateString,
  getUniqueElementsFromMultidimensionalArray,
  getUniqueSpeakers,
  getDateFormValue,
};
