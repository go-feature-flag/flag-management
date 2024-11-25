import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import type { typeVariation } from "../../../../@types/typeVariation.ts";
import { Section } from "../../../../components/section/section.tsx";
import { VariationSelector } from "../components/variationSelector.tsx";
import { variationTypes } from "../helpers/variations.ts";
import { Variations } from "./variations.tsx";

/**
 * VariationsTab is the component displaying the variation tab.
 * @constructor
 */
const translationBaseKey = "page.flag.variations";
export const VariationTab = ({ type }: { type: typeVariation }) => {
  const { t } = useTranslation();

  const variationType = variationTypes.find((i) => i.type === type);
  return (
    <>
      <h1 className={"my-2"}>Feature Flag Variations</h1>
      <TypeSelector type={type} />
      <Section
        outsideTitle={t(`${translationBaseKey}.section.variations`)}
        titleSize={"lg"}
      >
        {variationType && (
          <Variations
            key={`${variationType.name}_${variationType.type}`}
            fieldName={`variations`}
            maxVariations={variationType.maxVariations}
            inputValueType={variationType.inputValueType}
            step={variationType.step}
            notEditable={variationType.valueNotEditable}
            valueValidationOptions={variationType.valueValidationOptions}
          />
        )}
      </Section>
      <Section
        outsideTitle={t(`${translationBaseKey}.section.defaultRule`)}
        info={t(`${translationBaseKey}.defaultRule.info`)}
        titleSize={"lg"}
      >
        <VariationSelector baseName={`defaultRule`} />
      </Section>
    </>
  );
};

const TypeSelector = ({ type }: { type: typeVariation }) => {
  const { register } = useFormContext();
  return (
    <ul className="mt-3 grid w-fit grid-cols-2 gap-6 md:grid-cols-5">
      {variationTypes
        .filter((variation) => variation.type === type)
        .map((variation) => (
          <li key={variation.name}>
            <input
              {...register("type")}
              type="radio"
              id={variation.name}
              value={variation.name.toLowerCase()}
              className="peer hidden"
              required
              defaultChecked={variation.type === type}
            />
            <label
              htmlFor={variation.name}
              className="inline-flex w-full items-center justify-between rounded-lg border border-gray-200 bg-white p-4 text-gray-500 hover:bg-gray-100 hover:text-gray-600 peer-checked:border-goff-600 peer-checked:text-goff-600 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-gray-300 dark:peer-checked:text-goff-500"
            >
              <div className="block">
                <div className="w-full text-base font-semibold">
                  {variation.displayName}
                </div>
              </div>
              <variation.icon className="mx-2 h-7 w-7" />
            </label>
          </li>
        ))}
    </ul>
  );
};
