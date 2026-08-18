package container

import (
	"github.com/game-review/backend/internal/game"
	"github.com/game-review/backend/internal/review"
)

type Container struct {
	GameRepository *game.MemoryRepository
	GameService    *game.Service
	GameHandler    *game.Handler

	ReviewRepository *review.MemoryRepository
	ReviewService    *review.Service
	ReviewHandler    *review.Handler
}

func New() *Container {
	gameRepo := game.NewMemoryRepository()
	gameService := game.NewService(gameRepo)
	gameHandler := game.NewHandler(gameService)

	reviewRepo := review.NewMemoryRepository()
	reviewService := review.NewService(reviewRepo, gameRepo)
	reviewHandler := review.NewHandler(reviewService)

	return &Container{
		GameRepository:   gameRepo,
		GameService:      gameService,
		GameHandler:      gameHandler,
		ReviewRepository: reviewRepo,
		ReviewService:    reviewService,
		ReviewHandler:    reviewHandler,
	}
}
