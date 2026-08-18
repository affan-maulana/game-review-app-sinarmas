package game

type Game struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Platform string `json:"platform"`
}
