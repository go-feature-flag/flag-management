import { describe, expect, it } from "vitest";
import { formatAndLocalizedDate } from "./dateFormater";

describe("formatDate", () => {
  it("should return an empty string if no date is provided", () => {
    expect(formatAndLocalizedDate()).toBe("");
  });

  it("should format a valid date correctly", () => {
    const date = new Date(Date.UTC(2024, 10, 25, 14, 48, 0));
    expect(formatAndLocalizedDate(date, "en-US", { timeZone: "UTC" })).toBe(
      "Nov 25, 2024, 02:48 PM",
    );
  });

  it("should handle different date inputs", () => {
    const date1 = new Date("2022-01-01T00:00:00.000Z");
    const date2 = new Date("2022-12-31T23:59:59.999Z");
    expect(formatAndLocalizedDate(date1, "en-US", { timeZone: "UTC" })).toBe(
      "Jan 1, 2022, 12:00 AM",
    );
    expect(formatAndLocalizedDate(date2, "en-US", { timeZone: "UTC" })).toBe(
      "Dec 31, 2022, 11:59 PM",
    );
  });
});
