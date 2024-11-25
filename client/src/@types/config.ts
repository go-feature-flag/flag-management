import type { IconType } from "react-icons/lib";

export interface GoffConfiguration {
  // The title of the application will appear in each title bar
  title: string;

  // Stage is the actual stage of the app, if not empty a badge is displayed with the stage
  stage?: string;

  // The logo of the application will be used everywhere a logo is needed
  logo: string;

  // apiURL is the URL of the API, it will be used to make requests
  apiURL: string;

  // The items that will appear in the top bar menu
  topBarItems: NavbarItemProps[];

  // The items that will appear in the sidebar menu
  sidebarItems: SidebarItemGroupProps[];

  // The operators that can be used in the rules
  ruleOperators: { name: string; translationKey: string }[];

  flagPageConfiguration: FlagPageConfigurations;
}

export interface NavbarItemProps {
  title: string;
  icon: IconType;
  position: Position;
  href: string;
  disabled?: boolean;
}

export enum Position {
  topRight,
  sideTop,
  sideBottom,
}

export interface SidebarItemProps {
  link: string;
  title: string;
  icon: IconType;
}

export interface SidebarItemGroupProps {
  name: string;
  sidebarItems: SidebarItemProps[];
}

export interface FlagPageConfigurations {
  successToastDisplayTime: number;
}
