"use client";

import { useState, useEffect } from "react";
import { Game } from "@/types/game";
import { gameService } from "@/services/game.service";

export function useGame(id: string) {
  const [game, setGame] = useState<Game | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchGame = async () => {
      try {
        setLoading(true);
        const data = await gameService.getGame(id);
        setGame(data);
        setError(null);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Failed to load game details"
        );
      } finally {
        setLoading(false);
      }
    };

    fetchGame();
  }, [id]);

  return { game, loading, error };
}
