package movies

import (
	"project/ca/business/movies"
	"project/ca/business/users"
	"time"

	"gorm.io/gorm"
)

type GetMovieAPI struct {
	Title  string
	Year   string
	ImdbId string
	Type   string
	Poster string
	Genre  string
	Writer string
	Actors string
}

type Genres struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Movies struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Year      string
	ImdbId    string `gorm:"unique"`
	Type      string
	Poster    string
	Genre     []movies.Genre `gorm:"many2many:movie_Genre;"`
	Ratings   []users.User   `gorm:"many2many:ratings;foreignKey:Id;joinForeignKey:Movie_Id;References:Id;joinReferences:User_Id"`
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

func (movie *GetMovieAPI) ToDomainAPI() movies.GetAPI {
	return movies.GetAPI{
		Title:  movie.Title,
		Year:   movie.Year,
		ImdbId: movie.ImdbId,
		Type:   movie.Type,
		Poster: movie.Poster,
		Genre:  movie.Genre,
		Writer: movie.Writer,
		Actors: movie.Actors,
	}
}
func (movie *Movies) ToDomainMovie() movies.Movie {
	return movies.Movie{
		Id:        movie.Id,
		Title:     movie.Title,
		Year:      movie.Year,
		ImdbId:    movie.ImdbId,
		Type:      movie.Type,
		Poster:    movie.Poster,
		Genre:     movie.Genre,
		Writer:    movie.Writer,
		Actors:    movie.Actors,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: movie.UpdatedAt,
	}
}
func (genre *Genres) ToDomainGenre() movies.Genre {
	return movies.Genre{
		Id:        genre.Id,
		Name:      genre.Name,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}
}

func ToListDomain(data []Movies) (result []movies.Movie) {
	result = []movies.Movie{}
	for _, movie := range data {
		result = append(result, movie.ToDomainMovie())
	}
	return
}
func ToListDomainGenre(data []Genres) (result []movies.Genre) {
	result = []movies.Genre{}
	for _, genre := range data {
		result = append(result, genre.ToDomainGenre())
	}
	return
}

func FromDomainMovies(domain movies.Movie) Movies {
	return Movies{
		Id:        domain.Id,
		Title:     domain.Title,
		Year:      domain.Year,
		ImdbId:    domain.ImdbId,
		Type:      domain.Type,
		Poster:    domain.Poster,
		Genre:     domain.Genre,
		Writer:    domain.Writer,
		Actors:    domain.Actors,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainGenres(domain movies.Genre) Genres {
	return Genres{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
