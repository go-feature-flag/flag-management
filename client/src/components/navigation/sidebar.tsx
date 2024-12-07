import { Sidebar } from "flowbite-react";
import type { FC } from "react";
import { HiExternalLink } from "react-icons/hi";
import { Link } from "react-router-dom";
import { twMerge } from "tailwind-merge";
import { config } from "../../config.ts";
import { useSidebarContext } from "./SidebarContextHelper";

export const AppSidebar: FC = function () {
  const { isCollapsed } = useSidebarContext();
  const { sidebarItems } = config;
  return (
    <Sidebar
      aria-label="Sidebar with multi-level dropdown example"
      collapsed={isCollapsed}
      id="sidebar"
      className={twMerge(
        "fixed inset-y-0 left-0 z-20 mt-16 flex h-full shrink-0 flex-col border-r border-gray-200 duration-75 dark:border-gray-700 lg:flex",
        isCollapsed && "hidden w-16",
      )}
    >
      <Sidebar.Items>
        {sidebarItems.map((itemGroup) => (
          <Sidebar.ItemGroup key={itemGroup.name}>
            {itemGroup.sidebarItems.map((item) => (
              <Sidebar.Item
                to={item.link}
                as={Link}
                icon={item.icon}
                key={item.title}
                target={item.link.startsWith("http") ? "_blank" : "_self"}
              >
                <div className={"flex"}>
                  {item.title}{" "}
                  {item.link.startsWith("http") && <HiExternalLink />}
                </div>
              </Sidebar.Item>
            ))}
          </Sidebar.ItemGroup>
        ))}
      </Sidebar.Items>
    </Sidebar>
  );
};
