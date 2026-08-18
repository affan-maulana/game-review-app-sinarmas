package review

import "context"

type Repository interface {
	GetByGameID(ctx context.Context, gameID string) ([]Review, error)
	Create(ctx context.Context, review Review) (*Review, error)
}
