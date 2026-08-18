package server

import (
	"github.com/game-review/backend/internal/container"
	"github.com/game-review/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, c *container.Container) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return response.Success(ctx, fiber.StatusOK, fiber.Map{
			"status": "ok",
		})
	})

	api := app.Group("/api")

	games := api.Group("/games")
	games.Get("/", c.GameHandler.GetGames)
	games.Get("/:id", c.GameHandler.GetGameByID)
	games.Get("/:id/reviews", c.ReviewHandler.GetReviews)
	games.Post("/:id/reviews", c.ReviewHandler.CreateReview)
}
