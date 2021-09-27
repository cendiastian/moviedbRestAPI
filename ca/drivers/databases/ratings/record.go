package ratings

import (
	"project/ca/business/ratings"
	"time"

	"gorm.io/gorm"
)

type Ratings struct {
	Movie_Id  int `gorm:"primaryKey"`
	User_Id   int `gorm:"primaryKey"`
	Rate      float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (rate *Ratings) ToDomain() ratings.Rating {
	return ratings.Rating{
		Movie_Id:  rate.Movie_Id,
		User_Id:   rate.User_Id,
		Rate:      rate.Rate,
		CreatedAt: rate.CreatedAt,
		UpdatedAt: rate.UpdatedAt,
	}
}

func ToListDomain(data []Ratings) (result []ratings.Rating) {
	result = []ratings.Rating{}
	for _, rate := range data {
		result = append(result, rate.ToDomain())
	}
	return
}

func FromDomain(domain ratings.Rating) Ratings {
	return Ratings{
		Movie_Id:  domain.Movie_Id,
		User_Id:   domain.User_Id,
		Rate:      domain.Rate,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
