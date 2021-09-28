package movies

import (
	"project/ca/business/movies"
	"project/ca/drivers/databases/genres"
	"project/ca/drivers/databases/ratings"
	"time"

	"gorm.io/gorm"
)

type Movies struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Year      string
	ImdbId    string `gorm:"unique"`
	Type      string
	Poster    string
	Genre     []genres.Genres   `gorm:"many2many:movie_Genre;"`
	Ratings   []ratings.Ratings `gorm:"foreignKey:MovieId"`
	Rating    float32
	Writer    string
	Actors    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type Movies struct {
// 	Id        int `gorm:"primaryKey"`
// 	Title     string
// 	Year      string
// 	ImdbId    string `gorm:"unique"`
// 	Type      string
// 	Poster    string
// 	Genre     []movies.Genre `gorm:"many2many:movie_Genre;" json:"Genre"`
// 	Writer    string
// 	Actors    string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

func (movie *Movies) ToDomainMovie() movies.Movie {
	return movies.Movie{
		Id:        movie.Id,
		Title:     movie.Title,
		Year:      movie.Year,
		ImdbId:    movie.ImdbId,
		Type:      movie.Type,
		Poster:    movie.Poster,
		Genre:     genres.ToListDomainGenre(movie.Genre),
		Ratings:   ratings.ToListDomain(movie.Ratings),
		Rating:    movie.Rating,
		Writer:    movie.Writer,
		Actors:    movie.Actors,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: movie.UpdatedAt,
	}
}

func ToListDomain(data []Movies) (result []movies.Movie) {
	result = []movies.Movie{}
	for _, movie := range data {
		result = append(result, movie.ToDomainMovie())
	}
	return
}

func FromDomain(domain movies.Movie) Movies {
	return Movies{
		Id:     domain.Id,
		Title:  domain.Title,
		Year:   domain.Year,
		ImdbId: domain.ImdbId,
		Type:   domain.Type,
		Poster: domain.Poster,
		Genre:  genres.FromListDomainGenre(domain.Genre),
		// Ratings:   domain.Ratings,
		Rating:    domain.Rating,
		Writer:    domain.Writer,
		Actors:    domain.Actors,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
