package game_test

import (
	"context"
	"testing"

	"github.com/game-review/backend/internal/game"
)

func setupGameService() (*game.Service, *game.MemoryRepository) {
	repo := game.NewMemoryRepository()
	repo.Seed([]game.Game{
		{ID: "1", Title: "Elden Ring", Genre: "Action RPG", Platform: "PC"},
		{ID: "2", Title: "The Witcher 3", Genre: "RPG", Platform: "PC"},
	})
	service := game.NewService(repo)
	return service, repo
}

func TestGetGames(t *testing.T) {
	service, _ := setupGameService()

	games, err := service.GetGames(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(games) != 2 {
		t.Fatalf("expected 2 games, got %d", len(games))
	}
}

func TestGetGameByID(t *testing.T) {
	service, _ := setupGameService()

	game, err := service.GetGameByID(context.Background(), "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if game.Title != "Elden Ring" {
		t.Errorf("expected 'Elden Ring', got '%s'", game.Title)
	}
}

func TestGetGameByIDNotFound(t *testing.T) {
	service, _ := setupGameService()

	_, err := service.GetGameByID(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error for non-existent game, got nil")
	}
}
