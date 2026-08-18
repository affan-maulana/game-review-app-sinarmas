package review

import "time"

type Review struct {
	ID           string    `json:"id"`
	GameID       string    `json:"gameId"`
	ReviewerName string    `json:"reviewerName"`
	Text         string    `json:"text"`
	Rating       int       `json:"rating"`
	CreatedAt    time.Time `json:"createdAt"`
}
