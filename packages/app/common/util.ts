import moment from "moment";

function getHourRangeString(startDate: Date, endDate: Date) {
  const AMERICAN_HOUR_FORMAT = "LT";

  const startHour = moment(startDate).format(AMERICAN_HOUR_FORMAT);
  const endHour = moment(endDate).format(AMERICAN_HOUR_FORMAT);

  return `${startHour} - ${endHour}`;
}

function replacer(key: string, value: any) {
  if (value instanceof Date) return value.toString();
  return value;
}

// Converts Date values in an object into strings, otherwise our app would crash.
function serialize(obj: any) {
  return JSON.stringify(obj, replacer);
}

export default { getHourRangeString, serialize };
