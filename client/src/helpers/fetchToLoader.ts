import type { DeferredData } from "@remix-run/router/utils.ts";
import { defer } from "react-router-dom";

export function fetchToLoader<T>(
  fetcher: (request: Request) => Promise<T>,
  request: Request,
): DeferredData {
  const fetcherData: Promise<T> = fetcher(request);
  // @ts-expect-error: the type error because TS is lost with imports
  return defer({ result: fetcherData });
}
