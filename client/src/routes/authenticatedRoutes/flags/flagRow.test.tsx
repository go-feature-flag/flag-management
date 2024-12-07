import { render, screen } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";
import { Table, TableBody } from "flowbite-react";
import { MemoryRouter } from "react-router-dom";
import type { Mock } from "vitest";
import { beforeEach, expect, vi } from "vitest";
import { updateFeatureFlagStatusById } from "../../../api/goffApi.ts";
import type { FeatureFlagFormData } from "../../../models/featureFlagFormData.ts";
import { FlagRow } from "./flagRow.tsx";

vi.mock("../../../api/goffApi.ts", () => ({
  updateFeatureFlagStatusById: vi.fn(),
}));

describe("flag row", () => {
  const handleDelete = vi.fn();
  const handleDisable = vi.fn();
  let defaultFlag: FeatureFlagFormData;
  beforeEach(() => {
    vi.clearAllMocks();
    defaultFlag = {
      creationDate: new Date(2024, 0, 1),
      lastUpdatedDate: new Date(2024, 0, 2),
      id: "3b241101-e2bb-4255-8caf-4136c566a962",
      name: "my-feature-flag",
      description: "This is a feature flag",
      disable: false,
      trackEvent: true,
      type: "boolean",
      variations: [
        {
          id: "1b4e28ba-2fa1-11d2-883f-0016d3cca427",
          name: "enabled",
          value: true,
        },
        {
          id: "6fa459ea-ee8a-3ca4-894e-db77e160355e",
          name: "disabled",
          value: false,
        },
      ],
      defaultRule: {
        id: "16fd2706-8baf-433b-82eb-8c7fada847da",
        variationSelect: "6fa459ea-ee8a-3ca4-894e-db77e160355e",
      },

      metadata: [],
      targetingRules: [],
      version: "0.0.1",
    };
  });
  describe("labels", () => {
    it("should display a labels (version, variations, type) of the flag", () => {
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={defaultFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("label", { name: "0.0.1" })).toBeVisible();
      expect(screen.getByRole("label", { name: "2 variations" })).toBeVisible();
      expect(screen.getByRole("label", { name: "boolean" })).toBeVisible();
    });
    it("should display a label with the type of the flag", () => {
      const stringFlag: FeatureFlagFormData = {
        creationDate: new Date(2024, 0, 1),
        lastUpdatedDate: new Date(2024, 0, 2),
        id: "3b241101-e2bb-4255-8caf-4136c566a962",
        name: "my-feature-flag",
        description: "This is a feature flag",
        disable: false,
        trackEvent: true,
        type: "string",
        variations: [
          {
            id: "1b4e28ba-2fa1-11d2-883f-0016d3cca427",
            name: "enabled",
            value: "true",
          },
          {
            id: "6fa459ea-ee8a-3ca4-894e-db77e160355e",
            name: "disabled",
            value: "false",
          },
          {
            id: "6fa459ea-ee8a-3ca4-894e-db77e160355f",
            name: "yo",
            value: "yo",
          },
        ],
        defaultRule: {
          id: "16fd2706-8baf-433b-82eb-8c7fada847da",
          variationSelect: "6fa459ea-ee8a-3ca4-894e-db77e160355e",
        },

        metadata: [],
        targetingRules: [],
        version: "0.0.1",
      };
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={stringFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("label", { name: "3 variations" })).toBeVisible();
      expect(screen.getByRole("label", { name: "string" })).toBeVisible();
    });
    it("should not display label version if no version specified for the flag", () => {
      const stringFlag: FeatureFlagFormData = {
        creationDate: new Date(2024, 0, 1),
        lastUpdatedDate: new Date(2024, 0, 2),
        id: "3b241101-e2bb-4255-8caf-4136c566a962",
        name: "my-feature-flag",
        description: "This is a feature flag",
        disable: false,
        trackEvent: true,
        type: "string",
        variations: [
          {
            id: "6fa459ea-ee8a-3ca4-894e-db77e160355e",
            name: "disabled",
            value: "false",
          },
          {
            id: "6fa459ea-ee8a-3ca4-894e-db77e160355f",
            name: "yo",
            value: "yo",
          },
        ],
        defaultRule: {
          id: "16fd2706-8baf-433b-82eb-8c7fada847da",
          variationSelect: "6fa459ea-ee8a-3ca4-894e-db77e160355e",
        },

        metadata: [],
        targetingRules: [],
      };
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={stringFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.queryByTestId("label-flag-version")).toBeNull();
    });
  });

  describe("enable/disable flag", () => {
    it("should display a button to disable the flag", () => {
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={defaultFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("true");
    });

    it("should display a button to enable the flag", () => {
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={{ ...defaultFlag, disable: true }}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("false");
    });

    it("should ask for confirmation if click on enable button", async () => {
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={defaultFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("true");

      await userEvent.click(screen.getByRole("toggle-switch"));
      expect(
        screen.getByText("page.flags.flagList.row.modal.enable"),
      ).toBeVisible();
    });

    it("should not change value if click cancel", async () => {
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={defaultFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("true");

      await userEvent.click(screen.getByRole("toggle-switch"));
      expect(
        screen.getByText("page.flags.flagList.row.modal.enable"),
      ).toBeVisible();

      await userEvent.click(
        screen.getByRole("button", {
          name: "component.modal.cancelText",
        }),
      );
      expect(
        screen.queryByText("page.flags.flagList.row.modal.enable"),
      ).toBeNull();

      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("true");
    });

    it("should disable flag if click on yes", async () => {
      const mockUpdateFeatureFlagStatusById =
        updateFeatureFlagStatusById as Mock;
      mockUpdateFeatureFlagStatusById.mockResolvedValue({ status: 200 });
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={defaultFlag}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("true");

      await userEvent.click(screen.getByRole("toggle-switch"));
      expect(
        screen.getByText("page.flags.flagList.row.modal.enable"),
      ).toBeVisible();

      await userEvent.click(
        screen.getByRole("button", {
          name: "component.modal.okText",
        }),
      );

      // we close the modal
      expect(
        screen.queryByText("page.flags.flagList.row.modal.enable"),
      ).toBeNull();
      expect(mockUpdateFeatureFlagStatusById).toHaveBeenCalledWith(
        "3b241101-e2bb-4255-8caf-4136c566a962",
        { disable: true },
      );
      expect(handleDisable).toHaveBeenCalledWith(defaultFlag.id);
    });

    it("should enable flag if click on yes", async () => {
      const mockUpdateFeatureFlagStatusById =
        updateFeatureFlagStatusById as Mock;
      mockUpdateFeatureFlagStatusById.mockResolvedValue({ status: 200 });
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={{ ...defaultFlag, disable: true }}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("false");

      await userEvent.click(screen.getByRole("toggle-switch"));
      expect(
        screen.getByText("page.flags.flagList.row.modal.enable"),
      ).toBeVisible();

      await userEvent.click(
        screen.getByRole("button", {
          name: "component.modal.okText",
        }),
      );

      // we close the modal
      expect(
        screen.queryByText("page.flags.flagList.row.modal.enable"),
      ).toBeNull();
      expect(mockUpdateFeatureFlagStatusById).toHaveBeenCalledWith(
        "3b241101-e2bb-4255-8caf-4136c566a962",
        { disable: false },
      );
      expect(handleDisable).toHaveBeenCalledWith(defaultFlag.id);
    });

    it("should display an error in the modal if API is throwing an error", async () => {
      const mockUpdateFeatureFlagStatusById =
        updateFeatureFlagStatusById as Mock;
      mockUpdateFeatureFlagStatusById.mockRejectedValue(
        new Error("Internal Server Error"),
      );
      render(
        <MemoryRouter>
          <Table hoverable>
            <TableBody className="divide-y">
              <FlagRow
                handleDelete={handleDelete}
                handleDisable={handleDisable}
                flag={{ ...defaultFlag, disable: true }}
              />
            </TableBody>
          </Table>
        </MemoryRouter>,
      );
      expect(screen.getByRole("toggle-switch")).toBeVisible();
      expect(
        screen.getByRole("toggle-switch").getAttribute("aria-checked"),
      ).toBe("false");

      await userEvent.click(screen.getByRole("toggle-switch"));
      expect(
        screen.getByText("page.flags.flagList.row.modal.enable"),
      ).toBeVisible();

      await userEvent.click(
        screen.getByRole("button", {
          name: "component.modal.okText",
        }),
      );

      // we close the modal
      expect(
        screen.getByText("page.flags.flagList.row.modal.enable"),
      ).toBeVisible();

      expect(mockUpdateFeatureFlagStatusById).toHaveBeenCalledWith(
        "3b241101-e2bb-4255-8caf-4136c566a962",
        { disable: false },
      );
      expect(handleDisable).not.toHaveBeenCalled();
      expect(
        screen.getByText(
          "page.flags.flagList.row.errors.statusChange Error: Internal Server Error",
        ),
      ).toBeVisible();
    });
  });
});