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

function getFullDateString(date: Date) {
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

function getValueForDateField(date: string) {
  return moment(date).format("YYYY-MM-DD");
}

export default {
  getHourRangeString,
  getFullDateString,
  getUniqueElementsFromMultidimensionalArray,
  getValueForDateField,
  getUniqueSpeakers,
};
