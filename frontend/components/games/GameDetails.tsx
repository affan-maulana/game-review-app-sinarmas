import { Game } from "@/types/game";

interface GameDetailsProps {
  game: Game;
}

export function GameDetails({ game }: GameDetailsProps) {
  return (
    <div className="rounded-lg border border-gray-700 bg-gray-900 p-6">
      <h1 className="text-2xl font-bold text-white">{game.title}</h1>
      {game.genre && (
        <p className="text-gray-300 mt-2">
          <span className="font-medium text-white">Genre:</span> {game.genre}
        </p>
      )}
      {game.platform && (
        <p className="text-gray-300 mt-1">
          <span className="font-medium text-white">Platform:</span> {game.platform}
        </p>
      )}
    </div>
  );
}
