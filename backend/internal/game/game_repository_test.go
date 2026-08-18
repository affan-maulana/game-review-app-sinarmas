package game_test

import (
	"context"
	"testing"

	"github.com/game-review/backend/internal/game"
)

func TestMemoryRepositoryGetAll(t *testing.T) {
	repo := game.NewMemoryRepository()
	repo.Seed([]game.Game{
		{ID: "1", Title: "Game 1"},
		{ID: "2", Title: "Game 2"},
	})

	games, err := repo.GetAll(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(games) != 2 {
		t.Fatalf("expected 2 games, got %d", len(games))
	}
}

func TestMemoryRepositoryGetByID(t *testing.T) {
	repo := game.NewMemoryRepository()
	repo.Seed([]game.Game{
		{ID: "1", Title: "Game 1"},
	})

	game, err := repo.GetByID(context.Background(), "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if game.Title != "Game 1" {
		t.Errorf("expected 'Game 1', got '%s'", game.Title)
	}
}

func TestMemoryRepositoryGetByIDNotFound(t *testing.T) {
	repo := game.NewMemoryRepository()

	_, err := repo.GetByID(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error for non-existent game, got nil")
	}
}
