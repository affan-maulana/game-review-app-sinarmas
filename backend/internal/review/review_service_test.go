package review_test

import (
	"context"
	"testing"

	"github.com/game-review/backend/internal/game"
	"github.com/game-review/backend/internal/review"
)

func setupReviewService() (*review.Service, *game.MemoryRepository) {
	gameRepo := game.NewMemoryRepository()
	gameRepo.Seed([]game.Game{
		{ID: "1", Title: "Elden Ring", Genre: "Action RPG", Platform: "PC"},
	})

	reviewRepo := review.NewMemoryRepository()
	service := review.NewService(reviewRepo, gameRepo)
	return service, gameRepo
}

func TestCreateValidReview(t *testing.T) {
	service, _ := setupReviewService()

	req := review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "Great game!",
		Rating:       5,
	}

	result, err := service.CreateReview(context.Background(), "1", req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ReviewerName != "John" {
		t.Errorf("expected 'John', got '%s'", result.ReviewerName)
	}
	if result.Rating != 5 {
		t.Errorf("expected rating 5, got %d", result.Rating)
	}
	if result.GameID != "1" {
		t.Errorf("expected gameID '1', got '%s'", result.GameID)
	}
}

func TestCreateReviewEmptyReviewerName(t *testing.T) {
	service, _ := setupReviewService()

	req := review.CreateReviewRequest{
		ReviewerName: "",
		Text:         "Great game!",
		Rating:       5,
	}

	_, err := service.CreateReview(context.Background(), "1", req)
	if err != review.ErrInvalidReviewer {
		t.Fatalf("expected ErrInvalidReviewer, got %v", err)
	}
}

func TestCreateReviewEmptyText(t *testing.T) {
	service, _ := setupReviewService()

	req := review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "",
		Rating:       5,
	}

	_, err := service.CreateReview(context.Background(), "1", req)
	if err != review.ErrInvalidText {
		t.Fatalf("expected ErrInvalidText, got %v", err)
	}
}

func TestCreateReviewInvalidRating(t *testing.T) {
	service, _ := setupReviewService()

	tests := []struct {
		name   string
		rating int
	}{
		{"rating too low", 0},
		{"rating too high", 6},
		{"negative rating", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := review.CreateReviewRequest{
				ReviewerName: "John",
				Text:         "Great game!",
				Rating:       tt.rating,
			}

			_, err := service.CreateReview(context.Background(), "1", req)
			if err != review.ErrInvalidRating {
				t.Fatalf("expected ErrInvalidRating, got %v", err)
			}
		})
	}
}

func TestCreateReviewNonExistentGame(t *testing.T) {
	service, _ := setupReviewService()

	req := review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "Great game!",
		Rating:       5,
	}

	_, err := service.CreateReview(context.Background(), "nonexistent", req)
	if err != review.ErrGameNotFound {
		t.Fatalf("expected ErrGameNotFound, got %v", err)
	}
}

func TestGetReviewsByGameID(t *testing.T) {
	service, _ := setupReviewService()

	req := review.CreateReviewRequest{
		ReviewerName: "John",
		Text:         "Great game!",
		Rating:       5,
	}

	_, err := service.CreateReview(context.Background(), "1", req)
	if err != nil {
		t.Fatalf("failed to create review: %v", err)
	}

	reviews, err := service.GetReviewsByGameID(context.Background(), "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(reviews) != 1 {
		t.Fatalf("expected 1 review, got %d", len(reviews))
	}
}

func TestGetReviewsNonExistentGame(t *testing.T) {
	service, _ := setupReviewService()

	_, err := service.GetReviewsByGameID(context.Background(), "nonexistent")
	if err != review.ErrGameNotFound {
		t.Fatalf("expected ErrGameNotFound, got %v", err)
	}
}
