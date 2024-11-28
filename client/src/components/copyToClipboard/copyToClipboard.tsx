import { Button } from "flowbite-react";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { FaCheck } from "react-icons/fa6";
import { IoMdCopy } from "react-icons/io";

const baseTranslationKey = "component.copyToClipboard";
export const CopyToClipboard = ({
  textToCopy,
  timeDisplayCopied = 2000,
}: {
  textToCopy: string;
  timeDisplayCopied?: number;
}) => {
  const { t } = useTranslation();
  const [isCopied, setIsCopied] = useState(false);
  return (
    <Button
      color="gray"
      size={"xs"}
      className={"absolute right-2 top-2"}
      onClick={async (event: { preventDefault: () => void }) => {
        event.preventDefault();
        await navigator.clipboard.writeText(textToCopy);
        setIsCopied(true);
        setTimeout(() => {
          setIsCopied(false);
        }, timeDisplayCopied);
      }}
    >
      {isCopied ? (
        <FaCheck
          className={"mr-1 h-4 w-4 border-green-500 text-green-500"}
          data-testid={"check-icon"}
        />
      ) : (
        <IoMdCopy className={"mr-1 h-4 w-4"} data-testid={"copy-icon"} />
      )}
      {isCopied ? (
        <span className={"text-green-500"}>
          {t(`${baseTranslationKey}.copied`)}
        </span>
      ) : (
        t(`${baseTranslationKey}.copy`)
      )}
    </Button>
  );
};
