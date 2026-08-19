"use client";

import { Game } from "@/types/game";
import { GameCard } from "./GameCard";

interface GameListProps {
  games: Game[];
  loading: boolean;
  error: string | null;
}

export function GameList({ games, loading, error }: GameListProps) {
  if (loading) {
    return <p className="text-gray-300">Loading games...</p>;
  }

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (games.length === 0) {
    return <p className="text-gray-300">No games found.</p>;
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      {games.map((game) => (
        <GameCard key={game.id} game={game} />
      ))}
    </div>
  );
}
