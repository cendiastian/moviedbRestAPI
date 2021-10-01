package responses

import (
	"project/business/ratings"
	_user "project/controllers/users/responses"
	"time"
)

type RatingResponse struct {
	// MovieId int `json:"movie"`
	// UserId    int       `json:"UserId"`
	User      _user.UserResponse `json:"user"`
	Rate      float32            `json:"rating"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func FromDomain(domain ratings.Ratings) RatingResponse {
	return RatingResponse{
		// MovieId: domain.MovieId,
		// UserId:    domain.UserId,
		User:      _user.FromDomain(domain.User),
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
