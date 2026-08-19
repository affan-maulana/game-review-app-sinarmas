"use client";

import { GameList } from "@/components/games/GameList";
import { useGames } from "@/hooks/useGames";

export default function HomePage() {
  const { games, loading, error } = useGames();

  return (
    <div>
      <h1 className="text-2xl font-bold mb-6 text-white">Games</h1>
      <GameList games={games} loading={loading} error={error} />
    </div>
  );
}
