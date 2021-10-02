package requests

import (
	"project/business/ratings"
)

type RatingCreate struct {
	MovieId int     `json:"MovieId"`
	UserId  int     `json:"UserId"`
	Rate    float32 `json:"Rate"`
}

func (rate *RatingCreate) ToDomain() ratings.Ratings {
	return ratings.Ratings{
		MovieId: rate.MovieId,
		UserId:  rate.UserId,
		Rate:    rate.Rate,
	}
}
