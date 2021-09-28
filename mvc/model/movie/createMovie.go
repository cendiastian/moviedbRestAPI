package movie

type CreateMovie struct {
	Title         string `json:"Title"`
	Year          string `json:"Year"`
	Genre         string `json:"Genre"`
	Director      string `json:"Director"`
	Writer        string `json:"Writer"`
	Actors        string `json:"Actors"`
	Plot          string `json:"Plot"`
	Poster        string `json:"Poster"`
	ImdbId        string `json:"imdbId" gorm:"unique"`
	Type          string `json:"Type"`
	GetAllMovieID int    `json:"GetAllMovieID"`
}
