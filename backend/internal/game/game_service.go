package game

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetGames(ctx context.Context) ([]Game, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetGameByID(ctx context.Context, id string) (*Game, error) {
	return s.repo.GetByID(ctx, id)
}
