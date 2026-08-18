package review

type CreateReviewRequest struct {
	ReviewerName string `json:"reviewerName"`
	Text         string `json:"text"`
	Rating       int    `json:"rating"`
}

type ReviewResponse struct {
	ID           string `json:"id"`
	GameID       string `json:"gameId"`
	ReviewerName string `json:"reviewerName"`
	Text         string `json:"text"`
	Rating       int    `json:"rating"`
	CreatedAt    string `json:"createdAt"`
}

func ToReviewResponse(r Review) ReviewResponse {
	return ReviewResponse{
		ID:           r.ID,
		GameID:       r.GameID,
		ReviewerName: r.ReviewerName,
		Text:         r.Text,
		Rating:       r.Rating,
		CreatedAt:    r.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToReviewResponseList(reviews []Review) []ReviewResponse {
	responses := make([]ReviewResponse, len(reviews))
	for i, r := range reviews {
		responses[i] = ToReviewResponse(r)
	}
	return responses
}
