import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { GameCard } from "@/components/games/GameCard";

const mockGame = {
  id: "1",
  title: "Elden Ring",
  genre: "Action RPG",
  platform: "PC, PS5",
};

describe("GameCard", () => {
  it("renders game title", () => {
    render(<GameCard game={mockGame} />);
    expect(screen.getByText("Elden Ring")).toBeInTheDocument();
  });

  it("renders game genre", () => {
    render(<GameCard game={mockGame} />);
    expect(screen.getByText("Action RPG")).toBeInTheDocument();
  });

  it("renders game platform", () => {
    render(<GameCard game={mockGame} />);
    expect(screen.getByText("PC, PS5")).toBeInTheDocument();
  });

  it("links to game detail page", () => {
    render(<GameCard game={mockGame} />);
    const link = screen.getByRole("link");
    expect(link).toHaveAttribute("href", "/games/1");
  });
});
