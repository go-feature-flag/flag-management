import { isBrowser } from "./isBrowser.ts";

export function isSmallScreen(): boolean {
  return isBrowser() && window.innerWidth < 768;
}
