import { Badge, DarkThemeToggle, Navbar } from "flowbite-react";
import type { FC } from "react";
import { HiMenuAlt1, HiX } from "react-icons/hi";
import { Link } from "react-router-dom";
import { Position } from "../../@types/config";
import { config } from "../../config";
import { isSmallScreen } from "../../helpers/isSmallScreen";
import Image from "../image/Image";
import { useSidebarContext } from "./SidebarContextHelper";

export const AppNavbar: FC<Record<string, never>> = function () {
  const { isCollapsed: isSidebarCollapsed, setCollapsed: setSidebarCollapsed } =
    useSidebarContext();

  const { topBarItems, stage } = config;
  return (
    <header>
      <Navbar
        fluid
        className="fixed top-0 z-30 w-full border-b border-gray-200 bg-white p-0 dark:border-gray-700 dark:bg-gray-800 sm:p-0"
      >
        <div className="w-full p-3 pr-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <button
                aria-controls="sidebar"
                aria-expanded
                className="mr-2 cursor-pointer rounded p-2 text-gray-600 hover:bg-gray-100 hover:text-gray-900 focus:bg-gray-100 focus:ring-2 focus:ring-gray-100 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white dark:focus:bg-gray-700 dark:focus:ring-gray-700"
                onClick={() => setSidebarCollapsed(!isSidebarCollapsed)}
              >
                {isSidebarCollapsed || !isSmallScreen() ? (
                  <HiMenuAlt1 className="h-6 w-6" />
                ) : (
                  <HiX className="h-6 w-6" />
                )}
              </button>
              <Navbar.Brand as={Link} to={"/"}>
                <Image
                  alt="GO Feature Flag logo"
                  height={50}
                  src={config.logo}
                  width={50}
                />
                <span className="self-center whitespace-nowrap px-3 text-xl font-semibold tracking-tight dark:text-white">
                  {config.title}
                </span>
              </Navbar.Brand>
            </div>
            <div className={"flex items-center"}>
              {stage && (
                <Badge color={"indigo"} size={"sm"} className={"mr-2"}>
                  Alpha
                </Badge>
              )}
              {topBarItems
                .filter((i) => !i.disabled)
                .filter((i) => i.position === Position.topRight)
                .map((i) => (
                  <Link
                    to={i.href}
                    key={i.title + i.href}
                    className={
                      "px-1 py-2 text-3xl no-underline hover:text-goff-300 dark:text-white"
                    }
                    target={"_blank"}
                  >
                    <i.icon title={i.title} />
                  </Link>
                ))}
              <DarkThemeToggle />
            </div>
          </div>
        </div>
      </Navbar>
    </header>
  );
};
