import { Breadcrumb } from "flowbite-react";
import type { FC, ReactNode } from "react";
import { useTranslation } from "react-i18next";
import { HiMiniRocketLaunch } from "react-icons/hi2";
import { Link } from "react-router-dom";
import { CopyToClipboard } from "../../../components/copyToClipboard/copyToClipboard.tsx";
import styles from "./style.module.css";

const translationBaseKey = "page.flag";

export interface FlagLayoutProps {
  children: ReactNode;
  name: string;
}

export const FlagLayout: FC<FlagLayoutProps> = ({ children, name }) => {
  return (
    <>
      <FlagHierarchy name={name} />
      <div className="me:py-3 grid grid-cols-1 place-items-center px-3 py-2 md:px-6 lg:px-10 lg:py-5">
        <div className="w-full max-w-7xl grid-cols-2 overflow-x-auto">
          {children}
        </div>
      </div>
    </>
  );
};

const FlagHierarchy = ({ name }: { name: string }) => {
  const { t } = useTranslation();
  return (
    <Breadcrumb
      aria-label="Feature Flags"
      className={"flex pl-3 pt-3 font-thin"}
    >
      <Breadcrumb.Item
        className={styles.breadcrumbItem}
        icon={HiMiniRocketLaunch}
      >
        <Link to="/">{t("component.navigation.featureFlags")}</Link>
      </Breadcrumb.Item>
      <Breadcrumb.Item className={"text-base font-light lg:text-lg"}>
        {name}
      </Breadcrumb.Item>
      <CopyToClipboard
        tooltipText={t(`${translationBaseKey}.copyTooltip`)}
        textToCopy={"toto"}
        size={"xs"}
        className={
          "ml-1 border-0 p-0 text-goff-600 hover:bg-goff-50 dark:text-goff-400"
        }
      />
    </Breadcrumb>
  );
};
