package game

type GameResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Platform string `json:"platform"`
}

func ToGameResponse(g Game) GameResponse {
	return GameResponse{
		ID:       g.ID,
		Title:    g.Title,
		Genre:    g.Genre,
		Platform: g.Platform,
	}
}

func ToGameResponseList(games []Game) []GameResponse {
	responses := make([]GameResponse, len(games))
	for i, g := range games {
		responses[i] = ToGameResponse(g)
	}
	return responses
}
