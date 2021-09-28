package ratings

import (
	"context"
	"project/ca/business/ratings"

	"gorm.io/gorm"
)

type MysqlRatingRepository struct {
	Connect *gorm.DB
}

func NewMysqlRatingRepository(connect *gorm.DB) ratings.Repository {
	return &MysqlRatingRepository{
		Connect: connect,
	}
}

func (rep *MysqlRatingRepository) Delete(ctx context.Context, MovieId int, UserId int) error {
	var Rating Ratings
	result := rep.Connect.Delete(&Rating, "movieid = ? AND userid", MovieId, UserId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlRatingRepository) Update(ctx context.Context, domain ratings.Ratings) error {
	Rating := FromDomain(domain)
	result := rep.Connect.Where("movieid = ? AND userid", Rating.MovieId, Rating.UserId).Updates(&Ratings{Rate: Rating.Rate})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlRatingRepository) Create(ctx context.Context, domain ratings.Ratings) (ratings.Ratings, error) {
	Rating := FromDomain(domain)
	result := rep.Connect.Create(&Rating)

	if result.Error != nil {
		return ratings.Ratings{}, result.Error
	}

	return Rating.ToDomain(), nil
}

func (rep *MysqlRatingRepository) Detail(ctx context.Context, movie int, user int) (ratings.Ratings, error) {
	var pay Ratings
	result := rep.Connect.Preload("MovieId").First(&pay, "MovieId= ? AND UserId  = ? ", movie, user)
	if result.Error != nil {
		return ratings.Ratings{}, result.Error
	}
	return pay.ToDomain(), nil
}
