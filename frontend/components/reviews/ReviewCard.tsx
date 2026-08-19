import { Review } from "@/types/review";

interface ReviewCardProps {
  review: Review;
}

export function ReviewCard({ review }: ReviewCardProps) {
  return (
    <div className="rounded-lg border border-gray-700 bg-gray-900 p-4">
      <div className="flex items-center justify-between">
        <span className="font-medium text-white">
          {review.reviewerName}
        </span>
        <span className="text-sm text-yellow-400 font-semibold">
          {"★".repeat(review.rating)}
          {"☆".repeat(5 - review.rating)}
        </span>
      </div>
      <p className="text-gray-300 mt-2">{review.text}</p>
      <p className="text-xs text-gray-500 mt-2">
        {new Date(review.createdAt).toLocaleDateString()}
      </p>
    </div>
  );
}
