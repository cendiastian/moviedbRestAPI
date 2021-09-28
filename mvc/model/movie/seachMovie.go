package movie

type SearchMovie struct {
	Result      []Movie `json:"search"`
	TotalResult string  `json:"totalResults"`
	Response    string  `json:"response"`
}

type SearchRequest struct {
	MovieTitle string
}

type SearchResponse struct {
	Code     int     `json:"code"`
	Message  string  `json:"message"`
	Response []Movie `json:"Search"`
}
