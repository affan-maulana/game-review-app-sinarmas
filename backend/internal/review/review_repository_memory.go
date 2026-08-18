package review

import (
	"context"
	"sync"
)

type MemoryRepository struct {
	mu      sync.RWMutex
	reviews []Review
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		reviews: make([]Review, 0),
	}
}

func (r *MemoryRepository) Seed(reviews []Review) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reviews = reviews
}

func (r *MemoryRepository) GetByGameID(_ context.Context, gameID string) ([]Review, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []Review
	for _, rev := range r.reviews {
		if rev.GameID == gameID {
			result = append(result, rev)
		}
	}
	if result == nil {
		result = make([]Review, 0)
	}
	return result, nil
}

func (r *MemoryRepository) Create(_ context.Context, review Review) (*Review, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.reviews = append(r.reviews, review)
	return &review, nil
}
