package requests

import (
	"project/ca/business/ratings"
)

type RatingCreate struct {
	Movie_Id int     `json:"Movie_Id"`
	User_Id  int     `json:"User_Id"`
	Rate     float32 `json:"Rate"`
}

func (rate *RatingCreate) ToDomain() ratings.Rating {
	return ratings.Rating{
		Movie_Id: rate.Movie_Id,
		User_Id:  rate.User_Id,
		Rate:     rate.Rate,
	}
}
