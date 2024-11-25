import { Alert, Button, FloatingLabel, Tooltip } from "flowbite-react";
import { useState } from "react";
import { useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { FaTrashAlt } from "react-icons/fa";
import { FaCirclePlus } from "react-icons/fa6";
import { v4 as uuidv4 } from "uuid";
import { Section } from "../../../../components/section/section";
import {
  extractValidationError,
  hasError,
} from "../../../../helpers/extractValidationError.ts";
import type { FeatureFlagFormData } from "../../../../models/featureFlagFormData.ts";

const translationBaseKey = "page.flag.metadata";
export const Metadata = () => {
  const { t } = useTranslation();
  const { metadata } = useWatch<FeatureFlagFormData>();
  const { setValue } = useFormContext();

  function addItem() {
    const v = metadata ?? [];
    v.push({ name: "", value: "", id: uuidv4() });
    setValue("metadata", v, { shouldValidate: true, shouldDirty: true });
  }

  function removeItem(id: string) {
    setValue("metadata", metadata?.filter((m) => m.id !== id));
  }

  return (
    <Section
      outsideTitle={t(`${translationBaseKey}.title`)}
      info={t(`${translationBaseKey}.description`)}
    >
      {metadata?.length === 0 && <NoMetadata />}
      {metadata && metadata?.length > 0 && (
        <div className={"my-6"}>
          {metadata?.map((m, index) => {
            return (
              <MetadataLine
                key={m.id}
                id={m.id ?? ""}
                name={`metadata[${index}]`}
                onClickRemoveButton={removeItem}
              />
            );
          })}
        </div>
      )}

      <Button className="mt-4" onClick={addItem}>
        <FaCirclePlus className={"mr-2 h-4 w-4"} />
        Add metadata
      </Button>
    </Section>
  );
};

const NoMetadata = () => {
  const { t } = useTranslation();
  const [displayInfo, setDisplayInfo] = useState(true);
  return (
    <>
      {displayInfo && (
        <Alert color="success" onDismiss={() => setDisplayInfo(false)}>
          <span className="font-bold">Info!</span>
          <br /> {t(`${translationBaseKey}.messageNoMetadata`)}
        </Alert>
      )}
    </>
  );
};

const MetadataLine = ({
  name,
  id,
  onClickRemoveButton,
}: {
  name: string;
  id: string;
  onClickRemoveButton?: (id: string) => void;
}) => {
  const { t } = useTranslation();
  const {
    register,
    formState: { errors },
  } = useFormContext();
  return (
    <div className={"my-2 grid grid-cols-11 gap-4"}>
      <div className={"col-span-5"}>
        <FloatingLabel
          {...register(`${name}.name`, {
            required: t(`${translationBaseKey}.fields.name.errorRequired`),
          })}
          variant="outlined"
          label={t(`${translationBaseKey}.fields.name.title`)}
          placeholder={t(`${translationBaseKey}.fields.name.title`)}
          color={hasError(errors, `${name}.value`) ? "error" : "default"}
          helperText={extractValidationError(errors, `${name}.name`)?.message}
        />
      </div>
      <div className={"col-span-5"}>
        <FloatingLabel
          {...register(`${name}.value`, {
            required: t(`${translationBaseKey}.fields.value.errorRequired`),
          })}
          variant="outlined"
          label={t(`${translationBaseKey}.fields.value.title`)}
          placeholder={t(`${translationBaseKey}.fields.value.title`)}
          color={hasError(errors, `${name}.value`) ? "error" : "default"}
          helperText={extractValidationError(errors, `${name}.value`)?.message}
        />
      </div>
      <div className={"mt-4"}>
        <Tooltip content={t(`${translationBaseKey}.removeButton`)}>
          <button
            type="button"
            className={"text-gray-800 dark:text-gray-50"}
            onClick={
              onClickRemoveButton
                ? () => onClickRemoveButton(id)
                : () => {
                    return;
                  }
            }
          >
            <FaTrashAlt className="h-4 w-4" />
          </button>
        </Tooltip>
      </div>
    </div>
  );
};
