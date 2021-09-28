package genres

import (
	"project/ca/business/genres"
	"time"

	"gorm.io/gorm"
)

type Genres struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (genre *Genres) ToDomainGenre() genres.Genre {
	return genres.Genre{
		Id:        genre.Id,
		Name:      genre.Name,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}
}

func ToListDomainGenre(data []Genres) (result []genres.Genre) {
	result = []genres.Genre{}
	for _, genre := range data {
		result = append(result, genre.ToDomainGenre())
	}
	return result
}
func FromListDomainGenre(data []genres.Genre) (result []Genres) {
	result = []Genres{}
	for _, genre := range data {
		result = append(result, FromDomain(genre))
	}
	return result
}

func FromDomain(domain genres.Genre) Genres {
	return Genres{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
