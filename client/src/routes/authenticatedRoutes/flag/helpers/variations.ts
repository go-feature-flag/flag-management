import { t } from "i18next";
import _ from "lodash";
import type { RegisterOptions } from "react-hook-form";
import { AiOutlineFieldString } from "react-icons/ai";
import { BsToggleOn } from "react-icons/bs";
import type { IconType } from "react-icons/lib";
import { TbNumbers } from "react-icons/tb";
import { VscJson } from "react-icons/vsc";
import { DEFAULT_VALUE_NEW_JSON_VARIATION } from "../../../../models/const.ts";
import type { SelectedVariationType } from "../../../../models/featureFlagFormData.ts";

const translationBaseKey = "page.flag.variations";

export const variationTypes: {
  name: string; // technical name used in the code
  displayName: string; // used to be printed on the screen
  icon: IconType;
  type: SelectedVariationType;
  maxVariations?: number;
  // inputValueType: it is the type of the field to display the value.
  // "json" is a custom type to use the JSON editor.
  inputValueType: "text" | "number" | "json";
  defaultValueForNewVariation?: string;
  // step is used for the number input field.
  step?: string | number;
  // valueNotEditable is used to disable the value field modification (useful for boolean).
  valueNotEditable?: boolean;
  valueValidationOptions?: RegisterOptions;
}[] = [
  {
    name: "Boolean",
    type: "boolean",
    displayName: "Bool",
    icon: BsToggleOn,
    maxVariations: 2,
    inputValueType: "text",
    valueNotEditable: true,
    valueValidationOptions: {
      setValueAs: (value: boolean | string) => {
        if (_.isBoolean(value)) {
          return value;
        }
        return value === "true";
      },
      validate: (value: boolean | string) => {
        if (value !== true && value !== false) {
          return t(`${translationBaseKey}.fields.error.valueRequired`);
        }
      },
    },
  },
  {
    name: "String",
    type: "string",
    displayName: "String",
    icon: AiOutlineFieldString,
    inputValueType: "text",
    valueValidationOptions: {
      required: t(`${translationBaseKey}.fields.error.valueRequired`),
    },
  },
  {
    name: "Integer",
    type: "integer",
    displayName: "Integer",
    icon: TbNumbers,
    inputValueType: "number",
    step: 1,
    valueValidationOptions: {
      required: t(`${translationBaseKey}.fields.error.valueRequired`),
      valueAsNumber: true,
    },
  },
  {
    name: "Double",
    type: "double",
    displayName: "Float",
    icon: TbNumbers,
    inputValueType: "number",
    step: "any",
    valueValidationOptions: {
      required: t(`${translationBaseKey}.fields.error.valueRequired`),
      valueAsNumber: true,
    },
  },
  {
    name: "JSON",
    type: "json",
    displayName: "JSON",
    icon: VscJson,
    inputValueType: "json",
    defaultValueForNewVariation: DEFAULT_VALUE_NEW_JSON_VARIATION,
    valueValidationOptions: {
      required: t(`${translationBaseKey}.fields.error.valueRequired`),
    },
  },
];
