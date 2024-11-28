import { act, fireEvent, render, screen } from "@testing-library/react";
import { afterEach, beforeEach, describe } from "vitest";
import { CopyToClipboard } from "./copyToClipboard.tsx";

describe("Copy To Clipboard", () => {
  beforeEach(() => {
    Object.assign(navigator, {
      clipboard: {
        writeText: vi.fn(),
      },
    });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  it("should have only icon visible by default (no text)", () => {
    render(<CopyToClipboard textToCopy={"foobar"} />);
    expect(screen.getByTestId("copy-to-clipboard-button")).toBeVisible();
    expect(screen.queryByText("component.copyToClipboard.copy")).toBeNull();
    expect(screen.getByTestId("copy-icon")).toBeVisible();
  });

  it("should have text visible if copy to clipboard size sm", () => {
    render(<CopyToClipboard textToCopy={"foobar"} size={"sm"} />);
    expect(screen.getByTestId("copy-to-clipboard-button")).toBeVisible();
    expect(screen.getByText("component.copyToClipboard.copy")).toBeVisible();
    expect(screen.getByTestId("copy-icon")).toBeVisible();
  });

  it("should have a tooltip if text provided", () => {
    const tooltipTxt = "this is a tooltip";
    render(<CopyToClipboard textToCopy={"foobar"} tooltipText={tooltipTxt} />);
    expect(screen.getByTestId("tooltip-copy-clipboard")).toBeVisible();
    expect(screen.getByTestId("tooltip-copy-clipboard")).toHaveTextContent(
      tooltipTxt,
    );
  });

  it("should have no tooltip if no text provided", () => {
    render(<CopyToClipboard textToCopy={"foobar"} />);
    expect(screen.queryByTestId("tooltip-copy-clipboard")).toBeNull();
  });

  it("should copy value change icon and text if button is clicked", async () => {
    const textToCopy = "foobar";
    render(<CopyToClipboard textToCopy={textToCopy} size={"sm"} />);
    await act(async () => {
      fireEvent.click(screen.getByTestId("copy-to-clipboard-button"));
    });
    expect(navigator.clipboard.writeText).toHaveBeenCalledWith(textToCopy);
    expect(
      screen.getByText("component.copyToClipboard.copied"),
    ).toBeInTheDocument();
    expect(screen.getByTestId("check-icon")).toBeInTheDocument();
  });

  it("should come back to initial text after timeDisplayCopied", async () => {
    const textToCopy = "foobar";
    render(
      <CopyToClipboard
        textToCopy={textToCopy}
        timeDisplayCopied={500}
        size={"sm"}
      />,
    );
    await act(async () => {
      fireEvent.click(screen.getByRole("button"));
    });
    expect(navigator.clipboard.writeText).toHaveBeenCalledWith(textToCopy);

    // wait until the text is back to initial
    await act(async () => {
      await new Promise((r) => setTimeout(r, 600));
    });

    expect(
      screen.getByText("component.copyToClipboard.copy"),
    ).toBeInTheDocument();
    expect(screen.getByTestId("copy-icon")).toBeInTheDocument();
  });
});
