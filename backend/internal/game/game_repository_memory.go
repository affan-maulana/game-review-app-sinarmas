package game

import (
	"context"
	"fmt"
	"sync"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	games []Game
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		games: make([]Game, 0),
	}
}

func (r *MemoryRepository) Seed(games []Game) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.games = games
}

func (r *MemoryRepository) GetAll(_ context.Context) ([]Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Game, len(r.games))
	copy(result, r.games)
	return result, nil
}

func (r *MemoryRepository) GetByID(_ context.Context, id string) (*Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, g := range r.games {
		if g.ID == id {
			game := g
			return &game, nil
		}
	}
	return nil, fmt.Errorf("game not found: %s", id)
}
