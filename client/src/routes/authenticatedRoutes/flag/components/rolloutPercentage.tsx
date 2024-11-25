import { Alert, TextInput } from "flowbite-react";
import { Fragment } from "react";
import { useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { CgMathPercent } from "react-icons/cg";
import { FaInfoCircle } from "react-icons/fa";
import { HiInformationCircle } from "react-icons/hi2";
import { GoffPercentageProgressBar } from "../../../../components/progressBar/progressBar.tsx";
import { Section } from "../../../../components/section/section.tsx";
import {
  extractValidationErrors,
  hasErrors,
} from "../../../../helpers/extractValidationError.ts";
import type { FeatureFlagFormData } from "../../../../models/featureFlagFormData.ts";
import { getColorByIndex } from "../helpers/colors.ts";
import { isPercentageValid } from "../helpers/percentage.ts";

const translationBaseKey = "component.percentageRollout";
export const PercentageRollout = ({ baseName }: { baseName: string }) => {
  const { t } = useTranslation();
  const {
    register,
    formState: { errors },
    watch,
  } = useFormContext();
  const { variations } = useWatch<FeatureFlagFormData>();
  const numberVariations = variations?.length ?? 0;

  // computePercentages computes the sum of all percentages.
  function computePercentages(baseName: string): number {
    const percentages = watch(`${baseName}.percentage`);
    if (percentages === undefined) {
      return 0;
    }
    let computedPercentage = 0;
    variations?.forEach((variation) => {
      const percentage = variation?.name ? percentages[variation.name] : 0;
      computedPercentage += percentage ?? 0;
    });
    return computedPercentage;
  }

  const errorsToCheck =
    variations?.map((variation) => {
      return `${baseName}.percentage[${variation.name}]`;
    }) ?? [];

  console.log(errorsToCheck);

  const hasVariationErrors = hasErrors(errors, errorsToCheck);

  return (
    <Section
      innerTitle={t(`${translationBaseKey}.title`)}
      info={t(`${translationBaseKey}.description`)}
      maxWidth={"4xl"}
      infoStyle={"gray"}
      borderColour={hasVariationErrors ? "danger" : "base"}
    >
      {numberVariations <= 0 && (
        <Alert color="failure" icon={HiInformationCircle}>
          {t(`${translationBaseKey}.alertVariation`)}
        </Alert>
      )}
      {hasVariationErrors && (
        <Alert color="failure">
          <div className={"flex space-x-2"}>
            <FaInfoCircle className={"h-4 w-4"} />
            <span className="font-medium">
              {t(`${translationBaseKey}.errors.variationErrorTitle`)}
            </span>
          </div>
          <ul className="ml-5 mt-1.5 list-inside list-disc">
            {extractValidationErrors(errors, errorsToCheck).map(
              (error, index) => (
                <li key={index}>{error}</li>
              ),
            )}
          </ul>
        </Alert>
      )}

      {numberVariations > 0 && (
        <>
          <GoffPercentageProgressBar progress={computePercentages(baseName)} />
          {variations?.map((variation, index: number) => (
            <Fragment key={`${baseName}.percentage.${variation.id}`}>
              {variation.name && (
                <div className={"my-3 flex items-center"}>
                  <TextInput
                    {...register(`${baseName}.percentage[${variation.name}]`, {
                      required: t(`${translationBaseKey}.errors.required`),
                      valueAsNumber: true,
                      validate: {
                        validatePercentage: (value) => {
                          if (!isPercentageValid(value)) {
                            return t(
                              `${translationBaseKey}.errors.invalidPercentage`,
                            );
                          }
                        },
                        validateSum: () => {
                          const computedPercentages =
                            computePercentages(baseName);
                          if (
                            !isPercentageValid(computedPercentages) ||
                            computedPercentages !== 100
                          ) {
                            return t(
                              `${translationBaseKey}.errors.invalidPercentageSum`,
                            );
                          }
                        },
                      },
                    })}
                    placeholder="0"
                    defaultValue="0"
                    className={"w-20 appearance-none"}
                    rightIcon={CgMathPercent}
                    inputMode={"numeric"}
                    sizing={"sm"}
                  />
                  <label className={"ml-2"}>
                    {getColorByIndex(index)} {variation.name}
                  </label>
                </div>
              )}
            </Fragment>
          ))}
        </>
      )}
    </Section>
  );
};
