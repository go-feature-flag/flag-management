import { t } from "i18next";
import {
  SELECT_ROLLOUT_PERCENTAGE,
  SELECT_ROLLOUT_PROGRESSIVE,
} from "../../../../models/const.ts";
import type {
  FormDataVariationSelector,
  FormFlagPageTargetingRule,
  FormFlagPageVariationInfo,
} from "../../../../models/featureFlagFormData.ts";

export const isVariationRemovable = (
  currentId: string,
  variationName: string,
  variationValues: FormFlagPageVariationInfo[],
  defaultSelection: FormDataVariationSelector,
  targetingRules: FormFlagPageTargetingRule[] = [],
): string | undefined => {
  if (variationValues.length <= 2)
    return t("page.flag.variations.tooltip.removeMinVariations");

  const isDefaultSelection = defaultSelection.variationSelect === currentId;
  const isDefaultProgressive = [
    defaultSelection.progressive?.variation.from,
    defaultSelection.progressive?.variation.to,
  ].includes(currentId);
  if (isDefaultSelection || isDefaultProgressive)
    return t("page.flag.variations.tooltip.removeVariationUsedDefaultRule");

  return targetingRules.some((rule) => {
    const isRuleSelection = rule.variation.variationSelect === currentId;
    const isRuleProgressive =
      rule.variation.variationSelect === SELECT_ROLLOUT_PROGRESSIVE &&
      [
        rule.variation.progressive?.variation.from,
        rule.variation.progressive?.variation.to,
      ].includes(currentId);

    const isRulePercentage =
      rule.variation.variationSelect === SELECT_ROLLOUT_PERCENTAGE &&
      Object.entries(rule.variation.percentage ?? {})
        .filter(([, value]) => value !== 0)
        .map(([key]) => key)
        .includes(variationName ?? "");

    return isRuleSelection || isRuleProgressive || isRulePercentage;
  })
    ? t("page.flag.variations.tooltip.removeVariationUsedTargetingRule")
    : undefined;
};
