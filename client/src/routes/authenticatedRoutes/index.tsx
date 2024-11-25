import type { FC } from "react";
import type { RouteObject } from "react-router-dom";
import { RootLayout } from "../../components/layout/RootLayout.tsx";
import type { LayoutProps } from "../../components/layout/layoutProps.ts";
import { fetchToLoader } from "../../helpers/fetchToLoader.ts";
import type { FeatureFlagFormData } from "../../models/featureFlagFormData.ts";
import {
  getDefaultFormData,
  getFlagById,
  getFlags,
} from "../../service/featureFlagService.ts";
import ContactPage from "./contact/page.tsx";
import { FlagPage } from "./flag/page.tsx";
import { FlagListPage } from "./flags/page.tsx";

export const authenticatedLayout: FC<LayoutProps> = RootLayout;
export const authenticatedRoutes: RouteObject[] = [
  // { path: "/", element: <HomePage /> },
  { path: "/contact", element: <ContactPage /> },
  {
    path: "/",
    element: <FlagListPage />,
    loader: ({ request }) =>
      fetchToLoader<FeatureFlagFormData[]>(getFlags, request),
  },

  {
    path: "/flags/:flagId",
    element: <FlagPage />,
    loader: ({ request }) =>
      fetchToLoader<FeatureFlagFormData>(getFlagById, request),
  },
  {
    path: "/flags/new/:flagName",
    element: <FlagPage isNew={true} />,
    loader: ({ request }) => {
      return fetchToLoader<FeatureFlagFormData>(getDefaultFormData, request);
    },
  },
];
