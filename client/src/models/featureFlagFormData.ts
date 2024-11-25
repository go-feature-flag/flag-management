import type { typeVariation } from "../@types/typeVariation.ts";

export type SelectedVariationType =
  | "boolean"
  | "string"
  | "integer"
  | "json"
  | "double";
export interface FeatureFlagFormData {
  name?: string;
  id: string;
  disable?: boolean;
  trackEvent?: boolean;
  type: SelectedVariationType;
  variations: FormFlagPageVariationInfo[];
  defaultRule: FormDataVariationSelector;
  description?: string;
  version?: string;
  metadata?: { name: string; value: string; id: string }[];
  targetingRules: FormFlagPageTargetingRule[];
  creationDate?: Date;
  lastUpdatedDate?: Date;
}

export interface FormFlagPageTargetingRule {
  disabled: boolean;
  query: string;
  name: string;
  variation: FormDataVariationSelector;
}

export interface FormFlagPageVariationInfo {
  id?: string;
  name: string;
  value?: typeVariation;
}

export interface FormDataVariationSelector {
  id: string;
  percentage?: Record<string, number>;
  progressive?: IProgressiveRollout;
  variationSelect?: string; // uuid of the variation
}

export interface IProgressiveRollout {
  variation: {
    from: string;
    to: string;
  };
  range: {
    start: Date;
    end: Date;
  };
  percentage: {
    from: number;
    to: number;
  };
}
