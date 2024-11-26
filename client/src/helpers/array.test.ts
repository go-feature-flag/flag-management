import { describe, expect, it } from "vitest";
import { flattenDeep } from "./array";

describe("flattenDeep", () => {
  it("should flatten a nested array", () => {
    const input = [1, [2, [3, [4]], 5]];
    const output = [1, 2, 3, 4, 5];
    expect(flattenDeep(input)).toEqual(output);
  });

  it("should return an empty array when given an empty array", () => {
    const input: string[] = [];
    const output: string[] = [];
    expect(flattenDeep(input)).toEqual(output);
  });

  it("should handle arrays with different types", () => {
    const input = [1, ["a", [true, [null]], 5]];
    const output = [1, "a", true, null, 5];
    expect(flattenDeep(input)).toEqual(output);
  });

  it("should return the same array if it is already flat", () => {
    const input = [1, 2, 3, 4, 5];
    const output = [1, 2, 3, 4, 5];
    expect(flattenDeep(input)).toEqual(output);
  });

  it("should handle deeply nested arrays", () => {
    const input = [[[[[1]]]]];
    const output = [1];
    expect(flattenDeep(input)).toEqual(output);
  });
});
