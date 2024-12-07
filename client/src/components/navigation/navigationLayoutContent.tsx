import type { FC, PropsWithChildren } from "react";
import { twMerge } from "tailwind-merge";
import { useSidebarContext } from "./SidebarContextHelper";
import { AppNavbar } from "./navbar";
import { AppSidebar } from "./sidebar";

export const NavigationLayoutContent: FC<PropsWithChildren> = function ({
  children,
}) {
  const { isCollapsed } = useSidebarContext();

  return (
    <>
      <AppNavbar />
      <div className="mt-16 flex items-start">
        <AppSidebar />
        <div
          id="main-content"
          className={twMerge(
            "relative h-full w-full overflow-y-auto",
            isCollapsed ? "lg:ml-[4.5rem]" : "lg:ml-64",
          )}
        >
          {children}
        </div>
      </div>
    </>
  );
};
