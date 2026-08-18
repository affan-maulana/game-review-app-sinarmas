import { apiClient } from "@/lib/api/client";
import { routes } from "@/lib/api/routes";
import { Game } from "@/types/game";

export const gameService = {
  getGames: () => apiClient.get<Game[]>(routes.games.list),

  getGame: (id: string) => apiClient.get<Game>(routes.games.detail(id)),
};
