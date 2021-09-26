package movies

import (
	"context"
	"time"
)

type Movie struct {
	Id     int
	Title  string
	Year   string
	ImdbId string
	Type   string
	Poster string
	Genre  []Genre
	Writer string
	Actors string
	// Rating    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type GetAPI struct {
	Title  string
	Year   string
	ImdbId string
	Type   string
	Poster string
	Genre  string
	Writer string
	Actors string
}

type Genre struct {
	Name      string
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	//
	GetAllMovie(ctx context.Context) ([]Movie, error)
	GetAPI(ctx context.Context, ImdbId string) (GetAPI, error)
	CreateMovieAPI(ctx context.Context, Movie Movie) (Movie, error)
	MovieDetail(ctx context.Context, id int) (Movie, error)
	ScanGenre(ctx context.Context, Genre Genre) (Genre, error)
	SearchMovie(ctx context.Context, title string) ([]Movie, error)
	FilterGenre(ctx context.Context, genre string) ([]Movie, error)
	FilterOrder(ctx context.Context, order string) ([]Movie, error)
	DeleteAll(ctx context.Context) error
	DeleteMovie(ctx context.Context, id int) (Movie, error)
	UpdateMovie(ctx context.Context, Movie Movie) error
	// CreateGenreAPI(ctx context.Context, Genre Genre) (Genre, error)
	// GetMovieAPI(ctx context.Context, Movie Movie) (Movie, error)
	// MovieDetail(ctx context.Context, id int) (Movie, error)
	// Delete(ctx context.Context, Movie Movie) error
}

type Repository interface {
	GetAllMovie(ctx context.Context) ([]Movie, error)
	DeleteAll(ctx context.Context) error
	UpdateMovie(ctx context.Context, id int, Title string, Type string) error
	DeleteMovie(ctx context.Context, id int) (Movie, error)
	GetAPI(ctx context.Context, ImdbId string) (GetAPI, error)
	CreateMovieAPI(ctx context.Context, Title string, Year string, ImdbId string, Type string, Poster string, Genre []Genre, Writer string, Actors string) (Movie, error)
	MovieDetail(ctx context.Context, id int) (Movie, error)
	ScanGenre(ctx context.Context, name string) (Genre, error)
	SearchMovie(ctx context.Context, title string) ([]Movie, error)
	FilterGenre(ctx context.Context, genre string) ([]Movie, error)
	FilterOrder(ctx context.Context, order string) ([]Movie, error)
	// CreateGenreAPI(ctx context.Context, name string) (Genre, error)
	// GetMovieAPI(ctx context.Context, Movie Movie) (Movie, error)
}
