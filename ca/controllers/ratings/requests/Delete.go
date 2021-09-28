package requests

import "project/ca/business/ratings"

type RatingDelete struct {
	MovieId int `json:"MovieId"`
	UserId  int `json:"UserId"`
}

func (rate *RatingDelete) ToDomain() ratings.Ratings {
	return ratings.Ratings{
		MovieId: rate.MovieId,
		UserId:  rate.UserId,
	}
}
