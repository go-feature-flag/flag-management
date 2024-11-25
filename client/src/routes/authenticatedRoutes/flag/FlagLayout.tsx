import type { FC } from "react";
import { useTranslation } from "react-i18next";
import { HiMiniRocketLaunch } from "react-icons/hi2";
import { GoffBreadcrumb } from "../../../components/breadcrumb/GoffBreadcrumb.tsx";
import type { LayoutProps } from "../../../components/layout/layoutProps.ts";

const translationBaseKey = "page.flag";
export interface FlagLayoutProps extends LayoutProps {
  name: string;
}

export const FlagLayout: FC<FlagLayoutProps> = ({ children, name }) => {
  const { t } = useTranslation();
  return (
    <>
      <GoffBreadcrumb
        title={{
          icon: HiMiniRocketLaunch,
          text: t("component.navigation.featureFlags"),
          to: "/",
        }}
        items={[{ name: name }]}
        tooltipText={t(`${translationBaseKey}.copyTooltip`)}
      />
      <div className="me:py-3 grid grid-cols-1 place-items-center px-3 py-2 md:px-6 lg:px-10 lg:py-5">
        <div className="w-full max-w-7xl grid-cols-2 overflow-x-auto">
          {children}
        </div>
      </div>
    </>
  );
};
