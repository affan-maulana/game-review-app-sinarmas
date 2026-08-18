import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { ReviewList } from "@/components/reviews/ReviewList";

const mockReviews = [
  {
    id: "r1",
    gameId: "1",
    reviewerName: "John",
    text: "Great game!",
    rating: 5,
    createdAt: "2024-01-15T10:00:00Z",
  },
  {
    id: "r2",
    gameId: "1",
    reviewerName: "Jane",
    text: "Good but buggy",
    rating: 3,
    createdAt: "2024-02-20T14:30:00Z",
  },
];

describe("ReviewList", () => {
  it("renders reviews", () => {
    render(<ReviewList reviews={mockReviews} loading={false} error={null} />);
    expect(screen.getByText("John")).toBeInTheDocument();
    expect(screen.getByText("Jane")).toBeInTheDocument();
  });

  it("shows loading state", () => {
    render(<ReviewList reviews={[]} loading={true} error={null} />);
    expect(screen.getByText("Loading reviews...")).toBeInTheDocument();
  });

  it("shows error state", () => {
    render(
      <ReviewList reviews={[]} loading={false} error="Failed to load" />
    );
    expect(screen.getByText(/Failed to load/)).toBeInTheDocument();
  });

  it("shows empty state", () => {
    render(<ReviewList reviews={[]} loading={false} error={null} />);
    expect(
      screen.getByText("No reviews yet. Be the first to review!")
    ).toBeInTheDocument();
  });
});
