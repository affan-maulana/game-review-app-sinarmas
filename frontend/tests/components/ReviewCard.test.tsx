import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { ReviewCard } from "@/components/reviews/ReviewCard";

const mockReview = {
  id: "r1",
  gameId: "1",
  reviewerName: "John",
  text: "Great game!",
  rating: 5,
  createdAt: "2024-01-15T10:00:00Z",
};

describe("ReviewCard", () => {
  it("renders reviewer name", () => {
    render(<ReviewCard review={mockReview} />);
    expect(screen.getByText("John")).toBeInTheDocument();
  });

  it("renders review text", () => {
    render(<ReviewCard review={mockReview} />);
    expect(screen.getByText("Great game!")).toBeInTheDocument();
  });

  it("renders rating stars", () => {
    render(<ReviewCard review={mockReview} />);
    expect(screen.getByText("★★★★★")).toBeInTheDocument();
  });
});
