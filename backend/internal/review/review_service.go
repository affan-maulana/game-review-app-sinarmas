package review

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/game-review/backend/internal/game"
	"github.com/google/uuid"
)

var (
	ErrGameNotFound    = errors.New("game not found")
	ErrInvalidReviewer = errors.New("reviewer name is required")
	ErrInvalidText     = errors.New("review text is required")
	ErrInvalidRating   = errors.New("rating must be between 1 and 5")
)

type Service struct {
	repo     Repository
	gameRepo game.Repository
}

func NewService(repo Repository, gameRepo game.Repository) *Service {
	return &Service{
		repo:     repo,
		gameRepo: gameRepo,
	}
}

func (s *Service) GetReviewsByGameID(ctx context.Context, gameID string) ([]Review, error) {
	_, err := s.gameRepo.GetByID(ctx, gameID)
	if err != nil {
		return nil, ErrGameNotFound
	}
	return s.repo.GetByGameID(ctx, gameID)
}

func (s *Service) CreateReview(ctx context.Context, gameID string, req CreateReviewRequest) (*Review, error) {
	if strings.TrimSpace(req.ReviewerName) == "" {
		return nil, ErrInvalidReviewer
	}

	if strings.TrimSpace(req.Text) == "" {
		return nil, ErrInvalidText
	}

	if req.Rating < 1 || req.Rating > 5 {
		return nil, ErrInvalidRating
	}

	_, err := s.gameRepo.GetByID(ctx, gameID)
	if err != nil {
		return nil, ErrGameNotFound
	}

	review := Review{
		ID:           uuid.New().String(),
		GameID:       gameID,
		ReviewerName: strings.TrimSpace(req.ReviewerName),
		Text:         strings.TrimSpace(req.Text),
		Rating:       req.Rating,
		CreatedAt:    time.Now(),
	}

	created, err := s.repo.Create(ctx, review)
	if err != nil {
		return nil, fmt.Errorf("failed to create review: %w", err)
	}

	return created, nil
}
