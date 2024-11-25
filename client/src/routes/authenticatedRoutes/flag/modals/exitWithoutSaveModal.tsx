import { useTranslation } from "react-i18next";
import { ConfirmationModal } from "../../../../components/modal/ConfirmationModal.tsx";

export const ExitWithoutSaveModal = ({
  isOpen,
  onClickYes,
  onClickCancel,
}: {
  isOpen: boolean;
  onClickYes: () => void;
  onClickCancel: () => void;
}) => {
  const { t } = useTranslation();
  return (
    <ConfirmationModal
      isOpen={isOpen}
      text={t("component.exitWithoutSaveModal.text")}
      onClickYes={onClickYes}
      onClickCancel={onClickCancel}
    />
  );
};
