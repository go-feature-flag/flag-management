import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import { NewFlagModal } from "./newFlagModal.tsx";

describe("New flag modal", () => {
  it("should render the new flag modal", () => {
    render(
      <MemoryRouter>
        <NewFlagModal newFlag={true} setNewFlag={() => {}} featureFlags={[]} />
      </MemoryRouter>,
    );
    expect(
      screen.getByAltText("page.flags.newModal.labelNewFlag"),
    ).toBeVisible();
  });
});
