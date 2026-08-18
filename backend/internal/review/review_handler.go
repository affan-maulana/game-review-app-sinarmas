package review

import (
	"errors"

	"github.com/game-review/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetReviews(c *fiber.Ctx) error {
	gameID := c.Params("id")

	reviews, err := h.service.GetReviewsByGameID(c.Context(), gameID)
	if err != nil {
		if errors.Is(err, ErrGameNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Game not found")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve reviews")
	}
	return response.Success(c, fiber.StatusOK, ToReviewResponseList(reviews))
}

func (h *Handler) CreateReview(c *fiber.Ctx) error {
	gameID := c.Params("id")

	var req CreateReviewRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	review, err := h.service.CreateReview(c.Context(), gameID, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrGameNotFound):
			return response.Error(c, fiber.StatusNotFound, "Game not found")
		case errors.Is(err, ErrInvalidReviewer):
			return response.Error(c, fiber.StatusBadRequest, "Reviewer name is required")
		case errors.Is(err, ErrInvalidText):
			return response.Error(c, fiber.StatusBadRequest, "Review text is required")
		case errors.Is(err, ErrInvalidRating):
			return response.Error(c, fiber.StatusBadRequest, "Rating must be between 1 and 5")
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create review")
		}
	}
	return response.Success(c, fiber.StatusCreated, ToReviewResponse(*review))
}
