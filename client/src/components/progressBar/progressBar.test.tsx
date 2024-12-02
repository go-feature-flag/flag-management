import { render, screen } from "@testing-library/react";
import { expect } from "vitest";
import type { FormFlagPageVariationInfo } from "../../models/featureFlagFormData.ts";
import { PercentageProgressBar } from "./progressBar.tsx";

describe("ProgressBar", () => {
  const variations: FormFlagPageVariationInfo[] = [
    { name: "variation1", value: "value1" },
    { name: "variation2", value: "value2" },
    { name: "variation3", value: "value3" },
    { name: "variation4", value: "value4" },
  ];
  it("should have a portion of the progressbar for all the percentages", () => {
    const percentages = {
      variation1: 25,
      variation2: 25,
      variation3: 25,
      variation4: 25,
    };

    render(
      <PercentageProgressBar
        percentages={percentages}
        variations={variations}
      />,
    );

    expect(screen.getAllByTestId("percentage-progress-bar-item").length).toBe(
      variations.length,
    );
  });
  it("should not display in progressbar if no percentages", () => {
    const percentages = {
      variation1: 25,
      variation2: 25,
      variation4: 25,
    };

    render(
      <PercentageProgressBar
        percentages={percentages}
        variations={variations}
      />,
    );
    expect(screen.getAllByTestId("percentage-progress-bar-item").length).toBe(
      3,
    );
    expect(screen.getByTestId("percentage-progress-display")).toBeVisible();
    expect(screen.getByTestId("percentage-progress-display")).toHaveTextContent(
      "75%",
    );
  });
  it("should ignore percentages if not part of the variations", () => {
    const percentages = {
      variation1: 25,
      variation2: 25,
      variation423: 25,
    };
    render(
      <PercentageProgressBar
        percentages={percentages}
        variations={variations}
      />,
    );
    expect(screen.getAllByTestId("percentage-progress-bar-item").length).toBe(
      2,
    );
  });

  it("should display an error icon if percentage more that 100%", () => {
    const percentages = {
      variation1: 25,
      variation2: 25,
      variation4: 75,
    };
    render(
      <PercentageProgressBar
        percentages={percentages}
        variations={variations}
      />,
    );
    expect(screen.getByTestId("percentage-progress-bar-error")).toBeVisible();
    expect(
      screen.getByTestId("percentage-progress-bar-error-icon"),
    ).toBeVisible();
  });

  it("should display a check icon if percentage is 100%", () => {
    const percentages = {
      variation1: 25,
      variation2: 25,
      variation4: 50,
    };
    render(
      <PercentageProgressBar
        percentages={percentages}
        variations={variations}
      />,
    );
    expect(screen.getAllByTestId("percentage-progress-bar-item").length).toBe(
      3,
    );
    expect(
      screen.getByTestId("percentage-progress-bar-check-icon"),
    ).toBeVisible();
  });
});
