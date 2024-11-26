import { v4 as uuidv4 } from "uuid";
import {
  createFeatureFlag,
  deleteFeatureFlagById,
  getFeatureFlagById,
  getFeatureFlags,
  updateFeatureFlagById,
  updateFeatureFlagStatusById,
} from "../api/goffApi.ts";
import {
  VARIATION_BOOLEAN_DEFAULT_VALUE,
  VARIATION_DOUBLE_DEFAULT_VALUE,
  VARIATION_INTEGER_DEFAULT_VALUE,
  VARIATION_JSON_DEFAULT_VALUE,
  VARIATION_STRING_DEFAULT_VALUE,
} from "../models/const.ts";
import type {
  FeatureFlagFormData,
  FormFlagPageVariationInfo,
  SelectedVariationType,
} from "../models/featureFlagFormData.ts";
import { dtoToFlagFormConverter } from "./converterDTOToFormData.ts";
import { convertToDTO } from "./converterFormDataToDTO.ts";

/**
 * Get all flags from the GO Feature Flag API
 * @param req
 */
export const getFlags = async (
  req: Request,
): Promise<FeatureFlagFormData[]> => {
  try {
    const flags = await getFeatureFlags(req);
    return flags.map(dtoToFlagFormConverter);
  } catch (e) {
    console.log(e);
    throw e;
  }
};

const availableDefaultVariations: Record<string, FormFlagPageVariationInfo[]> =
  {
    boolean: VARIATION_BOOLEAN_DEFAULT_VALUE,
    string: VARIATION_STRING_DEFAULT_VALUE,
    integer: VARIATION_INTEGER_DEFAULT_VALUE,
    json: VARIATION_JSON_DEFAULT_VALUE,
    double: VARIATION_DOUBLE_DEFAULT_VALUE,
  };

export const getDefaultFormData = (
  req: Request,
): Promise<FeatureFlagFormData> => {
  const url = new URL(req.url);
  const searchParams = new URLSearchParams(url.search);
  const type: string = searchParams.get("type") ?? "boolean";

  const defaultValueVariations = availableDefaultVariations[type];
  return Promise.resolve({
    id: uuidv4(),
    disable: false,
    trackEvent: false,
    description: "",
    version: "0.0.1",
    type: type as SelectedVariationType,
    creationDate: new Date(),
    lastUpdatedDate: new Date(),
    variations: defaultValueVariations,
    metadata: [],
    defaultRule: {
      id: uuidv4(),
    },
    targetingRules: [],
  });
};

/**
 * Get a flag by its ID from the GO Feature Flag API
 * @param req
 */
export const getFlagById = async (
  req: Request,
): Promise<FeatureFlagFormData> => {
  try {
    const flag = await getFeatureFlagById(req);
    console.log(flag);
    return dtoToFlagFormConverter(flag);
  } catch (e) {
    console.log(e);
    throw e;
  }
};

/**
 * Delete a flag by its ID from the GO Feature Flag API
 * @param id
 */
export const deleteFlagById = async (id: string): Promise<void> => {
  return deleteFeatureFlagById(id);
};

export const changeFlagStatusById = async (
  id: string,
  disable: boolean,
): Promise<FeatureFlagFormData> => {
  const requestBody = {
    disable: disable,
  };
  const flag = await updateFeatureFlagStatusById(id, requestBody);
  return dtoToFlagFormConverter(flag);
};

export const updateFlagById = async (flag: FeatureFlagFormData) => {
  const flagDTO = convertToDTO(flag);
  console.log(flagDTO);
  return updateFeatureFlagById(flag.id, flagDTO);
};

export const createFlag = async (flag: FeatureFlagFormData) => {
  const flagDTO = convertToDTO(flag);
  return createFeatureFlag(flagDTO);
};
