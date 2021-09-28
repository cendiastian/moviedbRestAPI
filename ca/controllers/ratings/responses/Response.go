package responses

import (
	"project/ca/business/ratings"
	"time"
)

type RatingResponse struct {
	MovieId   int       `json:"MovieId"`
	Username  string    `json:"Username"`
	UserId    int       `json:"UserId"`
	Rate      float32   `json:"Rate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain ratings.Ratings) RatingResponse {
	return RatingResponse{
		MovieId:   domain.MovieId,
		Username:  domain.Username,
		Rate:      domain.Rate,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(domain []ratings.Ratings) (response []RatingResponse) {
	for _, rate := range domain {
		response = append(response, FromDomain(rate))
	}
	return
}
