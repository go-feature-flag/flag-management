import { v4 as uuidv4 } from "uuid";

export const SELECT_ROLLOUT_PERCENTAGE = "rollout_percentage";
export const SELECT_ROLLOUT_PROGRESSIVE = "rollout_progressive";
export const DEFAULT_VALUE_NEW_JSON_VARIATION = "{}";
export const VARIATION_STRING_DEFAULT_VALUE = [
  { name: "variation-A", value: "", id: uuidv4() },
  { name: "variation-B", value: "", id: uuidv4() },
];
export const VARIATION_BOOLEAN_DEFAULT_VALUE = [
  { name: "enabled", value: true, id: uuidv4() },
  { name: "disabled", value: false, id: uuidv4() },
];
export const VARIATION_JSON_DEFAULT_VALUE = [
  {
    name: "variation-A",
    value: DEFAULT_VALUE_NEW_JSON_VARIATION,
    id: uuidv4(),
  },
  {
    name: "variation-B",
    value: DEFAULT_VALUE_NEW_JSON_VARIATION,
    id: uuidv4(),
  },
];
export const VARIATION_INTEGER_DEFAULT_VALUE = [
  { name: "variation-A", value: 0, id: uuidv4() },
  { name: "variation-B", value: 0, id: uuidv4() },
];
export const VARIATION_DOUBLE_DEFAULT_VALUE = [
  { name: "variation-A", value: 1.1, id: uuidv4() },
  { name: "variation-B", value: 1.1, id: uuidv4() },
];
