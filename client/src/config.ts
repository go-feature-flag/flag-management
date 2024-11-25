import { t } from "i18next";
import { BiBuoy } from "react-icons/bi";
import {
  HiAdjustmentsHorizontal,
  HiBars4,
  HiMiniRocketLaunch,
  HiSquare3Stack3D,
} from "react-icons/hi2";
import { IoLogoGithub, IoMdHelpCircle } from "react-icons/io";
import type { GoffConfiguration } from "./@types/config.ts";
import { Position } from "./@types/config.ts";

export const config: GoffConfiguration = {
  title: "GO Feature Flag",
  stage: "ALPHA",
  logo: "/gofeatureflag.svg",
  apiURL: "http://localhost:3001",
  flagPageConfiguration: {
    successToastDisplayTime: 3000,
  },
  topBarItems: [
    {
      title: "Documentation",
      icon: IoMdHelpCircle,
      href: "https://docs.gofeatureflag.org",
      position: Position.topRight,
    },
    {
      title: "GitHub",
      icon: IoLogoGithub,
      href: "https://github.com/thomaspoignant/go-feature-flag",
      position: Position.topRight,
    },
  ],
  sidebarItems: [
    {
      name: "Main",
      sidebarItems: [
        {
          link: "/",
          title: t("component.navigation.featureFlags"),
          icon: HiMiniRocketLaunch,
        },
        {
          link: "#",
          title: t("component.navigation.integrations"),
          icon: HiSquare3Stack3D,
        },
        {
          link: "/settings",
          title: t("component.navigation.settings"),
          icon: HiAdjustmentsHorizontal,
        },
      ],
    },
    {
      name: "help",
      sidebarItems: [
        {
          link: "https://docs.gofeatureflag.org",
          title: t("component.navigation.documentation"),
          icon: HiBars4,
        },
        {
          link: "/contact",
          title: t("component.navigation.help"),
          icon: BiBuoy,
        },
      ],
    },
  ],
  ruleOperators: [
    { name: "eq", translationKey: "page.flag.ruleOperators.eq" },
    { name: "ne", translationKey: "page.flag.ruleOperators.ne" },
    { name: "lt", translationKey: "page.flag.ruleOperators.lt" },
    { name: "gt", translationKey: "page.flag.ruleOperators.gt" },
    { name: "le", translationKey: "page.flag.ruleOperators.le" },
    { name: "ge", translationKey: "page.flag.ruleOperators.ge" },
    { name: "co", translationKey: "page.flag.ruleOperators.co" },
    { name: "sw", translationKey: "page.flag.ruleOperators.sw" },
    { name: "ew", translationKey: "page.flag.ruleOperators.ew" },
    { name: "in", translationKey: "page.flag.ruleOperators.in" },
    { name: "pr", translationKey: "page.flag.ruleOperators.pr" },
    { name: "not", translationKey: "page.flag.ruleOperators.not" },
  ],
};
