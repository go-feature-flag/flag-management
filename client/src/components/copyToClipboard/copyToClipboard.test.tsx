import { act, fireEvent, render, screen } from "@testing-library/react";
import { afterEach, beforeEach, describe } from "vitest";
import { CopyToClipboard } from "./copyToClipboard.tsx";

describe("GoffBreadcrumb", () => {
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

  it("should display the copy button", () => {
    render(<CopyToClipboard textToCopy={"foobar"} />);
    expect(screen.getByRole("button")).toBeInTheDocument();
    expect(
      screen.getByText("component.copyToClipboard.copy"),
    ).toBeInTheDocument();
    expect(screen.getByTestId("copy-icon")).toBeInTheDocument();
  });

  it("should copy value change icon and text if button is clicked", async () => {
    const textToCopy = "foobar";
    render(<CopyToClipboard textToCopy={textToCopy} />);
    await act(async () => {
      fireEvent.click(screen.getByRole("button"));
    });
    expect(navigator.clipboard.writeText).toHaveBeenCalledWith(textToCopy);
    expect(
      screen.getByText("component.copyToClipboard.copied"),
    ).toBeInTheDocument();
    expect(screen.getByTestId("check-icon")).toBeInTheDocument();
  });

  it("should come back to initial text after timeDisplayCopied", async () => {
    const textToCopy = "foobar";
    render(<CopyToClipboard textToCopy={textToCopy} timeDisplayCopied={500} />);
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
