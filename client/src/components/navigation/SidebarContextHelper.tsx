import { useContext } from "react";
import type { SidebarContextProps } from "./SidebarContext.tsx";
import { SidebarContext } from "./SidebarContext.tsx";

export function useSidebarContext(): SidebarContextProps {
  const context = useContext(SidebarContext);

  if (typeof context === "undefined") {
    throw new Error(
      "useSidebarContext should be used within the SidebarContext provider!",
    );
  }

  return context;
}
