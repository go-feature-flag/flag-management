import { Badge, Button, Drawer, Textarea } from "flowbite-react";
import { useState } from "react";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { IoMdInformationCircle } from "react-icons/io";
import { RxCross1 } from "react-icons/rx";
import {
  extractValidationError,
  hasError,
} from "../../../../helpers/extractValidationError.ts";
import { QueryFormatSidebar } from "./queryFormatSidebar.tsx";

const translationBaseKey = "page.flag.targeting.advancedRule";
export const RuleAdvancedQuery = ({ fieldName }: { fieldName?: string }) => {
  const [showInfo, setShowInfo] = useState(false);
  const { t } = useTranslation();
  const {
    register,
    watch,
    formState: { errors },
  } = useFormContext();
  return (
    <>
      <div className={"flex items-start"}>
        <Badge size="sm" className={"mr-3 mt-1"} color={"green"}>
          IF
        </Badge>
        <Textarea
          className={"font-mono"}
          {...register(`${fieldName}.query`, {
            required: t(`${translationBaseKey}.errors.queryRequired`, {
              name: watch(`${fieldName}.name`),
            }),
          })}
          color={hasError(errors, `${fieldName}.query`) ? "failure" : "gray"}
        />
        <button
          type="button"
          className={"ml-3 text-gray-800 dark:text-gray-50"}
          onClick={() => setShowInfo(true)}
        >
          <IoMdInformationCircle className="h-5 w-5" />
        </button>
      </div>
      {hasError(errors, `${fieldName}.query`) && (
        <span className={"mt-2 flex text-xs text-red-600 dark:text-red-400"}>
          {extractValidationError(errors, `${fieldName}.query`)?.message}
        </span>
      )}
      <Drawer
        open={showInfo}
        onClose={() => setShowInfo(false)}
        backdrop={false}
        position="right"
        edge={true}
      >
        <div>
          <h5
            className="mb-4 inline-flex items-center text-base font-semibold text-gray-500 dark:text-gray-400"
            id=":r4r:"
          >
            <IoMdInformationCircle className={"mr-3"} />
            {t(`${translationBaseKey}.queryFormat.title`)}
          </h5>
          <Button
            onClick={() => setShowInfo(false)}
            className="absolute end-2.5 top-4 flex h-6 w-6 items-center justify-center rounded-lg"
          >
            <RxCross1 />
            <span className="sr-only">Close menu</span>
          </Button>
          <span className="hidden" id="flowbite-drawer-header-:r5f:"></span>
        </div>
        <Drawer.Items className={"mb-6 text-gray-500 dark:text-gray-400"}>
          <QueryFormatSidebar />
        </Drawer.Items>
      </Drawer>
    </>
  );
};
