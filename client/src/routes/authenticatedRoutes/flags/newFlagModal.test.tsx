import { render, screen } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";
import { MemoryRouter } from "react-router-dom";
import { expect } from "vitest";
import { variationTypes } from "../flag/helpers/variations.ts";
import { NewFlagModal } from "./newFlagModal.tsx";

const mockedUsedNavigate = vi.fn();
describe("New flag modal", () => {
  it("should render the new flag modal", () => {
    render(
      <MemoryRouter>
        <NewFlagModal handleClose={vi.fn()} featureFlags={[]} />
      </MemoryRouter>,
    );
    expect(
      screen.getByAltText("page.flags.newModal.labelNewFlag"),
    ).toBeVisible();
  });

  it("should have a radio button for all variation types items", () => {
    render(
      <MemoryRouter>
        <NewFlagModal handleClose={vi.fn()} featureFlags={[]} />
      </MemoryRouter>,
    );

    variationTypes.forEach((variation) => {
      expect(
        screen.getByRole("radio", { name: variation.displayName }),
      ).toBeVisible();
    });

    // expect the first radio button to be checked
    expect(
      screen.getByRole("radio", { name: variationTypes[0].displayName }),
    ).toBeChecked();
  });

  it("should change radio checked if user click", async () => {
    render(
      <MemoryRouter>
        <NewFlagModal handleClose={vi.fn()} featureFlags={[]} />
      </MemoryRouter>,
    );
    expect(
      screen.getByRole("radio", { name: variationTypes[1].displayName }),
    ).not.toBeChecked();

    await userEvent.click(
      screen.getByRole("radio", { name: variationTypes[1].displayName }),
    );
    expect(
      screen.getByRole("radio", { name: variationTypes[1].displayName }),
    ).toBeChecked();

    // read value of the radio input field
    const selectedRadio = screen.getByRole("radio", { checked: true });
    const selectedValue = selectedRadio.getAttribute("value");
    expect(selectedValue).toBe(variationTypes[1].type);
  });

  it("should display an error if flag name is empty", async () => {
    render(
      <MemoryRouter>
        <NewFlagModal handleClose={vi.fn()} featureFlags={[]} />
      </MemoryRouter>,
    );

    await userEvent.click(
      screen.getByRole("button", { name: "page.flags.newModal.createButton" }),
    );
    expect(
      screen.getByText("page.flags.newModal.error.required"),
    ).toBeVisible();
    expect(
      screen.getByLabelText("page.flags.newModal.labelNewFlag"),
    ).toBeInvalid();
  });

  it("should display an error if flag name already exists", async () => {
    render(
      <MemoryRouter>
        <NewFlagModal handleClose={vi.fn()} featureFlags={["flag1"]} />
      </MemoryRouter>,
    );

    await userEvent.type(
      screen.getByLabelText("page.flags.newModal.labelNewFlag"),
      "flag1",
    );
    await userEvent.click(
      screen.getByRole("button", { name: "page.flags.newModal.createButton" }),
    );
    expect(screen.getByText("page.flags.newModal.error.exists")).toBeVisible();
  });

  it("should redirect when creating a new flag", async () => {
    vi.mock("react-router-dom", async () => {
      const actual = await vi.importActual("react-router-dom");
      return {
        ...actual,
        useNavigate: () => mockedUsedNavigate,
      };
    });

    render(
      <MemoryRouter>
        <NewFlagModal handleClose={vi.fn()} featureFlags={[]} />,
      </MemoryRouter>,
    );
    await userEvent.click(
      screen.getByRole("radio", { name: variationTypes[4].displayName }),
    );
    await userEvent.type(
      screen.getByLabelText("page.flags.newModal.labelNewFlag"),
      "flag1",
    );
    await userEvent.click(
      screen.getByRole("button", { name: "page.flags.newModal.createButton" }),
    );

    expect(mockedUsedNavigate).toHaveBeenCalledWith(
      `/flags/new/flag1?type=${variationTypes[4].type}`,
    );
  });

  it("should close the modal on click on cancel cross", async () => {
    const handleClose = vi.fn();
    render(
      <MemoryRouter>
        <NewFlagModal handleClose={handleClose} featureFlags={[]} />,
      </MemoryRouter>,
    );
    expect(
      screen.getByAltText("page.flags.newModal.labelNewFlag"),
    ).toBeVisible();
    await userEvent.click(screen.getByRole("button", { name: "Close" }));
    expect(handleClose).toBeCalled();
  });
});
