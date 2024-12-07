import { fireEvent, render, screen } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";
import { HiExclamationCircle } from "react-icons/hi2";
import { describe, expect } from "vitest";
import { ConfirmationModal } from "./ConfirmationModal.tsx";

describe("ConfirmationModal", () => {
  it("should not render confirmation modal if not open", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={false}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
      />,
    );
    expect(screen.queryByTestId("confirmation-modal")).toBeNull();
  });

  it("should render confirmation modal if open", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
      />,
    );
    expect(screen.getByTestId("confirm-modal-header")).toBeVisible();
    expect(screen.getByTestId("confirm-modal-body")).toBeVisible();
    expect(screen.getByText("Are you sure?")).toBeVisible();
  });

  it("should use default text for button names", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
      />,
    );

    const yesButton = screen.getByTestId("confirm-modal-yes-button");
    expect(yesButton).toBeVisible();
    expect(yesButton).toHaveTextContent("component.modal.okText");

    const cancelButton = screen.getByTestId("confirm-modal-cancel-button");
    expect(cancelButton).toBeVisible();
    expect(cancelButton).toHaveTextContent("component.modal.cancelText");
  });

  it("should possible to override button names", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    const cancelButtonText = "Nope";
    const okButtonText = "Sure";
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
        cancelText={cancelButtonText}
        okText={okButtonText}
      />,
    );

    const yesButton = screen.getByTestId("confirm-modal-yes-button");
    expect(yesButton).toBeVisible();
    expect(yesButton).toHaveTextContent(okButtonText);

    const cancelButton = screen.getByTestId("confirm-modal-cancel-button");
    expect(cancelButton).toBeVisible();
    expect(cancelButton).toHaveTextContent(cancelButtonText);
  });

  it("should have a default icon if props not specified", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
      />,
    );
    expect(screen.getByTestId("default-confirm-modal-icon")).toBeVisible();
  });

  it("should have a custom icon if props not specified", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
        icon={<HiExclamationCircle data-testid="custom-confirm-modal-icon" />}
      />,
    );
    expect(screen.getByTestId("custom-confirm-modal-icon")).toBeVisible();
    expect(screen.queryByTestId("default-confirm-modal-icon")).toBeNull();
  });

  it("should call the onClickYes is yes button is clicked", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickYes={yesClick}
        onClickCancel={cancelClick}
      />,
    );

    fireEvent.click(screen.getByTestId("confirm-modal-yes-button"));
    expect(yesClick).toBeCalled();
  });

  it("should call the onClickYes is yes button is clicked", () => {
    const cancelClick = vi.fn();
    const yesClick = vi.fn();
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickCancel={cancelClick}
        onClickYes={yesClick}
      />,
    );

    fireEvent.click(screen.getByTestId("confirm-modal-cancel-button"));
    expect(cancelClick).toBeCalled();
  });

  it("should ask for confirmation when clicking on yes", async () => {
    const yesClick = vi.fn();
    const cancelClick = vi.fn();
    const confirmationText = "ok to delete";

    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickCancel={cancelClick}
        onClickYes={yesClick}
        confirmationText={confirmationText}
      />,
    );

    expect(yesClick).not.toBeCalled();
    expect(screen.getByTestId("confirm-modal-text-input")).toBeVisible();
    expect(screen.getByTestId("confirm-modal-yes-button")).toBeDisabled();

    fireEvent.focus(screen.getByTestId("confirm-modal-text-input"));
    await userEvent.type(
      screen.getByTestId("confirm-modal-text-input"),
      "not ok to delete",
    );
    expect(screen.getByTestId("confirm-modal-yes-button")).toBeDisabled();

    await userEvent.clear(screen.getByTestId("confirm-modal-text-input"));
    await userEvent.type(
      screen.getByTestId("confirm-modal-text-input"),
      confirmationText,
    );

    expect(screen.getByTestId("confirm-modal-yes-button")).not.toBeDisabled();
    await userEvent.click(screen.getByTestId("confirm-modal-yes-button"));
    expect(yesClick).toBeCalled();
  });

  it("should display an error if error in props", () => {
    const yesClick = vi.fn();
    const cancelClick = vi.fn();
    const confirmationText = "ok to delete";
    const errMsg = "Error message";
    render(
      <ConfirmationModal
        isOpen={true}
        text={"Are you sure?"}
        onClickCancel={cancelClick}
        onClickYes={yesClick}
        confirmationText={confirmationText}
        error={errMsg}
      />,
    );
    expect(screen.getByTestId("confirm-modal-error")).toBeVisible();
    expect(screen.getByTestId("confirm-modal-error")).toHaveTextContent(errMsg);
  });
});
