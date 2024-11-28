import { Button, Tooltip } from "flowbite-react";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { FaCheck } from "react-icons/fa6";
import { GoCopy } from "react-icons/go";

const baseTranslationKey = "component.copyToClipboard";
export const CopyToClipboard = ({
  textToCopy,
  timeDisplayCopied = 2000,
  size = "xs",
  className,
  tooltipText,
}: {
  textToCopy: string;
  timeDisplayCopied?: number;
  size?: "xs" | "sm";
  className?: string;
  tooltipText?: string;
}) => {
  const [isCopied, setIsCopied] = useState(false);
  return (
    <TooltipWrapper tooltipText={tooltipText} isCopied={isCopied}>
      <CopyToClipboardButton
        textToCopy={textToCopy}
        timeDisplayCopied={timeDisplayCopied}
        size={size}
        className={className}
        isCopied={isCopied}
        setIsCopied={setIsCopied}
      />
    </TooltipWrapper>
  );
};

const TooltipWrapper = ({
  tooltipText,
  children,
  isCopied,
}: {
  tooltipText?: string;
  children: React.ReactNode;
  isCopied: boolean;
}) => {
  const { t } = useTranslation();
  const dynamicTooltipText = () => {
    return isCopied ? (
      <div className={"flex text-goff-400"}>
        <FaCheck className={"mr-1 h-4 w-4 "} />
        {t(`${baseTranslationKey}.copied`)}
      </div>
    ) : (
      tooltipText
    );
  };
  return tooltipText ? (
    <Tooltip
      content={dynamicTooltipText()}
      placement={"right"}
      data-testid={"tooltip-copy-clipboard"}
    >
      {children}
    </Tooltip>
  ) : (
    <>{children}</>
  );
};

const CopyToClipboardButton = ({
  textToCopy,
  timeDisplayCopied = 2000,
  size = "xs",
  className,
  isCopied,
  setIsCopied,
}: {
  textToCopy: string;
  timeDisplayCopied?: number;
  size?: "xs" | "sm";
  className?: string;
  isCopied: boolean;
  setIsCopied: (isCopied: boolean) => void;
}) => {
  return (
    <Button
      data-testid={"copy-to-clipboard-button"}
      color="gray"
      size={"xs"}
      className={className ?? "text-goff-600 dark:text-goff-400"}
      onClick={async (event: { preventDefault: () => void }) => {
        event.preventDefault();
        await navigator.clipboard.writeText(textToCopy);
        setIsCopied(true);
        setTimeout(() => {
          setIsCopied(false);
        }, timeDisplayCopied);
      }}
    >
      <CopyToClipboardIcon isCopied={isCopied} />
      {size !== "xs" && <CopyToClipboardText isCopied={isCopied} />}
    </Button>
  );
};

const CopyToClipboardIcon = ({ isCopied }: { isCopied: boolean }) => {
  return isCopied ? (
    <FaCheck
      className={"mr-1 h-4 w-4 border-goff-400 text-goff-400"}
      data-testid={"check-icon"}
    />
  ) : (
    <GoCopy className={"mr-1 h-4 w-4 "} data-testid={"copy-icon"} />
  );
};
const CopyToClipboardText = ({ isCopied }: { isCopied: boolean }) => {
  const { t } = useTranslation();
  return isCopied ? (
    <span className={"text-goff-400"}>{t(`${baseTranslationKey}.copied`)}</span>
  ) : (
    t(`${baseTranslationKey}.copy`)
  );
};
