import { Alert, Button, Modal, TextInput } from "flowbite-react";
import type { ReactElement } from "react";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import {
  HiInformationCircle,
  HiOutlineExclamationCircle,
} from "react-icons/hi2";
import type { IconType } from "react-icons/lib";

const baseTranslationKey = "component.modal";

interface DeleteModalProps {
  isOpen: boolean;
  text: string;
  okText?: string;
  cancelText?: string;
  onClickCancel?: () => void;
  onClickYes: () => void;
  icon?: ReactElement<IconType>;
  error?: string;
  // confirmationText is used to confirm the deletion of an item, if set the user will have to type the confirmationText to confirm the action
  confirmationText?: string;
}

export function ConfirmationModal(props: DeleteModalProps) {
  const [confirmationText, setConfirmationText] = useState("");
  const { t } = useTranslation();

  function isButtonDisabled() {
    return (
      props.confirmationText !== undefined &&
      confirmationText !== props.confirmationText
    );
  }

  return (
    <Modal show={props.isOpen} size="lg" onClose={props.onClickCancel} popup>
      <Modal.Header data-testid={"confirm-modal-header"} />
      <Modal.Body data-testid={"confirm-modal-body"}>
        <div className="text-center">
          {props.icon ?? (
            <HiOutlineExclamationCircle
              className="mx-auto mb-4 h-14 w-14"
              data-testid={"default-confirm-modal-icon"}
            />
          )}
          <h3 className="mb-5 text-lg font-normal">{props.text}</h3>
          {props.confirmationText && (
            <>
              <hr />
              <div className="mb-5 justify-center gap-4">
                <div className={"mb-1 text-left font-light"}>
                  Please type <strong>{props.confirmationText}</strong> to
                  confirm.
                </div>
                <TextInput
                  id="base"
                  type="text"
                  sizing="sm"
                  className={"w-full"}
                  placeholder={props.confirmationText}
                  onChange={(event) => setConfirmationText(event.target.value)}
                  data-testid={"confirm-modal-text-input"}
                />
              </div>
            </>
          )}
          <div className="flex justify-center gap-4">
            <Button
              color="failure"
              onClick={props.onClickYes}
              disabled={isButtonDisabled()}
              data-testid={"confirm-modal-yes-button"}
            >
              {props.okText ?? t(`${baseTranslationKey}.okText`)}
            </Button>
            <Button
              color="gray"
              onClick={props.onClickCancel}
              data-testid={"confirm-modal-cancel-button"}
            >
              {props.cancelText ?? t(`${baseTranslationKey}.cancelText`)}
            </Button>
          </div>
        </div>
        {props.error && (
          <Alert
            color="failure"
            icon={HiInformationCircle}
            data-testid={"confirm-modal-error"}
          >
            {props.error}
          </Alert>
        )}
      </Modal.Body>
    </Modal>
  );
}
