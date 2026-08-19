import Link from "next/link";
import { Game } from "@/types/game";

interface GameCardProps {
  game: Game;
}

export function GameCard({ game }: GameCardProps) {
  return (
    <Link
      href={`/games/${game.id}`}
      className="block rounded-lg border border-gray-700 bg-gray-800 p-4 hover:border-blue-500 hover:shadow-md transition-all"
    >
      <h3 className="text-lg font-semibold text-white">{game.title}</h3>
      {game.genre && <p className="text-sm mt-1 text-gray-300">{game.genre}</p>}
      {game.platform && <p className="text-xs mt-2 text-gray-400">{game.platform}</p>}
    </Link>
  );
}
