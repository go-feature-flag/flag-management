import { v4 as uuidv4 } from "uuid";
import {
  SELECT_ROLLOUT_PERCENTAGE,
  SELECT_ROLLOUT_PROGRESSIVE,
} from "../models/const.ts";
import type { FeatureFlagDTO, Rule } from "../models/featureFlagDTO.ts";
import type {
  FeatureFlagFormData,
  FormDataVariationSelector,
  FormFlagPageVariationInfo,
} from "../models/featureFlagFormData.ts";

export const dtoToFlagFormConverter = (
  flag: FeatureFlagDTO,
): FeatureFlagFormData => {
  const convertedVariations = flag?.variations
    ? Object.entries(flag.variations).map(([key, value]) => ({
        name: key ?? "",
        value: value ?? "",
        id: uuidv4(),
      }))
    : [];
  return {
    name: flag.name,
    id: flag.id,
    disable: flag.disable ?? false,
    trackEvent: flag.trackEvents ?? false,
    description: flag.description,
    version: flag.version,
    type: flag.type,
    creationDate: flag.createdDate ? new Date(flag.createdDate) : new Date(),
    lastUpdatedDate: flag.lastUpdatedDate
      ? new Date(flag.lastUpdatedDate)
      : new Date(),
    variations: convertedVariations,
    metadata: flag?.metadata
      ? Object.entries(flag.metadata).map(([key, value]) => {
          return {
            name: key ?? "",
            value: value.toString() ?? "",
            id: uuidv4(),
          };
        })
      : [],
    defaultRule: convertRuleToFlagForm(flag.defaultRule, convertedVariations),
    targetingRules:
      flag.targeting?.map((r) => {
        return {
          disabled: r.disable ?? false,
          query: r.query ?? "",
          name: r.name ?? "",
          id: r.id ?? uuidv4(),
          variation: convertRuleToFlagForm(r, convertedVariations),
        };
      }) ?? [],
  };
};

const convertRuleToFlagForm = (
  rule: Rule,
  variations: FormFlagPageVariationInfo[],
): FormDataVariationSelector => {
  const currentFormVariationSelector: FormDataVariationSelector = {
    id: rule.id,
  };
  if (rule.variation && rule.variation.length > 0) {
    currentFormVariationSelector.variationSelect = variations.find((v) => {
      return v.name === rule.variation;
    })?.id;
    return currentFormVariationSelector;
  }

  if (rule.percentage) {
    currentFormVariationSelector.percentage = rule.percentage;
    currentFormVariationSelector.variationSelect = SELECT_ROLLOUT_PERCENTAGE;
    return currentFormVariationSelector;
  }

  if (rule.progressiveRollout) {
    currentFormVariationSelector.variationSelect = SELECT_ROLLOUT_PROGRESSIVE;
    currentFormVariationSelector.progressive = {
      variation: {
        from:
          variations.find(
            (r) => r.name === rule.progressiveRollout?.initial?.variation,
          )?.id ?? "",
        to:
          variations.find(
            (r) => r.name === rule.progressiveRollout?.end?.variation,
          )?.id ?? "",
      },
      range: {
        start: rule.progressiveRollout?.initial?.date
          ? new Date(rule.progressiveRollout.initial.date)
          : new Date(),
        end: rule.progressiveRollout?.end?.date
          ? new Date(rule.progressiveRollout.end.date)
          : new Date(),
      },
      percentage: {
        from: rule.progressiveRollout?.initial?.percentage ?? 0,
        to: rule.progressiveRollout?.end?.percentage ?? 100,
      },
    };
  }
  return currentFormVariationSelector;
};
