import { Alert, Button } from "flowbite-react";
import { useTranslation } from "react-i18next";
import { HiExclamationCircle } from "react-icons/hi2";
import type { IconType } from "react-icons/lib";

const translationBaseKey = "component.alert";
export const AlertError = ({
  text,
  icon,
  proposeRefresh,
  errorDetails,
}: {
  text: string;
  icon?: IconType;
  proposeRefresh?: boolean;
  errorDetails?: string;
}) => {
  return (
    <Alert
      color="failure"
      icon={icon ?? HiExclamationCircle}
      additionalContent={proposeRefresh && <RefreshButton />}
    >
      <p>{text}</p>
      {errorDetails && <p>{errorDetails}</p>}
    </Alert>
  );
};

function RefreshButton() {
  const { t } = useTranslation();
  return (
    <Button
      color="failure"
      size={"xs"}
      onClick={() => {
        window.location.reload();
      }}
    >
      {t(`${translationBaseKey}.refreshButton`)}
    </Button>
  );
}
