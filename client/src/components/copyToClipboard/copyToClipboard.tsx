import { Button } from "flowbite-react";
import { useState } from "react";
import { FaCheck } from "react-icons/fa6";
import { IoMdCopy } from "react-icons/io";

export const CopyToClipboard = ({ value }: { value: string }) => {
  const [isCopied, setIsCopied] = useState(false);
  return (
    <Button
      color="gray"
      size={"xs"}
      className={"absolute right-2 top-2"}
      onClick={async (event: { preventDefault: () => void }) => {
        event.preventDefault();
        await navigator.clipboard.writeText(value);
        setIsCopied(true);
        setTimeout(() => {
          setIsCopied(false);
        }, 2000);
      }}
    >
      {isCopied ? (
        <FaCheck className={"mr-1 h-4 w-4 border-green-500 text-green-500"} />
      ) : (
        <IoMdCopy className={"mr-1 h-4 w-4"} />
      )}
      {isCopied ? <span className={"text-green-500"}>Copied</span> : "Copy"}
    </Button>
  );
};
