import { describe, it, expect, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { ReviewForm } from "@/components/reviews/ReviewForm";

describe("ReviewForm", () => {
  it("renders form fields", () => {
    const onSubmit = vi.fn();
    render(
      <ReviewForm onSubmit={onSubmit} isSubmitting={false} submitError={null} />
    );

    expect(screen.getByLabelText("Your Name")).toBeInTheDocument();
    expect(screen.getByLabelText("Rating")).toBeInTheDocument();
    expect(screen.getByLabelText("Review")).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: "Submit Review" })
    ).toBeInTheDocument();
  });

  it("shows validation errors when submitting empty form", async () => {
    const onSubmit = vi.fn();
    const user = userEvent.setup();

    render(
      <ReviewForm onSubmit={onSubmit} isSubmitting={false} submitError={null} />
    );

    await user.click(screen.getByRole("button", { name: "Submit Review" }));

    expect(screen.getByText("Reviewer name is required")).toBeInTheDocument();
    expect(screen.getByText("Review text is required")).toBeInTheDocument();
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it("submits form with valid data", async () => {
    const onSubmit = vi.fn().mockResolvedValue(undefined);
    const user = userEvent.setup();

    render(
      <ReviewForm onSubmit={onSubmit} isSubmitting={false} submitError={null} />
    );

    await user.type(screen.getByLabelText("Your Name"), "John");
    await user.type(screen.getByLabelText("Review"), "Great game!");
    await user.click(screen.getByRole("button", { name: "Submit Review" }));

    expect(onSubmit).toHaveBeenCalledWith({
      reviewerName: "John",
      text: "Great game!",
      rating: 5,
    });
  });

  it("displays submit error", () => {
    const onSubmit = vi.fn();

    render(
      <ReviewForm
        onSubmit={onSubmit}
        isSubmitting={false}
        submitError="Server error"
      />
    );

    expect(screen.getByText("Server error")).toBeInTheDocument();
  });
});
