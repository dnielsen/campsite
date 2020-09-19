import moment from "moment-timezone";
import { Option } from "./interfaces";

const TIMEZONE = "America/Los_Angeles";

function getHourRangeString(startDate: Date, endDate: Date) {
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

export default {
  getHourRangeString,
  getFullDateString,
  getUniqueElementsFromMultidimensionalArray,
};
