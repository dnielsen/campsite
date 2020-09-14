import moment from "moment-timezone";

const TIMEZONE = "America/Los_Angeles";

function getHourRangeString(startDate: string, endDate: string) {
  // for example: 12:55pm
  const startTime = moment(startDate).tz(TIMEZONE).format("h:mma");
  // For example: 12:55pm PDT
  const endTime = moment(endDate).tz(TIMEZONE).format("h:mma z");

  return `${startTime} - ${endTime}`;
}

function getFullDate(date: string) {
  return moment(date).tz(TIMEZONE).format("MM/DD/YYYY h:mma z");
}

function replacer(key: string, value: any) {
  if (value instanceof Date) return value.toString();
  return value;
}

// Converts Date values in an object into strings, otherwise our app would crash.
function serialize(obj: any) {
  return JSON.stringify(obj, replacer);
}

export default { getHourRangeString, serialize, getFullDate };
