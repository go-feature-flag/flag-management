import { Alert, Select, TextInput } from "flowbite-react";
import type { ReactNode } from "react";
import { Fragment } from "react";
import { Controller, useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { CgMathPercent } from "react-icons/cg";
import { CiCalendarDate, CiCircleChevDown } from "react-icons/ci";
import { FaInfoCircle } from "react-icons/fa";
import {
  TbCircleNumber1,
  TbCircleNumber2,
  TbCircleNumber3,
} from "react-icons/tb";
import DateTimePicker from "react-tailwindcss-datetimepicker";
import { twMerge } from "tailwind-merge";
import { Section } from "../../../../components/section/section.tsx";
import {
  extractValidationErrors,
  hasErrors,
} from "../../../../helpers/extractValidationError.ts";
import type { FeatureFlagFormData } from "../../../../models/featureFlagFormData.ts";
import { getColorByIndex } from "../helpers/colors.ts";
import { isPercentageValid } from "../helpers/percentage.ts";

const translationBaseKey = "component.progressiveRollout";

export const ProgressiveRollout = ({ baseName }: { baseName: string }) => {
  const { t } = useTranslation();
  const componentName = `${baseName}.progressive`;
  return (
    <Section
      innerTitle={t(`${translationBaseKey}.title`)}
      info={t(`${translationBaseKey}.description`)}
      maxWidth={"4xl"}
      infoStyle={"gray"}
    >
      <div className="grid grid-cols-12">
        <SelectVariationsStep componentName={componentName} />
        <ArrowRolloutStep />
        <GoffDatetimePicker baseName={componentName} />
        <ArrowRolloutStep />
        <SelectPercentagesStep componentName={componentName} />
      </div>
    </Section>
  );
};

const SelectVariationsStep = ({ componentName }: { componentName: string }) => {
  const { t } = useTranslation();
  const {
    formState: { errors },
  } = useFormContext();
  const errorsToCheck = [
    `${componentName}.variation.from`,
    `${componentName}.variation.to`,
  ];
  const hasVariationErrors = hasErrors(errors, errorsToCheck);
  return (
    <RolloutStep hasError={hasVariationErrors}>
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
      <div className={"flex flex-wrap items-center gap-2"}>
        <TbCircleNumber1 className="h-9 w-9" />{" "}
        {t(`${translationBaseKey}.step1.part1`)}
        <VariationSelect
          name={`${componentName}.variation.from`}
          requiredErrorMsg={`${translationBaseKey}.errors.variationFromRequired`}
        />{" "}
        {t(`${translationBaseKey}.step1.part2`)}
        <VariationSelect
          name={`${componentName}.variation.to`}
          requiredErrorMsg={`${translationBaseKey}.errors.variationToRequired`}
        />
      </div>
    </RolloutStep>
  );
};

const SelectPercentagesStep = ({
  componentName,
}: {
  componentName: string;
}) => {
  const {
    register,
    formState: { errors },
  } = useFormContext();
  const { t } = useTranslation();

  const errorsToCheck = [
    `${componentName}.percentage.from`,
    `${componentName}.percentage.to`,
  ];

  const hasPercentageErrors = hasErrors(errors, errorsToCheck);

  return (
    <RolloutStep hasError={hasPercentageErrors}>
      {hasPercentageErrors && (
        <Alert color="failure">
          <div className={"flex space-x-2"}>
            <FaInfoCircle className={"h-4 w-4"} />
            <span className="font-medium">
              {t(`${translationBaseKey}.errors.percentageErrorTitle`)}
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
      <div className={"flex flex-wrap items-center gap-2 "}>
        <TbCircleNumber3 className="h-9 w-9" />
        {t(`${translationBaseKey}.step3.part1`)}
        <TextInput
          {...register(`${componentName}.percentage.from`, {
            required: t(`${translationBaseKey}.errors.percentageFromRequired`),
            valueAsNumber: true,
            validate: {
              validatePercentage: (value) => {
                if (!isPercentageValid(value)) {
                  return t(
                    `${translationBaseKey}.errors.invalidPercentageFrom`,
                  );
                }
              },
            },
          })}
          placeholder="0"
          defaultValue="0"
          className={"w-20 appearance-none"}
          rightIcon={CgMathPercent}
          sizing={"sm"}
        />
        {t(`${translationBaseKey}.step3.part2`)}
        <TextInput
          {...register(`${componentName}.percentage.to`, {
            required: t(`${translationBaseKey}.errors.percentageToRequired`),
            valueAsNumber: true,
            validate: {
              validatePercentage: (value) => {
                if (!isPercentageValid(value)) {
                  return t(`${translationBaseKey}.errors.invalidPercentageTo`);
                }
              },
            },
          })}
          placeholder="100"
          defaultValue="100"
          className={"w-20 appearance-none"}
          rightIcon={CgMathPercent}
          sizing={"sm"}
        />
      </div>
    </RolloutStep>
  );
};

const GoffDatetimePicker = ({ baseName }: { baseName: string }) => {
  const { t } = useTranslation();
  const { setValue, control } = useFormContext();
  const now = new Date();
  const defaultRange = {
    start: new Date(),
    end: new Date(now.getTime() + 24 * 60 * 60 * 1000),
  };

  function handleApply(startDate: Date, endDate: Date) {
    setValue(`${baseName}.range`, { start: startDate, end: endDate });
  }

  const proposedRange: Record<string, [Date, Date]> = {
    "12 hours": [new Date(), new Date(now.getTime() + 12 * 60 * 60 * 1000)],
    "1 day": [new Date(), new Date(now.getTime() + 24 * 60 * 60 * 1000)],
    "1 week": [new Date(), new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000)],
    "2 weeks": [
      new Date(),
      new Date(now.getTime() + 2 * 7 * 24 * 60 * 60 * 1000),
    ],
    "4 weeks": [
      new Date(),
      new Date(now.getTime() + 4 * 7 * 24 * 60 * 60 * 1000),
    ],
  };

  return (
    <RolloutStep>
      <Controller
        name={`${baseName}.range`}
        control={control}
        defaultValue={defaultRange}
        render={({ field }) => {
          return (
            <DateTimePicker
              classNames={{
                rootContainer: "max-w-full w-full",
              }}
              ranges={proposedRange}
              smartMode={true}
              start={field?.value.start ?? defaultRange.start}
              end={field.value.end ?? defaultRange.end}
              applyCallback={handleApply}
              years={[now.getFullYear(), now.getFullYear() + 10]}
              autoApply={true}
              twelveHoursClock={true}
            >
              <div className={"flex items-center gap-2"}>
                <TbCircleNumber2 className="h-9 w-9" />{" "}
                {t(`${translationBaseKey}.step2.part1`)}
                <div
                  className={
                    "flex min-w-72 cursor-auto flex-row items-center gap-2 rounded-lg border border-gray-300 p-2 text-xs text-gray-900 focus:border-blue-600 dark:border-gray-600 dark:bg-gray-800 dark:text-white dark:focus:border-blue-500"
                  }
                >
                  {field.value?.start?.toISOString()}
                  <span
                    className={"font-bold text-goff-700 dark:text-goff-400"}
                  >
                    {" "}
                    {t(`${translationBaseKey}.step2.part2`)}{" "}
                  </span>
                  {field.value?.end?.toISOString()}
                  <CiCalendarDate className={"ml-3 h-5 w-5"} />
                </div>{" "}
                {t(`${translationBaseKey}.step2.part3`)}
              </div>
            </DateTimePicker>
          );
        }}
      />
    </RolloutStep>
  );
};

const VariationSelect = ({
  name,
  requiredErrorMsg,
}: {
  name: string;
  requiredErrorMsg: string;
}) => {
  const { register } = useFormContext();
  const { variations } = useWatch<FeatureFlagFormData>();
  const numberVariations = variations?.length ?? 0;
  return (
    <Select
      {...register(name, {
        required: requiredErrorMsg,
      })}
      disabled={numberVariations === 0}
      className={"m-0 w-fit p-0"}
      sizing={"sm"}
    >
      {variations?.map((variation, index: number) => (
        <Fragment key={`${name}.${variation.id}`}>
          {variation.id && (
            <option value={variation.id}>
              {getColorByIndex(index)} {variation.name}
            </option>
          )}
        </Fragment>
      ))}
    </Select>
  );
};

const RolloutStep = ({
  children,
  hasError = false,
}: {
  children: ReactNode;
  hasError?: boolean;
}) => {
  const errorClass = hasError ? "border-solid border-red-500" : "";

  return (
    <div
      className={twMerge(
        "col-span-12 rounded-lg border border-dotted border-gray-500 p-4",
        errorClass,
      )}
    >
      {children}
    </div>
  );
};

const ArrowRolloutStep = () => {
  return (
    <div
      className={
        "col-span-10 col-start-2 my-3 flex items-center justify-center"
      }
    >
      <CiCircleChevDown className="h-9 w-9" />
    </div>
  );
};
