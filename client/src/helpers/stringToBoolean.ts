import _ from "lodash";

export const stringToBoolean = (value: string): boolean => {
  const areTrue = ["yes", "true", true, "y", 1, "1"];
  if (_.isString(value)) {
    value = value.toLowerCase();
  }
  return _.indexOf(areTrue, value) > -1;
};
