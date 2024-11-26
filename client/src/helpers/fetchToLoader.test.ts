import { describe, expect, it, vi } from "vitest";
import { fetchToLoader } from "./fetchToLoader";

describe("fetchToLoader", () => {
  it("should return deferred data", async () => {
    const mockRequest = new Request("https://example.com");
    const mockResponse = { data: "test" };
    const mockFetcher = vi.fn().mockResolvedValue(mockResponse);

    const result = fetchToLoader(mockFetcher, mockRequest);
    expect(mockFetcher).toHaveBeenCalledWith(mockRequest);
    expect(result["deferredKeys"]).toEqual(["result"]);
  });
});
