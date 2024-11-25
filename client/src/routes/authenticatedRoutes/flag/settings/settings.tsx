import { Button, Textarea, TextInput, ToggleSwitch } from "flowbite-react";
import * as log from "loglevel";
import type { ReactNode } from "react";
import { useState } from "react";
import { Controller, useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { FaTrashAlt } from "react-icons/fa";
import { FaToggleOff, FaToggleOn } from "react-icons/fa6";
import { useNavigate } from "react-router-dom";
import { ConfirmationModal } from "../../../../components/modal/ConfirmationModal.tsx";
import { Section } from "../../../../components/section/section.tsx";
import type { FeatureFlagFormData } from "../../../../models/featureFlagFormData.ts";
import {
  changeFlagStatusById,
  deleteFlagById,
} from "../../../../service/featureFlagService.ts";

const translationBaseKey = "page.flag.settings";
export const Settings = () => {
  return (
    <>
      <NormalZone />
      <DangerZone />
    </>
  );
};

const NormalZone = () => {
  const { t } = useTranslation();

  return (
    <Section outsideTitle={t(`${translationBaseKey}.normalZone`)}>
      <div className="grid w-full grid-cols-3 gap-4 ">
        <Description />
        <hr className={"col-span-3 m-0 "} />
        <Version />
        <hr className={"col-span-3 m-0 "} />
        <TrackEvent />
      </div>
    </Section>
  );
};

const Description = () => {
  const { t } = useTranslation();
  const { register } = useFormContext();
  return (
    <SettingText
      title={t(`${translationBaseKey}.description.title`)}
      description={t(`${translationBaseKey}.description.description`)}
      component={
        <Textarea
          {...register("description")}
          placeholder={
            t(`${translationBaseKey}.description.description`) + ".."
          }
          className={"float-right max-w-sm md:max-w-xl"}
          rows={2}
        />
      }
    />
  );
};

const TrackEvent = () => {
  const { t } = useTranslation();
  const { control, setValue } = useFormContext();
  const { trackEvent } = useWatch<FeatureFlagFormData>();
  return (
    <SettingEndOfLine
      title={t(`${translationBaseKey}.trackEvent.title`)}
      description={t(`${translationBaseKey}.trackEvent.description`)}
      component={
        <Controller
          name={"trackEvent"}
          control={control}
          render={({ field }) => (
            <ToggleSwitch
              checked={field.value}
              className={"float-right pr-10"}
              value={field.value}
              onChange={() => {
                setValue("trackEvent", !trackEvent, { shouldDirty: true });
              }}
            />
          )}
        />
      }
    />
  );
};

const Version = () => {
  const { t } = useTranslation();
  const { register } = useFormContext();
  return (
    <SettingText
      title={t(`${translationBaseKey}.version.title`)}
      description={t(`${translationBaseKey}.version.description`)}
      component={
        <TextInput
          {...register("version")}
          type="text"
          placeholder="0.0.1"
          shadow
          className={"float-right max-w-sm md:max-w-xl"}
        />
      }
    />
  );
};

const DangerZone = () => {
  const { disable, name, id } = useWatch<FeatureFlagFormData>();
  const { t } = useTranslation();
  const [flagEnable, setFlagEnable] = useState(!disable);
  const [openDeleteModal, setOpenDeleteModal] = useState(false);
  const [openConfirmationDisable, setOpenConfirmationDisable] = useState(false);
  const [errorDelete, setErrorDelete] = useState("");
  const [errorDisable, setErrorDisable] = useState("");
  const navigate = useNavigate();
  const handleConfirmDelete = async () => {
    try {
      if (id) {
        await deleteFlagById(id);
        setOpenDeleteModal(false);
        navigate("/flags");
      } else {
        log.error("handleConfirmDelete: Flag ID is not defined");
      }
    } catch (error) {
      setErrorDelete(t(`${translationBaseKey}.delete.error`, { error: error }));
    }
  };

  /**
   * handleConfirmDisable is called when the user confirms the disabling/enabling of the feature flag
   */
  const handleConfirmDisable = async () => {
    const checkOldValue = flagEnable;
    try {
      setErrorDisable("");
      if (id) {
        await changeFlagStatusById(id, checkOldValue);
        setFlagEnable(!checkOldValue);
        setOpenConfirmationDisable(false);
      } else {
        log.error("handleConfirmDisable: Flag ID is not defined");
      }
    } catch (error) {
      setErrorDisable(
        t(`${translationBaseKey}.status.error`, { error: error }),
      );
      setFlagEnable(checkOldValue);
    }
  };

  return (
    <>
      <Section
        borderColour={"danger"}
        outsideTitle={t(`${translationBaseKey}.dangerZone`)}
      >
        <div className="grid w-full grid-cols-3 gap-4">
          <SettingEndOfLine
            title={
              flagEnable
                ? t(`${translationBaseKey}.status.disable.title`)
                : t(`${translationBaseKey}.status.enable.title`)
            }
            description={
              flagEnable
                ? t(`${translationBaseKey}.status.disable.description`)
                : t(`${translationBaseKey}.status.enable.description`)
            }
            component={
              <Button
                color={flagEnable ? "failure" : "blue"}
                className={"float-right min-w-64"}
                onClick={() => setOpenConfirmationDisable(true)}
              >
                {flagEnable ? (
                  <FaToggleOff className="mr-2 h-5 w-5" />
                ) : (
                  <FaToggleOn className="mr-2 h-5 w-5" />
                )}
                {flagEnable
                  ? t(`${translationBaseKey}.status.disable.button`)
                  : t(`${translationBaseKey}.status.enable.button`)}
              </Button>
            }
          />
          <hr className={"col-span-3 m-0 "} />
          <SettingEndOfLine
            title={t(`${translationBaseKey}.delete.title`)}
            description={t(`${translationBaseKey}.delete.description`)}
            component={
              <Button
                color="failure"
                className={"float-right min-w-64"}
                onClick={() => setOpenDeleteModal(true)}
              >
                <FaTrashAlt className="mr-2 h-4 w-4" />{" "}
                {t(`${translationBaseKey}.delete.button`)}
              </Button>
            }
          />
        </div>
      </Section>
      <ConfirmationModal
        text={t(`${translationBaseKey}.delete.modal`, { name })}
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
        text={
          flagEnable
            ? t(`${translationBaseKey}.status.disable.modal`, { name })
            : t(`${translationBaseKey}.status.enable.modal`, { name })
        }
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
    </>
  );
};

const SettingEndOfLine = ({
  title,
  description,
  component,
}: {
  title: string;
  description: string;
  component: React.ReactNode;
}) => {
  return (
    <>
      <div className="col-span-2 my-2">
        <div className="text-base font-medium text-gray-900 dark:text-gray-100">
          {title}
        </div>
        <div className="text-sm text-gray-500 dark:text-gray-400">
          {description}
        </div>
      </div>
      <div className={" w-full"}>{component}</div>
    </>
  );
};

const SettingText = ({
  title,
  description,
  component,
}: {
  title: string;
  description: string;
  component: ReactNode;
}) => {
  return (
    <>
      <div className=" my-2">
        <div className="text-base font-medium text-gray-900 dark:text-gray-100">
          {title}
        </div>
        <div className="text-sm text-gray-500 dark:text-gray-400">
          {description}
        </div>
      </div>
      <div className={"col-span-2 float-right w-full"}>{component}</div>
    </>
  );
};
