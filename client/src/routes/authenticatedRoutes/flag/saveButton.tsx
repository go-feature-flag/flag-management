import { clsx } from "clsx";
import { Button, Toast } from "flowbite-react";
import { HiX } from "react-icons/hi";
import { HiCheck } from "react-icons/hi2";
import { MdOutlineSaveAlt } from "react-icons/md";
import Loader from "../../../components/loader/Loader.tsx";

export const FlagSaveButton = ({
  loadingSave,
  displayToastError,
  setDisplayToastError,
  displayToastSuccess,
  setDisplayToastSuccess,
  isDirty,
}: {
  isDirty: boolean;
  loadingSave: boolean;
  displayToastError: boolean;
  displayToastSuccess: boolean;
  setDisplayToastError: (value: boolean) => void;
  setDisplayToastSuccess: (value: boolean) => void;
}) => {
  return (
    <div className={"relative"}>
      <Button
        type="submit"
        size={"xs"}
        className={"absolute right-0"}
        disabled={!isDirty}
      >
        {!loadingSave && <MdOutlineSaveAlt className={"mr-2 h-4 w-4"} />}
        {loadingSave && <Loader className={"mr-2 h-4 w-4"} />}
        Save
      </Button>
      {displayToastError && (
        <SaveToast
          type="error"
          message={"Error in your flag configuration."}
          handleClose={() => setDisplayToastError(false)}
        />
      )}
      {displayToastSuccess && (
        <SaveToast
          type="success"
          message={"Feature Flag successfully saved."}
          handleClose={() => setDisplayToastSuccess(false)}
        />
      )}
    </div>
  );
};

const SaveToast = ({
  type,
  message,
  handleClose,
}: {
  type: "success" | "error";
  message: string;
  handleClose?: () => void;
}) => {
  const colorClass =
    type === "error"
      ? "bg-red-100 text-red-500 dark:bg-red-800 dark:text-red-200"
      : "bg-green-100 text-green-500 dark:bg-green-800 dark:text-green-200";
  return (
    <Toast className={"absolute right-0 mt-10"}>
      <div
        className={clsx(
          colorClass,
          "inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-lg",
        )}
      >
        {type === "error" && <HiX className="h-5 w-5" />}
        {type === "success" && <HiCheck className="h-5 w-5" />}
      </div>
      <div className="ml-3 text-sm font-normal">{message}</div>
      <Toast.Toggle onClick={handleClose} />
    </Toast>
  );
};
