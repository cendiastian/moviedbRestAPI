package requests

import "project/ca/business/ratings"

type RatingDelete struct {
	Movie_Id int `json:"Movie_Id"`
	User_Id  int `json:"User_Id"`
}

func (rate *RatingDelete) ToDomain() ratings.Rating {
	return ratings.Rating{
		Movie_Id: rate.Movie_Id,
		User_Id:  rate.User_Id,
	}
}
