package ratings

import (
	"context"
	// "project/ca/drivers/databases/movies"

	"time"
)

type Ratings struct {
	MovieId int
	// Title   string
	// Movie     movies.Movie
	UserId   int
	Username string
	// User      users.User
	Rate      float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	// GetAllRate(ctx context.Context, id int) (Ratings, error)
	Detail(ctx context.Context, movie int, user int) (Ratings, error)
	Update(ctx context.Context, Ratings Ratings) error
	Delete(ctx context.Context, Ratings Ratings) error
	Create(ctx context.Context, Ratings Ratings) (Ratings, error)
}

type Repository interface {
	// GetAllRate(ctx context.Context, id int) (Ratings, error)
	Detail(ctx context.Context, movie int, user int) (Ratings, error)
	Update(ctx context.Context, Ratings Ratings) error
	Delete(ctx context.Context, MovieId int, UserId int) error
	Create(ctx context.Context, Ratings Ratings) (Ratings, error)
}
