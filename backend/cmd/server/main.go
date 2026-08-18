package main

import (
	"log"

	"github.com/game-review/backend/internal/config"
	"github.com/game-review/backend/internal/container"
	"github.com/game-review/backend/internal/middleware"
	"github.com/game-review/backend/internal/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	c := container.New()

	c.GameRepository.Seed(config.SeedGames())
	c.ReviewRepository.Seed(config.SeedReviews())

	app := fiber.New(fiber.Config{
		AppName: "Game Review API",
	})

	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.RecoverMiddleware())
	app.Use(middleware.CORSMiddleware(cfg.FrontendURL))

	server.RegisterRoutes(app, c)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
