package config

import (
	"time"

	"github.com/game-review/backend/internal/game"
	"github.com/game-review/backend/internal/review"
)

func SeedGames() []game.Game {
	return []game.Game{
		{ID: "1", Title: "Elden Ring", Genre: "Action RPG", Platform: "PC, PS5, Xbox"},
		{ID: "2", Title: "The Witcher 3: Wild Hunt", Genre: "RPG", Platform: "PC, PS5, Xbox, Switch"},
		{ID: "3", Title: "Cyberpunk 2077", Genre: "Action RPG", Platform: "PC, PS5, Xbox"},
		{ID: "4", Title: "God of War Ragnarök", Genre: "Action Adventure", Platform: "PS5, PC"},
		{ID: "5", Title: "Baldur's Gate 3", Genre: "RPG", Platform: "PC, PS5, Xbox"},
	}
}

func SeedReviews() []review.Review {
	return []review.Review{
		{
			ID:           "r1",
			GameID:       "1",
			ReviewerName: "Alex",
			Text:         "An absolute masterpiece. The open world is breathtaking and the combat is incredibly satisfying.",
			Rating:       5,
			CreatedAt:    time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			ID:           "r2",
			GameID:       "1",
			ReviewerName: "Sarah",
			Text:         "Challenging but rewarding. Some bosses felt unfair, but the exploration makes up for it.",
			Rating:       4,
			CreatedAt:    time.Date(2024, 2, 20, 14, 30, 0, 0, time.UTC),
		},
		{
			ID:           "r3",
			GameID:       "2",
			ReviewerName: "Mike",
			Text:         "One of the best RPGs ever made. The story, characters, and world-building are top-notch.",
			Rating:       5,
			CreatedAt:    time.Date(2024, 1, 10, 9, 0, 0, 0, time.UTC),
		},
		{
			ID:           "r4",
			GameID:       "3",
			ReviewerName: "Jordan",
			Text:         "After all the patches, this game finally lives up to its potential. Night City is stunning.",
			Rating:       4,
			CreatedAt:    time.Date(2024, 3, 5, 16, 0, 0, 0, time.UTC),
		},
		{
			ID:           "r5",
			GameID:       "3",
			ReviewerName: "Taylor",
			Text:         "Great story and characters, but the gameplay still has some rough edges.",
			Rating:       3,
			CreatedAt:    time.Date(2024, 3, 10, 11, 0, 0, 0, time.UTC),
		},
		{
			ID:           "r6",
			GameID:       "4",
			ReviewerName: "Chris",
			Text:         "An emotional journey with incredible combat. A worthy sequel.",
			Rating:       5,
			CreatedAt:    time.Date(2024, 2, 1, 13, 0, 0, 0, time.UTC),
		},
		{
			ID:           "r7",
			GameID:       "5",
			ReviewerName: "Pat",
			Text:         "The gold standard for CRPGs. Every choice matters and the replayability is insane.",
			Rating:       5,
			CreatedAt:    time.Date(2024, 4, 1, 10, 0, 0, 0, time.UTC),
		},
	}
}
