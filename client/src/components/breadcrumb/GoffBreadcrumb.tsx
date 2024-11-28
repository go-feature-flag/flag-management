import { Breadcrumb } from "flowbite-react";
import type { IconType } from "react-icons/lib";
import type { To } from "react-router-dom";
import { Link } from "react-router-dom";
import styles from "../../routes/authenticatedRoutes/flag/style.module.css";
import { CopyToClipboard } from "../copyToClipboard/copyToClipboard.tsx";

export interface GoffBreadcrumbProps {
  tooltipText?: string;
  items: { name: string; to?: To }[];
  title: { icon: IconType; text: string; to?: To };
}

export const GoffBreadcrumb = ({
  tooltipText,
  items,
  title,
}: GoffBreadcrumbProps) => {
  return (
    <Breadcrumb
      aria-label="Feature Flags"
      className={"flex pl-3 pt-3 font-thin"}
    >
      <Breadcrumb.Item className={styles.breadcrumbItem} icon={title.icon}>
        {title.to !== undefined && <Link to={title.to}>{title.text}</Link>}
        {title.to === undefined && title.text}
      </Breadcrumb.Item>
      {items.map((item, index) => (
        <Breadcrumb.Item
          className={styles.breadcrumbItem}
          key={item.name + index}
        >
          {item.name}
        </Breadcrumb.Item>
      ))}

      <CopyToClipboard
        tooltipText={tooltipText}
        textToCopy={"toto"}
        size={"xs"}
        className={
          "ml-1 border-0 p-0 text-goff-600 hover:bg-goff-50 dark:text-goff-400"
        }
      />
    </Breadcrumb>
  );
};
