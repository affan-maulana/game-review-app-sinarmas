import Link from "next/link";
import { Game } from "@/types/game";

interface GameCardProps {
  game: Game;
}

export function GameCard({ game }: GameCardProps) {
  return (
    <Link
      href={`/games/${game.id}`}
      className="block rounded-lg border border-gray-200 p-4 hover:border-blue-500 hover:shadow-md transition-all"
    >
      <h3 className="text-lg font-semibold text-gray-900">{game.title}</h3>
      {game.genre && (
        <p className="text-sm text-gray-600 mt-1">{game.genre}</p>
      )}
      {game.platform && (
        <p className="text-xs text-gray-400 mt-2">{game.platform}</p>
      )}
    </Link>
  );
}
