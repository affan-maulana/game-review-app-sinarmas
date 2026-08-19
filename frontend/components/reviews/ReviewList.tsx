"use client";

import { Review } from "@/types/review";
import { ReviewCard } from "./ReviewCard";

interface ReviewListProps {
  reviews: Review[];
  loading: boolean;
  error: string | null;
}

export function ReviewList({ reviews, loading, error }: ReviewListProps) {
  if (loading) {
    return <p className="text-gray-300">Loading reviews...</p>;
  }

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (reviews.length === 0) {
    return (
      <p className="text-gray-300">
        No reviews yet. Be the first to review!
      </p>
    );
  }

  return (
    <div className="space-y-4">
      {reviews.map((review) => (
        <ReviewCard key={review.id} review={review} />
      ))}
    </div>
  );
}
