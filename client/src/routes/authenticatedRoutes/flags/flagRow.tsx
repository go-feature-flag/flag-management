import {
  Button,
  TableCell,
  TableRow,
  ToggleSwitch,
  Tooltip,
} from "flowbite-react";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { FaTrashAlt } from "react-icons/fa";
import { FaToggleOn } from "react-icons/fa6";
import {
  HiInformationCircle,
  HiOutlineInformationCircle,
} from "react-icons/hi2";
import { Link, useNavigate } from "react-router-dom";
import {
  deleteFeatureFlagById,
  updateFeatureFlagStatusById,
} from "../../../api/goffApi.ts";
import { ConfirmationModal } from "../../../components/modal/ConfirmationModal.tsx";
import { formatAndLocalizedDate } from "../../../helpers/dateFormater.ts";
import type { FeatureFlagFormData } from "../../../models/featureFlagFormData.ts";
import styles from "./styles.module.css";

const translationBaseKey = "page.flags.flagList.row";

export function FlagRow(props: {
  currentFlag: FeatureFlagFormData;
  flags: FeatureFlagFormData[];
  setFlags: (flags: FeatureFlagFormData[]) => void;
}) {
  const {
    creationDate,
    disable,
    name,
    id,
    description,
    lastUpdatedDate,
    type,
    version,
    variations,
  } = props.currentFlag;
  const [flagEnable, setFlagEnable] = useState(!disable);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [openConfirmationDisable, setOpenConfirmationDisable] = useState(false);
  const [errorDisable, setErrorDisable] = useState("");
  const [errorDelete, setErrorDelete] = useState("");

  const navigate = useNavigate();
  const flagDetailsPageLocation = `/flags/${id}`;
  const { t } = useTranslation();
  /**
   * handleConfirmDisable is called when the user confirms the disabling/enabling of the feature flag
   */
  const handleConfirmDisable = async () => {
    const checkOldValue = flagEnable;
    try {
      await updateFeatureFlagStatusById(id, { disable: checkOldValue });
      setFlagEnable(!checkOldValue);
      setOpenConfirmationDisable(false);
    } catch (error) {
      setErrorDisable(
        `${t(`${translationBaseKey}.errors.statusChange`)} ${error}`,
      );
      setFlagEnable(checkOldValue);
    }
  };

  const handleConfirmDelete = async () => {
    try {
      await deleteFeatureFlagById(id);
      props.setFlags(props.flags.filter((flag) => flag.id !== id));
      setOpenDeleteModal(false);
    } catch (error) {
      setErrorDelete(`${t(`${translationBaseKey}.errors.delete`)} ${error}`);
    }
  };
  return (
    <TableRow className="bg-white dark:border-gray-700 dark:bg-gray-800">
      <TableCell
        className="cursor-pointer text-base"
        onClick={() => navigate(flagDetailsPageLocation)}
      >
        <div className="flex place-items-center">
          <span className={"text-lg text-gray-800 dark:text-gray-200"}>
            {name}
          </span>
          <div className={styles.infoIcon}>
            <Tooltip
              content={`${t(
                `${translationBaseKey}.tooltip.created`,
              )} ${formatAndLocalizedDate(creationDate)}`}
            >
              <HiOutlineInformationCircle />
            </Tooltip>
          </div>
        </div>
        {description !== "" && (
          <p className="text-sm font-normal text-gray-500 dark:text-gray-500">
            {description?.slice(0, 70)}
            {(description ? description.length : 0) > 70 && "..."}
          </p>
        )}
        <div className={"mt-1 max-md:hidden"}>
          <span className="me-2 rounded bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800 dark:bg-gray-700 dark:text-gray-300">
            {type}
          </span>
          {variations && (
            <span className="me-2 rounded bg-green-100 px-2.5 py-0.5 text-sm font-medium text-green-800 dark:bg-green-900 dark:text-green-300">
              {variations.length} variations
            </span>
          )}
          {version && (
            <span className="me-2 rounded bg-indigo-100 px-2.5 py-0.5 text-xs font-medium text-indigo-800 dark:bg-indigo-900 dark:text-indigo-300">
              {version}
            </span>
          )}
        </div>
      </TableCell>
      <TableCell
        className={styles.lastUpdated}
        onClick={() => navigate(flagDetailsPageLocation)}
      >
        <span className={"max-lg:hidden"}>
          {t(`${translationBaseKey}.lastUpdated`)}{" "}
          {formatAndLocalizedDate(lastUpdatedDate)}
        </span>
      </TableCell>
      <TableCell className={"max-w-fit"}>
        <Tooltip content={t(`${translationBaseKey}.tooltip.status`)}>
          <ToggleSwitch
            checked={flagEnable}
            onChange={() => setOpenConfirmationDisable(true)}
          />
        </Tooltip>
      </TableCell>
      <TableCell className={"max-w-fit"}>
        <div className={"right-1 flex justify-end"}>
          <Link to={flagDetailsPageLocation}>
            <Button size="sm" className={"mr-5"}>
              <Tooltip content={t(`${translationBaseKey}.tooltip.info`)}>
                <HiInformationCircle className="h-4 w-4" />
              </Tooltip>
            </Button>
          </Link>
          <Button size="sm" onClick={() => setOpenDeleteModal(true)}>
            <Tooltip content={t(`${translationBaseKey}.tooltip.delete`)}>
              <FaTrashAlt className="h-4 w-4" />
            </Tooltip>
          </Button>
          <ConfirmationModal
            text={t(`${translationBaseKey}.modal.delete`, { name })}
            isOpen={openDeleteModal}
            onClickYes={handleConfirmDelete}
            onClickCancel={() => {
              setOpenDeleteModal(false);
              setErrorDelete("");
            }}
            error={errorDelete}
            confirmationText={name}
          />
          <ConfirmationModal
            text={t(`${translationBaseKey}.modal.enable`, {
              name: name,
              action: flagEnable ? "disable" : "enable",
            })}
            isOpen={openConfirmationDisable}
            onClickYes={handleConfirmDisable}
            onClickCancel={() => {
              setOpenConfirmationDisable(false);
              setErrorDisable("");
            }}
            error={errorDisable}
            icon={
              <FaToggleOn className="mx-auto mb-4 h-14 w-14 text-gray-400 dark:text-gray-200" />
            }
          />
        </div>
      </TableCell>
    </TableRow>
  );
}
