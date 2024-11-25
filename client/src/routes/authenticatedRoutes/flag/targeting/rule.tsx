import { useBooleanFlagValue } from "@openfeature/react-sdk";
import { Alert, Dropdown, FloatingLabel } from "flowbite-react";
import type { Dispatch, SetStateAction } from "react";
import { useState } from "react";
import type { DeepPartialSkipArrayKey } from "react-hook-form";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { BsThreeDotsVertical, BsWrenchAdjustable } from "react-icons/bs";
import { FaInfoCircle, FaTrashAlt } from "react-icons/fa";
import { FaToggleOff, FaToggleOn } from "react-icons/fa6";
import { HiOutlinePencilSquare } from "react-icons/hi2";
import { ConfirmationModal } from "../../../../components/modal/ConfirmationModal.tsx";
import type { FormFlagPageTargetingRule } from "../../../../models/featureFlagFormData.ts";
import { VariationSelector } from "../components/variationSelector.tsx";
import { RuleAdvancedQuery } from "./ruleAdvancedQuery.tsx";
import { RuleDisabled } from "./ruleDisabled.tsx";
import { RuleQueryBuilder } from "./ruleQueryBuilder.tsx";

const translationBaseKey = "page.flag.targeting.rule";

export const Rule = ({
  rule,
  handleDelete,
  fieldName,
}: {
  rule: DeepPartialSkipArrayKey<FormFlagPageTargetingRule>;
  handleDelete: (ruleId?: string) => void;
  fieldName: string;
}) => {
  const { t } = useTranslation();
  const appEnableQueryBuilder = useBooleanFlagValue(
    "app-enable-query-builder",
    false,
  );
  // if the app-enable-query-builder flag is not enabled, we should default to the advanced query, else we should default to the query builder.
  const [advancedQuery, setAdvancedQuery] = useState(!appEnableQueryBuilder);
  const [confirmDelete, setConfirmDelete] = useState(false);
  return (
    <>
      {rule.disabled && (
        <Alert color="failure" icon={FaInfoCircle}>
          This rule is disabled
        </Alert>
      )}
      <RuleNameBar
        advancedQuery={advancedQuery}
        setAdvancedQuery={setAdvancedQuery}
        fieldName={fieldName}
        handleDelete={() => setConfirmDelete(true)}
        disabled={rule.disabled ?? false}
      />
      {rule.disabled ? (
        <RuleDisabled fieldName={fieldName} />
      ) : advancedQuery ? (
        <RuleAdvancedQuery fieldName={fieldName} />
      ) : (
        <RuleQueryBuilder />
      )}
      <div className={"mt-3"}>
        <VariationSelector
          baseName={`${fieldName}.variation`}
          disabled={rule.disabled}
        />
      </div>
      <ConfirmationModal
        isOpen={confirmDelete}
        text={t(`${translationBaseKey}.deleteConfirmation`, {
          name: rule.name,
        })}
        onClickCancel={() => setConfirmDelete(false)}
        onClickYes={() => {
          handleDelete(rule.variation?.id);
          setConfirmDelete(false);
        }}
      />
    </>
  );
};

const DisableRuleDropdownItem = ({ fieldName }: { fieldName: string }) => {
  const { watch, setValue } = useFormContext();
  const { t } = useTranslation();

  const buttonFieldName = `${fieldName}.disabled`;
  const isDisabled = watch(buttonFieldName);

  return (
    <>
      {isDisabled ? (
        <Dropdown.Item
          icon={FaToggleOn}
          onClick={() =>
            setValue(buttonFieldName, false, { shouldDirty: true })
          }
        >
          {t(`${translationBaseKey}.dropdown.enable`)}
        </Dropdown.Item>
      ) : (
        <Dropdown.Item
          icon={FaToggleOff}
          onClick={() => setValue(buttonFieldName, true, { shouldDirty: true })}
        >
          {t(`${translationBaseKey}.dropdown.disable`)}
        </Dropdown.Item>
      )}
    </>
  );
};

const RuleNameBar = ({
  advancedQuery,
  setAdvancedQuery,
  handleDelete,
  fieldName,
  disabled,
}: {
  advancedQuery: boolean;
  setAdvancedQuery: Dispatch<SetStateAction<boolean>>;
  fieldName: string;
  handleDelete: () => void;
  disabled: boolean;
}) => {
  const { t } = useTranslation();
  const { register } = useFormContext();

  const appEnableQueryBuilder = useBooleanFlagValue(
    "app-enable-query-builder",
    true,
  );

  return (
    <div className="flex items-center">
      <div className="grow">
        <FloatingLabel
          label={"Rule name"}
          sizing={"sm"}
          variant={"standard"}
          {...register(`${fieldName}.name`)}
          disabled={disabled}
        />
      </div>
      <div className="ml-auto shrink-0">
        <Dropdown
          label={<BsThreeDotsVertical className="h-6 w-6" />}
          arrowIcon={false}
          inline
        >
          {advancedQuery && appEnableQueryBuilder && (
            <Dropdown.Item
              icon={BsWrenchAdjustable}
              onClick={() => setAdvancedQuery(!advancedQuery)}
            >
              {t(`${translationBaseKey}.dropdown.queryBuilder`)}
            </Dropdown.Item>
          )}
          {!advancedQuery && (
            <Dropdown.Item
              icon={HiOutlinePencilSquare}
              onClick={() => setAdvancedQuery(!advancedQuery)}
            >
              {t(`${translationBaseKey}.dropdown.advanced`)}
            </Dropdown.Item>
          )}
          <Dropdown.Divider />
          <DisableRuleDropdownItem fieldName={fieldName} />
          <Dropdown.Item icon={FaTrashAlt} onClick={handleDelete}>
            {t(`${translationBaseKey}.dropdown.delete`)}
          </Dropdown.Item>
        </Dropdown>
      </div>
    </div>
  );
};
