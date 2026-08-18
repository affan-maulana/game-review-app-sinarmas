import { apiClient } from "@/lib/api/client";
import { routes } from "@/lib/api/routes";
import { CreateReviewPayload, Review } from "@/types/review";

export const reviewService = {
  getReviews: (gameId: string) =>
    apiClient.get<Review[]>(routes.games.reviews(gameId)),

  createReview: (gameId: string, payload: CreateReviewPayload) =>
    apiClient.post<Review>(routes.games.reviews(gameId), payload),
};
