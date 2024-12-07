import { config } from "../config.ts";
import type {
  FeatureFlagDTO,
  FeatureFlagStatusUpdateDTO,
  NewFeatureFlagDTO,
} from "../models/featureFlagDTO.ts";

const BASE_URL = config.apiURL;

export function getFeatureFlags(req: Request): Promise<FeatureFlagDTO[]> {
  const url = new URL(req.url);
  const filter = url.searchParams.get("filter");
  return fetch(`${BASE_URL}/v1/flags${filter ? `?filter=${filter}` : ""}`, {
    cache: "no-cache",
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error("Error in getFeatureFlags:", error);
      throw error;
    });
}

export function createFeatureFlag(
  featureFlagInput: NewFeatureFlagDTO,
): Promise<FeatureFlagDTO> {
  return fetch(`${BASE_URL}/v1/flags`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(featureFlagInput),
    cache: "no-cache",
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error("Error in createFeatureFlag:", error);
      throw error;
    });
}

export function getFeatureFlagById(req: Request): Promise<FeatureFlagDTO> {
  const url = new URL(req.url);
  const pathSegments = url.pathname
    ?.split("/")
    .filter((segment) => segment !== "");
  const id = pathSegments?.[1];
  return fetch(`${BASE_URL}/v1/flags/${id}`, {
    cache: "no-cache",
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error("Error in getFeatureFlagById:", error);
      throw error;
    });
}

export function updateFeatureFlagById(
  id: string,
  featureFlagUpdate: FeatureFlagDTO,
): Promise<FeatureFlagDTO> {
  return fetch(`${BASE_URL}/v1/flags/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(featureFlagUpdate),
    cache: "no-cache",
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error("Error in updateFeatureFlagById:", error);
      throw error;
    });
}

export function deleteFeatureFlagById(id: string): Promise<void> {
  return fetch(`${BASE_URL}/v1/flags/${id}`, {
    method: "DELETE",
    cache: "no-cache",
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return;
    })
    .catch((error) => {
      console.error("Error in deleteFeatureFlagById:", error);
      throw error;
    });
}

export function updateFeatureFlagStatusById(
  id: string,
  featureFlagStatusUpdate: FeatureFlagStatusUpdateDTO,
): Promise<FeatureFlagDTO> {
  return fetch(`${BASE_URL}/v1/flags/${id}/status`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(featureFlagStatusUpdate),
    cache: "no-cache",
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json();
    })
    .catch((error) => {
      throw error;
    });
}
