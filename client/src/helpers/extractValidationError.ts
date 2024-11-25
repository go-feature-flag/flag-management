import _ from "lodash";
import type { FieldErrors, FieldValues } from "react-hook-form";

export const extractValidationError = (
  errors: FieldErrors,
  path: string,
): FieldValues => {
  return (_.get(errors, path) as FieldValues) || {};
};

export const hasError = (errors: FieldErrors, path: string): boolean => {
  return _.has(errors, path);
};

export const hasErrors = (errors: FieldErrors, paths: string[]): boolean => {
  for (const item of paths) {
    if (hasError(errors, item)) {
      return true;
    }
  }
  return false;
};

export const extractValidationErrors = (
  errors: FieldErrors,
  paths: string[],
): string[] => {
  const errorMessages = new Set<string>();
  paths
    .map((path) => extractValidationError(errors, path))
    .forEach((error) => {
      if (error.message) {
        errorMessages.add(error.message);
      }
    });
  return [...errorMessages];
};
