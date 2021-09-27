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

func (rep *MysqlRatingRepository) Delete(ctx context.Context, Movie_Id int, User_Id int) error {
	var Rating Ratings
	result := rep.Connect.Delete(&Rating, "movie_id = ? AND user_id", Movie_Id, User_Id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlRatingRepository) Update(ctx context.Context, Movie_Id int, User_Id int, Rate float32) error {
	result := rep.Connect.Where("movie_id = ? AND user_id", Movie_Id, User_Id).Updates(&Ratings{Rate: Rate})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlRatingRepository) Create(ctx context.Context, Movie_Id int, User_Id int, Rate float32) (ratings.Rating, error) {
	Rating := Ratings{
		Movie_Id: Movie_Id,
		User_Id:  User_Id,
		Rate:     Rate,
	}
	result := rep.Connect.Create(&Rating)

	if result.Error != nil {
		return ratings.Rating{}, result.Error
	}

	return Rating.ToDomain(), nil
}
