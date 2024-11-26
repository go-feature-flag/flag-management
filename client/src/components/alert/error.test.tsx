import { render } from "@testing-library/react";
import { FaAnchor } from "react-icons/fa6";
import { describe, expect, it } from "vitest";
import { AlertError } from "./error.tsx";

describe("Error Alert", () => {
  describe("display text", () => {
    it("should display the text in the document", () => {
      const title = "my custom text";
      const { getByText } = render(<AlertError text={title} />);
      expect(getByText(title)).toBeInTheDocument();
    });
  });

  describe("icon", () => {
    it("should display the default icon", () => {
      const title = "my custom text";
      const { getByTestId, queryByTestId } = render(
        <AlertError text={title} />,
      );
      expect(getByTestId("flowbite-alert-icon")).toBeInTheDocument();
      expect(queryByTestId("custom-icon")).toBeNull();
    });

    it("should display the custom icon", () => {
      const title = "my custom text";
      const { getByTestId, queryByTestId } = render(
        <AlertError
          text={title}
          icon={() => <FaAnchor data-testid="custom-icon" />}
        />,
      );
      expect(queryByTestId("flowbite-alert-icon")).toBeNull();
      expect(getByTestId("custom-icon")).toBeInTheDocument();
    });
  });

  describe("refresh button", () => {
    it("should have a refresh button if propsRefresh is true", () => {
      const { getByRole } = render(
        <AlertError text={"foo"} proposeRefresh={true} />,
      );
      expect(
        getByRole("button", {
          name: "component.alert.refreshButton",
        }),
      ).toBeInTheDocument();
    });

    it("should not have a refresh button if propsRefresh is false", () => {
      const { queryByRole } = render(
        <AlertError text={"foo"} proposeRefresh={false} />,
      );
      expect(
        queryByRole("button", {
          name: "component.alert.refreshButton",
        }),
      ).toBeNull();
    });

    it("should reload the page when the refresh button is clicked", () => {
      const { getByRole } = render(
        <AlertError text={"foo"} proposeRefresh={true} />,
      );
      const refreshButton = getByRole("button", {
        name: "component.alert.refreshButton",
      });

      const reloadMock = vi.fn();
      Object.defineProperty(window, "location", {
        value: {
          reload: reloadMock,
        },
        writable: true,
      });
      refreshButton.click();
      expect(reloadMock).toHaveBeenCalled();
      reloadMock.mockRestore();
    });
  });
});
