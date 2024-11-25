import type { FC, PropsWithChildren } from "react";
import { createContext, useEffect, useState } from "react";
import { isBrowser } from "../../helpers/isBrowser.ts";
import { isSmallScreen } from "../../helpers/isSmallScreen.ts";

export interface SidebarContextProps {
  isCollapsed: boolean;
  setCollapsed: (isOpen: boolean) => void;
}

export const SidebarContext = createContext<SidebarContextProps>(
  {} as SidebarContextProps,
);

export const SidebarProvider: FC<PropsWithChildren> = function ({ children }) {
  const location = isBrowser() ? window.location.pathname : "/";
  const storedIsCollapsed = isBrowser()
    ? localStorage.getItem("isSidebarCollapsed") === "true"
    : false;

  const [isCollapsed, setCollapsed] = useState(storedIsCollapsed);

  // Close Sidebar on page change on mobile
  useEffect(() => {
    if (isSmallScreen()) {
      setCollapsed(true);
    }
  }, [location]);

  // Close Sidebar on mobile tap inside main content
  useEffect(() => {
    function handleMobileTapInsideMain(event: MouseEvent) {
      const main = document.querySelector("#main-content");
      const isClickInsideMain = main?.contains(event.target as Node);

      if (isSmallScreen() && isClickInsideMain) {
        setCollapsed(true);
      }
    }

    document.addEventListener("mousedown", handleMobileTapInsideMain);

    return () => {
      document.removeEventListener("mousedown", handleMobileTapInsideMain);
    };
  }, []);

  // Update local storage when collapsed state changed
  useEffect(() => {
    localStorage.setItem("isSidebarCollapsed", isCollapsed ? "true" : "false");
  }, [isCollapsed]);

  return (
    <SidebarContext.Provider
      value={{
        isCollapsed,
        setCollapsed,
      }}
    >
      {children}
    </SidebarContext.Provider>
  );
};
