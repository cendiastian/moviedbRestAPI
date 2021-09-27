package responses

import (
	"project/ca/business/ratings"
	"time"
)

type RatingResponse struct {
	Movie_Id  int       `json:"Movie_Id"`
	User_Id   int       `json:"User_Id"`
	Rate      float32   `json:"Rate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain ratings.Rating) RatingResponse {
	return RatingResponse{
		Movie_Id:  domain.Movie_Id,
		User_Id:   domain.User_Id,
		Rate:      domain.Rate,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(domain []ratings.Rating) (response []RatingResponse) {
	for _, rate := range domain {
		response = append(response, FromDomain(rate))
	}
	return
}
