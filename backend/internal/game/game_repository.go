package game

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Game, error)
	GetByID(ctx context.Context, id string) (*Game, error)
}
