package genres

import (
	"context"
	"project/ca/business/genres"

	"gorm.io/gorm"
)

type MysqlGenreRepository struct {
	Connect *gorm.DB
}

func NewMysqlGenreRepository(connect *gorm.DB) genres.Repository {
	return &MysqlGenreRepository{
		Connect: connect,
	}
}

func (rep *MysqlGenreRepository) FirstOrCreate(ctx context.Context, name string) (genres.Genre, error) {
	var genre Genres
	genre.Name = name
	result := rep.Connect.FirstOrCreate(&genre, "Name= ?", name)

	if result.Error != nil {
		return genres.Genre{}, result.Error
	}

	return genre.ToDomainGenre(), nil
}

func (rep *MysqlGenreRepository) GetAllGenre(ctx context.Context) ([]genres.Genre, error) {
	var genre []Genres
	result := rep.Connect.Find(&genre)
	if result.Error != nil {
		return []genres.Genre{}, result.Error
	}
	return ToListDomainGenre(genre), nil
}
