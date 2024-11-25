import { Button, Label, Modal, Radio, TextInput } from "flowbite-react";
import { useRef, useState } from "react";
import { Trans, useTranslation } from "react-i18next";
import { HiOutlineInformationCircle } from "react-icons/hi2";
import { useNavigate } from "react-router-dom";
import newFlagIcon from "../../../assets/pages/flags/newflag.svg";
import type { FeatureFlagFormData } from "../../../models/featureFlagFormData.ts";
import { variationTypes } from "../flag/helpers/variations.ts";

const translationBaseKey = "page.flags.newModal";
export const NewFlagModal = ({
  newFlag,
  setNewFlag,
  featureFlags,
}: {
  newFlag: boolean;
  setNewFlag: (flagState: boolean) => void;
  featureFlags: FeatureFlagFormData[];
}) => {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const flagNameRef = useRef<HTMLInputElement>(null);
  const [flagName, setFlagName] = useState("");
  const [flagType, setFlagType] = useState("boolean");
  const [errorMessage, setErrorMessage] = useState("");

  const handleCreate = () => {
    const message = isFlagNameValid(flagName);
    setErrorMessage(message);
    if (message.length > 0) {
      return;
    }
    navigate(`/flags/new/${encodeURIComponent(flagName)}?type=${flagType}`);
  };

  const isFlagNameValid = (flagName: string): string => {
    if (!flagName || flagName.length < 1) {
      return t(`${translationBaseKey}.error.required`);
    }

    const res =
      featureFlags.find((flag) => flag.name === flagName) === undefined;
    if (!res) {
      return t(`${translationBaseKey}.error.exists`);
    }
    return "";
  };

  return (
    <Modal
      show={newFlag}
      size="2xl"
      popup
      onClose={() => setNewFlag(false)}
      initialFocus={flagNameRef}
      dismissible
    >
      <Modal.Header />
      <Modal.Body>
        <div className="space-y-6">
          <h3 className="text-xl font-medium text-gray-900 dark:text-white">
            {t(`${translationBaseKey}.title`)}
          </h3>
          <div className={"flex max-h-72 w-full max-2xl:hidden"}>
            <img src={newFlagIcon} alt={"blabla"} />
          </div>
          <div>
            <Label
              htmlFor="flag_name"
              value={t(`${translationBaseKey}.labelNewFlag`)}
              className={"mb-1 block"}
            />{" "}
            <TextInput
              id="flag_name"
              ref={flagNameRef}
              placeholder="my-new-flag"
              required
              sizing={"lg"}
              value={flagName}
              color={errorMessage === "" ? "gray" : "failure"}
              onChange={(event) => setFlagName(event.target.value)}
              helperText={errorMessage}
            />
            <div
              className={
                "mt-1 flex space-x-4 text-xs font-light italic text-gray-400"
              }
            >
              <HiOutlineInformationCircle className={"h-4 w-4"} />
              <Trans
                i18nKey={`${translationBaseKey}.infoFlagName`}
                components={{ 1: <br /> }}
              />
            </div>
            <Label
              htmlFor="flag_name"
              value={t(`${translationBaseKey}.labelSelectType`)}
              className={"mb-1 mt-5 block"}
            />
            <div
              className={
                "grid grid-cols-2 space-y-2 md:grid-cols-5 md:space-y-0"
              }
            >
              {variationTypes.map((variationType, index) => (
                <div key={variationType.name}>
                  <Label
                    htmlFor={variationType.name}
                    className={
                      "inline-flex w-28 cursor-pointer items-center justify-between rounded-lg border border-gray-200 bg-white p-4 text-gray-500 hover:bg-gray-200 hover:text-gray-600 peer-checked:border-goff-600 peer-checked:text-goff-600 has-[:checked]:border-4 has-[:checked]:border-goff-300 has-[:checked]:bg-gray-200 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-900 dark:hover:text-gray-300 dark:peer-checked:text-goff-500 has-[:checked]:dark:border-4 has-[:checked]:dark:border-goff-300 has-[:checked]:dark:bg-gray-900"
                    }
                  >
                    <Radio
                      className="hidden"
                      name={"flag_type"}
                      id={variationType.name}
                      value={variationType.type}
                      defaultChecked={index === 0}
                      onChange={(event) => setFlagType(event.target.value)}
                    />
                    <div className="block">
                      <div className="w-full text-base font-semibold">
                        {variationType.displayName}
                      </div>
                    </div>
                    <variationType.icon className="mx-1 h-4 w-4" />
                  </Label>
                </div>
              ))}
            </div>
            <div
              className={
                "mt-1 flex space-x-4 text-xs font-light italic text-gray-400"
              }
            >
              <HiOutlineInformationCircle className={"h-4 w-4"} />
              <Trans i18nKey={`${translationBaseKey}.infoFlagType`} />
            </div>
          </div>

          <div className="w-full">
            <Button onClick={handleCreate}>Create</Button>
          </div>
        </div>
      </Modal.Body>
    </Modal>
  );
};
