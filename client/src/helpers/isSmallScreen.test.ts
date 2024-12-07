import type { Mock } from "vitest";
import { describe, expect, it } from "vitest";
import { isBrowser } from "./isBrowser";
import { isSmallScreen } from "./isSmallScreen";

vi.mock("./isBrowser", () => ({
  isBrowser: vi.fn(),
}));

describe("isSmallScreen", () => {
  it("should return true if the screen width is less than 768", () => {
    (isBrowser as Mock).mockReturnValue(true);
    Object.defineProperty(global, "window", {
      value: {},
      configurable: true,
    });
    Object.defineProperty(window, "innerWidth", {
      value: 767,
      writable: true,
    });

    expect(isSmallScreen()).toBe(true);
  });

  it("should return false if the screen width is 768 or more", () => {
    (isBrowser as Mock).mockReturnValue(true);
    Object.defineProperty(global, "window", {
      value: {},
      configurable: true,
    });
    Object.defineProperty(window, "innerWidth", {
      value: 768,
      writable: true,
    });

    expect(isSmallScreen()).toBe(false);
  });

  it("should return false if not in a browser environment", () => {
    (isBrowser as Mock).mockReturnValue(false);
    expect(isSmallScreen()).toBe(false);
  });
});
