package movie

type GetAllMovie struct {
	Search      []CreateMovie `json:"Search"`
	TotalResult string        `json:"totalResults"`
}