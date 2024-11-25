import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeadCell,
  TableRow,
  TextInput,
} from "flowbite-react";
import { Suspense, useState } from "react";
import { useTranslation } from "react-i18next";
import { FaCirclePlus } from "react-icons/fa6";
import { HiMagnifyingGlass } from "react-icons/hi2";
import { Await, useAsyncError, useLoaderData } from "react-router-dom";
import { AlertError } from "../../../components/alert/error.tsx";
import Loader from "../../../components/loader/Loader.tsx";
import type { FeatureFlagFormData } from "../../../models/featureFlagFormData.ts";
import { FlagRow } from "./flagRow";
import { NewFlagModal } from "./newFlagModal.tsx";

const translationBaseKey = "page.flags.flagList";
export const FlagList = () => {
  const data = useLoaderData() as { result: FeatureFlagFormData[] };
  return (
    <Suspense fallback={<Loader />}>
      <Await resolve={data.result} errorElement={<ErrorLoadingPage />}>
        {(featureFlags) => <FlagTable featureFlags={featureFlags} />}
      </Await>
    </Suspense>
  );
};

/**
 * ErrorLoadingPage is the error state of the flag page.
 */
const ErrorLoadingPage = () => {
  const { t } = useTranslation();
  const error = useAsyncError() as Error;
  return (
    <AlertError
      text={t(`${translationBaseKey}.errors.loading`)}
      proposeRefresh={true}
      errorDetails={error.message}
    />
  );
};

/**
 * FlagTable contains the table of flags
 */
const FlagTable = ({
  featureFlags,
}: {
  featureFlags: FeatureFlagFormData[];
}) => {
  const [flagFilter, setFlagFilter] = useState("");
  const [flagList, setFlagList] = useState(featureFlags);
  const [newFlag, setNewFlag] = useState(false);
  const { t } = useTranslation();

  return (
    <>
      {newFlag && (
        <NewFlagModal
          setNewFlag={setNewFlag}
          newFlag={newFlag}
          featureFlags={featureFlags}
        />
      )}
      <Table hoverable>
        <TableHead>
          <TableHeadCell colSpan={3}>
            <div className="max-w-xs">
              <TextInput
                type="text"
                rightIcon={HiMagnifyingGlass}
                placeholder={t(
                  `${translationBaseKey}.topBar.searchPlaceholder`,
                )}
                value={flagFilter}
                onChange={(event) => setFlagFilter(event.target.value)}
              />
            </div>
          </TableHeadCell>
          <TableHeadCell>
            <div className="flex justify-end">
              <Button
                size={"xl"}
                onClick={() => setNewFlag(true)}
                className={"w-fit"}
              >
                <FaCirclePlus className={"h-5 w-5"} />
                <span className="pl-1.5 text-xs md:text-base">
                  {t(`${translationBaseKey}.topBar.createFlagButton`)}
                </span>
              </Button>
            </div>
          </TableHeadCell>
        </TableHead>
        <TableBody className="divide-y">
          {flagList
            .filter((flag) => flag.name?.includes(flagFilter))
            .map((flag) => {
              return (
                <FlagRow
                  flags={flagList}
                  setFlags={setFlagList}
                  currentFlag={flag}
                  key={flag.id}
                />
              );
            })}
          <TableRow className="bg-white dark:border-gray-700 dark:bg-gray-800">
            {flagList.filter((flag) => flag.name?.includes(flagFilter))
              .length === 0 && (
              <TableCell colSpan={4}>
                {t(`${translationBaseKey}.noResult`)}
              </TableCell>
            )}
          </TableRow>
        </TableBody>
      </Table>
    </>
  );
};
