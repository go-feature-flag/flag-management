import { clsx } from "clsx";
import type { ReactNode } from "react";

export interface SectionProps {
  children: ReactNode;
  outsideTitle?: string;
  innerTitle?: string;
  info?: string;
  infoStyle?: "green" | "gray";
  titleSize?: "xs" | "sm" | "base" | "lg" | "xl" | "2xl";
  borderColour?: "base" | "danger";
  maxWidth?: "sm" | "md" | "lg" | "xl" | "2xl" | "3xl" | "4xl";
  className?: string;
}

const borderColourConf = {
  base: "border-gray-300 dark:border-gray-800",
  danger: "border-red-500 dark:border-red-500",
};
export const Section = ({
  children,
  outsideTitle,
  titleSize,
  info,
  infoStyle,
  borderColour = "base",
  innerTitle,
  maxWidth,
  className,
}: SectionProps) => {
  const titleClassSize = titleSize ? `text-${titleSize}` : "";
  const maxWidthClass = maxWidth ? `max-w-${maxWidth}` : "";
  const infoClass =
    infoStyle === "gray"
      ? "text-sm text-gray-500"
      : "text-base text-goff-700 dark:text-goff-400";

  return (
    <div className={`${maxWidthClass} mt-2`}>
      {outsideTitle && (
        <h1 className={`${titleClassSize} items-baseline`}>{outsideTitle}</h1>
      )}
      <div
        className={
          className === undefined || className === ""
            ? `${borderColourConf[borderColour]} rounded-lg border p-4`
            : className
        }
      >
        {innerTitle && <h2 className={"mb-2 text-xl"}>{innerTitle}</h2>}
        {info && <div className={clsx(infoClass, "mb-3")}>{info}</div>}
        {children}
      </div>
    </div>
  );
};
