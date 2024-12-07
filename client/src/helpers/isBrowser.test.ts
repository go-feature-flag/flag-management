import { describe, expect, it } from "vitest";
import { isBrowser } from "./isBrowser";

describe("isBrowser", () => {
  it("should return true if running in a browser environment", () => {
    const originalWindow = global.window;
    Object.defineProperty(global, "window", { value: {}, configurable: true });
    expect(isBrowser()).toBe(true);
    Object.defineProperty(global, "window", {
      value: originalWindow,
      configurable: true,
    });
  });

  it("should return false if not running in a browser environment", () => {
    const originalWindow = global.window;
    Object.defineProperty(global, "window", {
      value: undefined,
      configurable: true,
    });
    expect(isBrowser()).toBe(false);
    Object.defineProperty(global, "window", {
      value: originalWindow,
      configurable: true,
    });
  });
});
