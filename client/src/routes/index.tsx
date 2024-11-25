import type { RouteObject } from "react-router-dom";
import { createBrowserRouter } from "react-router-dom";
import { RootLayout } from "../components/layout/RootLayout.tsx";
import { authenticatedRoutes } from "./authenticatedRoutes";
import { ErrorHandler } from "./authenticatedRoutes/error/page.tsx";

// Apply the correct layout to the routes
function getRoutes(): RouteObject[] {
  return authenticatedRoutes.map((route: RouteObject) => {
    if (route.errorElement === undefined) {
      route.errorElement = <ErrorHandler />;
    }
    route.element = <RootLayout>{route.element}</RootLayout>;
    route.errorElement = <RootLayout>{route.errorElement}</RootLayout>;
    return route;
  });
}

export const router = createBrowserRouter(getRoutes());
