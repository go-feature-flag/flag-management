import { describe, expect, it, vi } from "vitest";
import { isBrowser } from "./isBrowser";
import { isSmallScreen } from "./isSmallScreen";

vi.mock("./isBrowser", () => ({
  isBrowser: vi.fn(),
}));

describe("isSmallScreen", () => {
  it("should return true if the screen width is less than 768", () => {
    (isBrowser as vi.Mock).mockReturnValue(true);
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
    (isBrowser as vi.Mock).mockReturnValue(true);
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
    (isBrowser as vi.Mock).mockReturnValue(false);
    expect(isSmallScreen()).toBe(false);
  });
});
