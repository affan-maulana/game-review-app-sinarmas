"use client";

import { useState, FormEvent } from "react";
import { CreateReviewPayload } from "@/types/review";

interface ReviewFormProps {
  onSubmit: (payload: CreateReviewPayload) => Promise<unknown>;
  isSubmitting: boolean;
  submitError: string | null;
}

export function ReviewForm({
  onSubmit,
  isSubmitting,
  submitError,
}: ReviewFormProps) {
  const [reviewerName, setReviewerName] = useState("");
  const [text, setText] = useState("");
  const [rating, setRating] = useState(5);
  const [validationErrors, setValidationErrors] = useState<string[]>([]);

  const validate = (): boolean => {
    const errors: string[] = [];

    if (!reviewerName.trim()) {
      errors.push("Reviewer name is required");
    }
    if (!text.trim()) {
      errors.push("Review text is required");
    }
    if (rating < 1 || rating > 5) {
      errors.push("Rating must be between 1 and 5");
    }

    setValidationErrors(errors);
    return errors.length === 0;
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (!validate()) return;

    try {
      await onSubmit({
        reviewerName: reviewerName.trim(),
        text: text.trim(),
        rating,
      });
      setReviewerName("");
      setText("");
      setRating(5);
      setValidationErrors([]);
    } catch {
      // Error is handled by the parent hook
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <h3 className="text-lg font-semibold text-gray-900">Write a Review</h3>

      {validationErrors.length > 0 && (
        <div className="bg-red-50 border border-red-200 rounded p-3">
          {validationErrors.map((err, i) => (
            <p key={i} className="text-red-600 text-sm">
              {err}
            </p>
          ))}
        </div>
      )}

      {submitError && (
        <div className="bg-red-50 border border-red-200 rounded p-3">
          <p className="text-red-600 text-sm">{submitError}</p>
        </div>
      )}

      <div>
        <label
          htmlFor="reviewerName"
          className="block text-sm font-medium text-gray-700"
        >
          Your Name
        </label>
        <input
          id="reviewerName"
          type="text"
          value={reviewerName}
          onChange={(e) => setReviewerName(e.target.value)}
          className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
          placeholder="Enter your name"
        />
      </div>

      <div>
        <label
          htmlFor="rating"
          className="block text-sm font-medium text-gray-700"
        >
          Rating
        </label>
        <select
          id="rating"
          value={rating}
          onChange={(e) => setRating(Number(e.target.value))}
          className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
        >
          {[5, 4, 3, 2, 1].map((r) => (
            <option key={r} value={r}>
              {r} - {"★".repeat(r)}
            </option>
          ))}
        </select>
      </div>

      <div>
        <label
          htmlFor="reviewText"
          className="block text-sm font-medium text-gray-700"
        >
          Review
        </label>
        <textarea
          id="reviewText"
          value={text}
          onChange={(e) => setText(e.target.value)}
          rows={4}
          className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 text-gray-900 shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
          placeholder="Write your review..."
        />
      </div>

      <button
        type="submit"
        disabled={isSubmitting}
        className="w-full rounded-md bg-blue-600 px-4 py-2 text-white font-medium hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {isSubmitting ? "Submitting..." : "Submit Review"}
      </button>
    </form>
  );
}
