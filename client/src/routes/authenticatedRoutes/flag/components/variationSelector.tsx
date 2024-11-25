import { Badge, Select } from "flowbite-react";
import { Fragment } from "react";
import { useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import {
  extractValidationError,
  hasError,
} from "../../../../helpers/extractValidationError.ts";
import {
  SELECT_ROLLOUT_PERCENTAGE,
  SELECT_ROLLOUT_PROGRESSIVE,
} from "../../../../models/const.ts";
import type { FeatureFlagFormData } from "../../../../models/featureFlagFormData.ts";
import { getColorByIndex } from "../helpers/colors.ts";
import { PercentageRollout } from "./rolloutPercentage.tsx";
import { ProgressiveRollout } from "./rolloutProgressive.tsx";

const translationBaseKey = "component.variationSelector";
export const VariationSelector = ({
  selectSize,
  baseName,
  disabled = false,
}: {
  baseName: string;
  selectSize?: string;
  disabled?: boolean;
}) => {
  const { t } = useTranslation();
  const {
    register,
    watch,
    formState: { errors },
  } = useFormContext();
  const { variations } = useWatch<FeatureFlagFormData>();

  const numberVariations = variations?.length ?? 0;

  return (
    <div>
      <div className="flex items-center">
        <Badge size="sm" className={"mr-3"} color={"indigo"}>
          {t(`${translationBaseKey}.serve`)}
        </Badge>
        <Select
          sizing={selectSize ?? "md"}
          {...register(`${baseName}.variationSelect`, {
            required: t(
              `${translationBaseKey}.selectRollout.errors.noSelectedVariation`,
            ),
          })}
          disabled={numberVariations === 0 || disabled}
          className={"w-fit min-w-72"}
          defaultValue={""}
          color={
            hasError(errors, `${baseName}.variationSelect`) ? "failure" : "gray"
          }
          helperText={
            extractValidationError(errors, `${baseName}.variationSelect`)
              ?.message
          }
        >
          <option disabled={true} value="">
            Variations
          </option>
          {variations?.map((variation, index: number) => (
            <Fragment key={variation.id}>
              {variation.id && (
                <option value={variation.id}>
                  {getColorByIndex(index)} {variation.name}
                </option>
              )}
            </Fragment>
          ))}
          <option disabled={true}>Rollout Strategy</option>
          <option value={SELECT_ROLLOUT_PERCENTAGE}>
            {t(`${translationBaseKey}.selectRollout.percentage`)}
          </option>
          <option value={SELECT_ROLLOUT_PROGRESSIVE}>
            {t(`${translationBaseKey}.selectRollout.progressive`)}
          </option>
        </Select>
      </div>
      {watch(`${baseName}.variationSelect`) === SELECT_ROLLOUT_PERCENTAGE && (
        <PercentageRollout baseName={baseName} />
      )}

      {watch(`${baseName}.variationSelect`) === SELECT_ROLLOUT_PROGRESSIVE && (
        <ProgressiveRollout baseName={baseName} />
      )}
    </div>
  );
};
