import { Game } from "@/types/game";

interface GameDetailsProps {
  game: Game;
}

export function GameDetails({ game }: GameDetailsProps) {
  return (
    <div className="rounded-lg border border-gray-200 p-6">
      <h1 className="text-2xl font-bold text-gray-900">{game.title}</h1>
      {game.genre && (
        <p className="text-gray-600 mt-2">
          <span className="font-medium">Genre:</span> {game.genre}
        </p>
      )}
      {game.platform && (
        <p className="text-gray-600 mt-1">
          <span className="font-medium">Platform:</span> {game.platform}
        </p>
      )}
    </div>
  );
}
