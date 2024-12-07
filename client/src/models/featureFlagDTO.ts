import type { typeVariation } from "../@types/typeVariation.ts";
import type { SelectedVariationType } from "./featureFlagFormData.ts";

interface commonApiDto {
  id: string;
  name: string;
  createdDate: string;
  lastUpdatedDate?: string;
  description?: string;
  trackEvents?: boolean;
  disable?: boolean;
  version?: string;
  variations: Record<string, typeVariation>;
  targeting?: Rule[];
  defaultRule: Rule;
  experimentation?: {
    start: Date;
    end: Date;
  };
  metadata: Record<string, string | number | boolean>;
  type: SelectedVariationType;
  // NOTE: "scheduled" is on purpose not present in the DTO because it is not an implemented feature.
}

export interface Rule {
  name?: string;
  query?: string;
  variation?: string;
  percentage?: Record<string, number>;
  progressiveRollout?: {
    initial?: ProgressiveRolloutBand;
    end?: ProgressiveRolloutBand;
  };
  disable?: boolean;
  id: string;
}

export interface ProgressiveRolloutBand {
  variation: string;
  percentage: number;
  date: Date;
}

export interface FeatureFlagDTO extends commonApiDto {
  id: string;
}

export type NewFeatureFlagDTO = commonApiDto;

export interface FeatureFlagStatusUpdateDTO {
  disable: boolean;
}
