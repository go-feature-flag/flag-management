import { render, screen } from "@testing-library/react";
import { describe } from "vitest";
import Loader from "./loader.tsx";

describe("GoffBreadcrumb", () => {
  it("should display the flow-bite loader", () => {
    render(<Loader />);
    expect(screen.getByRole("status")).toBeVisible();
  });
});
