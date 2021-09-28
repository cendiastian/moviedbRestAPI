package movies

import (
	"context"
	"project/ca/business/genres"
	"project/ca/business/ratings"
	"time"
)

type Movie struct {
	Id        int
	Title     string
	Year      string
	ImdbId    string
	Type      string
	Poster    string
	Genre     []genres.Genre
	Ratings   []ratings.Ratings
	Rating    float32
	Writer    string
	Actors    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAllMovie(ctx context.Context) ([]Movie, error)
	CreateMovie(ctx context.Context, ImdbId string) (Movie, error)
	MovieDetail(ctx context.Context, id int) (Movie, error)
	SearchMovie(ctx context.Context, title string) ([]Movie, error)
	FilterGenre(ctx context.Context, genre string) ([]Movie, error)
	FilterOrder(ctx context.Context, order string) ([]Movie, error)
	DeleteAll(ctx context.Context) error
	DeleteMovie(ctx context.Context, id int) (Movie, error)
	UpdateMovie(ctx context.Context, Movie Movie) error
}

type Repository interface {
	GetAllMovie(ctx context.Context) ([]Movie, error)
	DeleteAll(ctx context.Context) error
	UpdateMovie(ctx context.Context, Movie Movie) error
	DeleteMovie(ctx context.Context, id int) (Movie, error)
	CreateMovie(ctx context.Context, Movie Movie, array []genres.Genre) (Movie, error)
	MovieDetail(ctx context.Context, id int) (Movie, error)
	SearchMovie(ctx context.Context, title string) ([]Movie, error)
	FilterGenre(ctx context.Context, genre string) ([]Movie, error)
	FilterOrder(ctx context.Context, order string) ([]Movie, error)
	// CreateGenreAPI(ctx context.Context, name string) (Genre, error)
	// GetMovieAPI(ctx context.Context, Movie Movie) (Movie, error)
}
