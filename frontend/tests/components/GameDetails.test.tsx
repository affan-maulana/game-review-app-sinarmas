import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { GameDetails } from "@/components/games/GameDetails";

const mockGame = {
  id: "1",
  title: "Elden Ring",
  genre: "Action RPG",
  platform: "PC, PS5, Xbox",
};

describe("GameDetails", () => {
  it("renders game title", () => {
    render(<GameDetails game={mockGame} />);
    expect(screen.getByText("Elden Ring")).toBeInTheDocument();
  });

  it("renders game genre", () => {
    render(<GameDetails game={mockGame} />);
    expect(screen.getByText("Action RPG")).toBeInTheDocument();
  });

  it("renders game platform", () => {
    render(<GameDetails game={mockGame} />);
    expect(screen.getByText("PC, PS5, Xbox")).toBeInTheDocument();
  });
});
