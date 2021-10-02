package ratings

import (
	"project/business/ratings"
	"project/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Ratings struct {
	MovieId int `gorm:"primaryKey"`
	// Movie     movies.Movies
	UserId    int `gorm:"primaryKey"`
	User      users.Users
	Rate      float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (rate *Ratings) ToDomain() ratings.Ratings {
	return ratings.Ratings{
		MovieId:   rate.MovieId,
		UserId:    rate.UserId,
		User:      rate.User.ToDomainUser(),
		Rate:      rate.Rate,
		CreatedAt: rate.CreatedAt,
		UpdatedAt: rate.UpdatedAt,
	}
}

func ToListDomain(data []Ratings) (result []ratings.Ratings) {
	result = []ratings.Ratings{}
	for _, rate := range data {
		result = append(result, rate.ToDomain())
	}
	return
}

func FromDomain(domain ratings.Ratings) Ratings {
	return Ratings{
		MovieId:   domain.MovieId,
		UserId:    domain.UserId,
		Rate:      domain.Rate,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
