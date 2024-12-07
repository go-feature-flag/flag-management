import type { RouteObject } from "react-router-dom";
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
