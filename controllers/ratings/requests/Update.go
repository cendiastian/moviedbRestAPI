package requests

import "project/business/ratings"

type RatingUpdate struct {
	MovieId int     `json:"MovieId"`
	UserId  int     `json:"UserId"`
	Rate    float32 `json:"Rate"`
}

func (rate *RatingUpdate) ToDomain() ratings.Ratings {
	return ratings.Ratings{
		MovieId: rate.MovieId,
		UserId:  rate.UserId,
		Rate:    rate.Rate,
	}
}
