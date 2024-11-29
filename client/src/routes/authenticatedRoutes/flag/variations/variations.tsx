import { Button, FloatingLabel, Tooltip } from "flowbite-react";
import { highlight, languages } from "prismjs";
import "prismjs/components/prism-json.js";
import "prismjs/themes/prism.css";
import type { HTMLInputTypeAttribute } from "react";
import type { RegisterOptions } from "react-hook-form";
import { Controller, useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { FaTrashAlt } from "react-icons/fa";
import { FaCirclePlus } from "react-icons/fa6";
import Editor from "react-simple-code-editor";
import { v4 as uuidv4 } from "uuid";
import type { typeVariation } from "../../../../@types/typeVariation.ts";
import {
  extractValidationError,
  hasError,
} from "../../../../helpers/extractValidationError.ts";
import type {
  FeatureFlagFormData,
  FormDataVariationSelector,
  FormFlagPageTargetingRule,
  FormFlagPageVariationInfo,
} from "../../../../models/featureFlagFormData.ts";
import { getColorByIndex } from "../helpers/colors.ts";
import { isVariationRemovable } from "../helpers/removeVariationHelper.ts";

const translationBaseKey = "page.flag.variations";

/**
 * Variations is the component displaying the list of variations for a feature flag.
 * @param fieldName  => the name of the field in the form
 * @param maxVariations => the maximum number of variations
 * @param inputValueType => the type of the input value
 * @param defaultValueForNewVariation => the default value for a new variation
 * @param step => the step for the input value
 * @param lockValue => if we disabled the value field modification (useful for boolean)
 * @constructor
 */
export const Variations = ({
  fieldName,
  maxVariations = undefined,
  inputValueType,
  defaultValueForNewVariation = "",
  step,
  notEditable = false,
  valueValidationOptions = {},
}: {
  fieldName: string;
  maxVariations?: number;
  inputValueType: "text" | "number" | "json";
  step?: string | number;
  defaultValueForNewVariation?: string;
  notEditable?: boolean;
  valueValidationOptions?: RegisterOptions;
}) => {
  const { t } = useTranslation();
  const { variations, targetingRules, defaultRule } =
    useWatch<FeatureFlagFormData>();
  const { setValue } = useFormContext();

  const variationValues = (): FormFlagPageVariationInfo[] => {
    return (variations as FormFlagPageVariationInfo[]) ?? [];
  };

  const displayAddButton = (): boolean => {
    if (maxVariations) {
      return variationValues()?.length < maxVariations;
    }
    return true;
  };

  function removeItem(index: number) {
    if (variations) {
      const filteredVariations = variations?.filter((_, i) => i !== index);
      setValue(fieldName, filteredVariations);
    }
  }

  function addItem() {
    if (variations) {
      const v = variations ?? [];
      v.push({ name: "", value: defaultValueForNewVariation, id: uuidv4() });
      setValue(fieldName, v);
    }
  }

  const removeTooltipText = (id: string, variationName: string) =>
    isVariationRemovable(
      id,
      variationName,
      variationValues(),
      defaultRule as FormDataVariationSelector,
      (targetingRules as FormFlagPageTargetingRule[]) ?? [],
    );

  return (
    <>
      {variationValues().map((variation, index) => (
        <VariationLine
          name={`${fieldName}[${index}]`}
          key={index}
          removeMessageError={removeTooltipText(
            variation.id ?? "",
            variation.name,
          )}
          defaultValue={variation.value}
          defaultName={variation.name}
          index={index}
          onClickRemoveButton={removeItem}
          inputValueType={inputValueType}
          step={step}
          notEditable={notEditable}
          valueValidationOptions={valueValidationOptions}
        />
      ))}
      {displayAddButton() && (
        <Button onClick={addItem} className={"max-w-xs"}>
          <FaCirclePlus className={"mr-2 h-4 w-4"} />
          {t("page.flag.addButton")}
        </Button>
      )}
    </>
  );
};

/**
 * VariationLine is the component displaying a single variation line.
 * @param removeMessageError => error message if not deletable variation, else undefined
 * @param index => the index of the variation
 * @param notEditable => if the variation is not editable
 * @param name => the name of the variation
 * @param onClickRemoveButton => the function to call when the remove button is clicked
 * @param inputValueType => the type of the input value
 * @param step => the step for the input value
 * @param lockValue => if we disabled the value field modification (useful for boolean)
 * @constructor
 */
const VariationLine = ({
  removeMessageError,
  index,
  notEditable = false,
  name,
  onClickRemoveButton,
  inputValueType = "text",
  step,
  valueValidationOptions = {},
}: {
  removeMessageError: string | undefined;
  defaultValue?: typeVariation;
  defaultName?: string;
  index: number;
  notEditable?: boolean;
  name: string; // the name for react-hook-form
  onClickRemoveButton?: (index: number) => void;
  inputValueType?: HTMLInputTypeAttribute;
  step?: string | number;
  valueValidationOptions?: RegisterOptions;
}) => {
  const { t } = useTranslation();
  const {
    register,
    control,
    setValue,
    formState: { errors },
  } = useFormContext();
  const dotColor = getColorByIndex(index).emoji;

  return (
    <>
      <div className={"grid grid-cols-12 gap-4"}>
        <div className={"col-span-6"}>
          <div className={"grid grid-cols-12 gap-2"}>
            <div
              className={
                "invisible flex items-center justify-center sm:visible"
              }
            >
              <span className={"text-2xl"}>{dotColor}</span>
            </div>
            <div className={"col-span-11"}>
              <FloatingLabel
                {...register(`${name}.name`, {
                  required: t(
                    `${translationBaseKey}.fields.error.nameRequired`,
                  ),
                })}
                variant="outlined"
                label={t(`${translationBaseKey}.fields.name`)}
                placeholder={t(`${translationBaseKey}.fields.name`)}
                color={hasError(errors, `${name}.name`) ? "error" : "default"}
                helperText={
                  extractValidationError(errors, `${name}.name`)?.message
                }
              />
            </div>
          </div>
        </div>
        <div className={"col-span-6"}>
          <div className={"grid grid-cols-12 gap-2"}>
            <div className={"col-span-11"}>
              {inputValueType !== "json" && (
                <FloatingLabel
                  {...register(`${name}.value`, valueValidationOptions)}
                  variant="outlined"
                  label={t(`${translationBaseKey}.fields.value`)}
                  placeholder={t(`${translationBaseKey}.fields.value`)}
                  className={
                    notEditable
                      ? "bg-gray-200 text-gray-500 dark:bg-gray-700 dark:text-gray-400"
                      : ""
                  }
                  disabled={notEditable}
                  type={inputValueType}
                  step={step}
                  color={
                    hasError(errors, `${name}.value`) ? "error" : "default"
                  }
                  helperText={
                    extractValidationError(errors, `${name}.value`)?.message
                  }
                />
              )}
              {inputValueType === "json" && (
                <Controller
                  name={`${name}.value`}
                  control={control}
                  render={({ field }) => (
                    <>
                      <Editor
                        {...field}
                        value={field.value || "{}"}
                        onValueChange={(code) => setValue(field.name, code)}
                        highlight={(code) =>
                          highlight(code, languages.json, "JSON")
                        }
                        padding={10}
                        className={
                          hasError(errors, `${name}.value`)
                            ? "min-h-12 rounded-lg border border-red-500 dark:border-red-500 dark:bg-gray-800"
                            : "min-h-12 rounded-lg border border-gray-200 dark:border-gray-600 dark:bg-gray-800"
                        }
                        style={{
                          fontFamily: '"Fira code", "Fira Mono", monospace',
                          fontSize: 14,
                        }}
                      />
                      <p
                        id="outlined_helper_text:r29:"
                        className="mt-2 text-xs text-red-600 dark:text-red-400"
                      >
                        {
                          extractValidationError(errors, `${name}.value`)
                            ?.message
                        }
                      </p>
                    </>
                  )}
                />
              )}
            </div>
            <div className={"mt-4"}>
              <Tooltip
                content={
                  removeMessageError ??
                  t(`${translationBaseKey}.tooltip.removeOk`)
                }
              >
                <button
                  type="button"
                  disabled={removeMessageError !== undefined}
                  className={
                    removeMessageError === undefined
                      ? "text-gray-800 dark:text-gray-50"
                      : "text-gray-300 dark:text-gray-600"
                  }
                  onClick={
                    onClickRemoveButton
                      ? () => onClickRemoveButton(index)
                      : () => {
                          /*nothing todo here*/
                        }
                  }
                >
                  <FaTrashAlt className="h-4 w-4" />
                </button>
              </Tooltip>
            </div>
          </div>
        </div>
        <div className={"col-span-5"}></div>
      </div>
    </>
  );
};
