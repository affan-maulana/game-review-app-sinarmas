import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { GameList } from "@/components/games/GameList";

const mockGames = [
  { id: "1", title: "Elden Ring", genre: "Action RPG", platform: "PC" },
  { id: "2", title: "The Witcher 3", genre: "RPG", platform: "PC" },
];

describe("GameList", () => {
  it("renders games", () => {
    render(<GameList games={mockGames} loading={false} error={null} />);
    expect(screen.getByText("Elden Ring")).toBeInTheDocument();
    expect(screen.getByText("The Witcher 3")).toBeInTheDocument();
  });

  it("shows loading state", () => {
    render(<GameList games={[]} loading={true} error={null} />);
    expect(screen.getByText("Loading games...")).toBeInTheDocument();
  });

  it("shows error state", () => {
    render(<GameList games={[]} loading={false} error="Failed to load" />);
    expect(screen.getByText(/Failed to load/)).toBeInTheDocument();
  });

  it("shows empty state", () => {
    render(<GameList games={[]} loading={false} error={null} />);
    expect(screen.getByText("No games found.")).toBeInTheDocument();
  });
});
