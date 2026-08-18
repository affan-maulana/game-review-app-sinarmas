package review_test

import (
	"context"
	"testing"
	"time"

	"github.com/game-review/backend/internal/review"
)

func TestMemoryRepositoryCreate(t *testing.T) {
	repo := review.NewMemoryRepository()

	r := review.Review{
		ID:           "1",
		GameID:       "game1",
		ReviewerName: "John",
		Text:         "Great!",
		Rating:       5,
		CreatedAt:    time.Now(),
	}

	created, err := repo.Create(context.Background(), r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if created.ID != "1" {
		t.Errorf("expected ID '1', got '%s'", created.ID)
	}
}

func TestMemoryRepositoryGetByGameID(t *testing.T) {
	repo := review.NewMemoryRepository()
	repo.Seed([]review.Review{
		{ID: "1", GameID: "game1", ReviewerName: "John", Text: "Good", Rating: 4, CreatedAt: time.Now()},
		{ID: "2", GameID: "game1", ReviewerName: "Jane", Text: "Great", Rating: 5, CreatedAt: time.Now()},
		{ID: "3", GameID: "game2", ReviewerName: "Bob", Text: "OK", Rating: 3, CreatedAt: time.Now()},
	})

	reviews, err := repo.GetByGameID(context.Background(), "game1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(reviews) != 2 {
		t.Fatalf("expected 2 reviews, got %d", len(reviews))
	}
}

func TestMemoryRepositoryGetByGameIDEmpty(t *testing.T) {
	repo := review.NewMemoryRepository()

	reviews, err := repo.GetByGameID(context.Background(), "nonexistent")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(reviews) != 0 {
		t.Fatalf("expected 0 reviews, got %d", len(reviews))
	}
}
