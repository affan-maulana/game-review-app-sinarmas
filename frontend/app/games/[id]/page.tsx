"use client";

import { useParams } from "next/navigation";
import Link from "next/link";
import { useGame } from "@/hooks/useGame";
import { useReviews } from "@/hooks/useReviews";
import { GameDetails } from "@/components/games/GameDetails";
import { ReviewList } from "@/components/reviews/ReviewList";
import { ReviewForm } from "@/components/reviews/ReviewForm";

export default function GameDetailPage() {
  const params = useParams();
  const id = params.id as string;

  const { game, loading: gameLoading, error: gameError } = useGame(id);
  const {
    reviews,
    loading: reviewsLoading,
    error: reviewsError,
    createReview,
    isSubmitting,
    submitError,
  } = useReviews(id);

  if (gameLoading) {
    return <p className="text-gray-300">Loading game details...</p>;
  }

  if (gameError || !game) {
    return (
      <div>
        <p className="text-red-500 mb-4">
          Error: {gameError || "Game not found"}
        </p>
        <Link href="/" className="text-blue-400 hover:underline">
          Back to games
        </Link>
      </div>
    );
  }

  return (
    <div className="space-y-8">
      <div>
        <Link
          href="/"
          className="text-blue-400 hover:underline text-sm mb-4 inline-block"
        >
          &larr; Back to games
        </Link>
        <GameDetails game={game} />
      </div>

      <div>
        <h2 className="text-xl font-bold text-white mb-4">Reviews</h2>
        <ReviewList
          reviews={reviews}
          loading={reviewsLoading}
          error={reviewsError}
        />
      </div>

      <div className="border-t border-gray-700 pt-6">
        <ReviewForm
          onSubmit={createReview}
          isSubmitting={isSubmitting}
          submitError={submitError}
        />
      </div>
    </div>
  );
}
