import type { typeVariation } from "../@types/typeVariation.ts";
import { ConvertToDTOError } from "../error.ts";
import {
  SELECT_ROLLOUT_PERCENTAGE,
  SELECT_ROLLOUT_PROGRESSIVE,
} from "../models/const.ts";
import type { FeatureFlagDTO, Rule } from "../models/featureFlagDTO.ts";
import type {
  FeatureFlagFormData,
  FormDataVariationSelector,
  FormFlagPageTargetingRule,
  FormFlagPageVariationInfo,
} from "../models/featureFlagFormData.ts";

export const convertToDTO = (flag: FeatureFlagFormData): FeatureFlagDTO => {
  if (!flag.name) {
    throw new ConvertToDTOError("no name for the feature flag");
  }

  if (!flag.id) {
    throw new ConvertToDTOError("no ID for the feature flag");
  }
  const currentVariations = flag.variations;

  return {
    name: flag.name,
    id: flag.id,
    disable: flag.disable,
    trackEvents: flag.trackEvent,
    description: flag.description,
    version: flag.version,
    type: flag.type,
    createdDate: flag.creationDate?.toISOString() ?? new Date().toISOString(),
    lastUpdatedDate: flag.lastUpdatedDate?.toISOString(),
    variations: Object.fromEntries(
      currentVariations
        .filter(({ value }) => value !== undefined)
        .map(({ name, value }) => [name, value]),
    ) as Record<string, typeVariation>,
    metadata: (flag.metadata ?? []).reduce(
      (acc, { name, value }) => {
        acc[name] = value;
        return acc;
      },
      {} as Record<string, string | number | boolean>,
    ),
    defaultRule: {
      name: "defaultRule",
      query: undefined,
      disable: false,
      ...convertRuleVariationToDto(flag.defaultRule, currentVariations),
    },
    targeting: flag.targetingRules.map((i) =>
      convertRuleToDto(i, currentVariations),
    ),
  };
};

const convertRuleToDto = (
  rule: FormFlagPageTargetingRule,
  variations: FormFlagPageVariationInfo[],
): Rule => {
  return {
    name: rule.name,
    query: rule.query,
    disable: rule.disabled ?? false,
    ...convertRuleVariationToDto(rule.variation, variations),
  };
};

const convertRuleVariationToDto = (
  rule: FormDataVariationSelector,
  variations: FormFlagPageVariationInfo[],
): Rule => {
  if (rule.variationSelect === SELECT_ROLLOUT_PERCENTAGE && rule.percentage) {
    // We put in the percentage only the variations that are in the variations list displayed in the UI
    const acceptedVariations = new Set(variations.map((v) => v.name));
    return {
      id: rule.id,
      percentage: Object.fromEntries(
        Object.entries(rule.percentage).filter(([key]) =>
          acceptedVariations.has(key),
        ),
      ),
    };
  }

  if (rule.variationSelect === SELECT_ROLLOUT_PROGRESSIVE && rule.progressive) {
    return {
      id: rule.id,
      progressiveRollout: {
        initial: {
          variation:
            variations.find((v) => v.id === rule.progressive?.variation.from)
              ?.name ?? "",
          percentage: rule.progressive.percentage.from,
          date: rule.progressive.range.start,
        },
        end: {
          variation:
            variations.find((v) => v.id === rule.progressive?.variation.to)
              ?.name ?? "",
          percentage: rule.progressive.percentage.to,
          date: rule.progressive.range.end,
        },
      },
    };
  }

  return {
    id: rule.id,
    variation: variations.find((v) => v.id === rule.variationSelect)?.name,
  };
};
