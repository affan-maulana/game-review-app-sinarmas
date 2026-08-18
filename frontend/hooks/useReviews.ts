"use client";

import { useState, useEffect, useCallback } from "react";
import { Review, CreateReviewPayload } from "@/types/review";
import { reviewService } from "@/services/review.service";

export function useReviews(gameId: string) {
  const [reviews, setReviews] = useState<Review[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState<string | null>(null);

  const fetchReviews = useCallback(async () => {
    try {
      setLoading(true);
      const data = await reviewService.getReviews(gameId);
      setReviews(data);
      setError(null);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Failed to load reviews"
      );
    } finally {
      setLoading(false);
    }
  }, [gameId]);

  useEffect(() => {
    fetchReviews();
  }, [fetchReviews]);

  const createReview = async (payload: CreateReviewPayload) => {
    try {
      setIsSubmitting(true);
      setSubmitError(null);
      const newReview = await reviewService.createReview(gameId, payload);
      setReviews((prev) => [...prev, newReview]);
      return newReview;
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Failed to submit review";
      setSubmitError(message);
      throw err;
    } finally {
      setIsSubmitting(false);
    }
  };

  return {
    reviews,
    loading,
    error,
    createReview,
    isSubmitting,
    submitError,
  };
}
