import { Breadcrumb, Tooltip } from "flowbite-react";
import { useState } from "react";
import { FaCheck } from "react-icons/fa6";
import { FiCopy } from "react-icons/fi";
import { GoCopy } from "react-icons/go";
import type { IconType } from "react-icons/lib";
import type { To } from "react-router-dom";
import { Link } from "react-router-dom";
import styles from "../../routes/authenticatedRoutes/flag/style.module.css";

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
  const clipBoardTooltipText = tooltipText;
  const [clipBoardClicked, setClipBoardClicked] = useState(false);
  const [clipBoardTooltip, setClipBoardTooltip] =
    useState(clipBoardTooltipText);

  // copyFlagNameToClipboard is a function that copies the name of the feature flag to the clipboard
  // and sets the state of clipBoardClicked to true, which will show a checkmark icon in the UI.
  function copyFlagNameToClipboard() {
    navigator.clipboard.writeText(items.map((item) => item.name).join("/"));
    setClipBoardClicked(true);
    setClipBoardTooltip("Copied!");
    setTimeout(() => {
      setClipBoardClicked(false);
      setClipBoardTooltip(clipBoardTooltipText);
    }, 2000);
  }

  return (
    <Breadcrumb aria-label="Feature Flags" className={"pl-3 pt-3 font-thin"}>
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

      <div
        onClick={copyFlagNameToClipboard}
        className={"ml-3 h-fit cursor-pointer text-goff-600 dark:text-goff-400"}
      >
        {tooltipText && (
          <Tooltip content={clipBoardTooltip}>
            {clipBoardClicked ? <FaCheck /> : <GoCopy />}
          </Tooltip>
        )}
        {!tooltipText && <>{clipBoardClicked ? <FaCheck /> : <FiCopy />}</>}
      </div>
    </Breadcrumb>
  );
};
