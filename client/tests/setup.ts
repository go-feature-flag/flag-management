import * as matchers from "@testing-library/jest-dom/matchers";
import "@testing-library/jest-dom/vitest";
import type { ReactNode } from "react";
import { beforeEach, expect, vi } from "vitest";

expect.extend(matchers);

beforeEach(() => {
  vi.mock("react-i18next", () => ({
    // this mock makes sure any components using the translate hook can use it without a warning being shown
    useTranslation: () => {
      return {
        t: (str: string) => str,
        i18n: {
          changeLanguage: () =>
            new Promise(() => {
              // do nothing
            }),
        },
      };
    },
    initReactI18next: {
      type: "3rdParty",
      init: () => {
        // do nothing
      },
    },
    Trans: ({ children }: { children: ReactNode }) => children,
  }));
});
