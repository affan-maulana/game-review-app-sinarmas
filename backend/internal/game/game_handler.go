package game

import (
	"github.com/game-review/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetGames(c *fiber.Ctx) error {
	games, err := h.service.GetGames(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve games")
	}
	return response.Success(c, fiber.StatusOK, ToGameResponseList(games))
}

func (h *Handler) GetGameByID(c *fiber.Ctx) error {
	id := c.Params("id")

	game, err := h.service.GetGameByID(c.Context(), id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Game not found")
	}
	return response.Success(c, fiber.StatusOK, ToGameResponse(*game))
}
