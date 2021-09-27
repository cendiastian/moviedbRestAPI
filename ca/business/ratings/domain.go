package ratings

import (
	"context"
	"time"
)

type Rating struct {
	Movie_Id  int
	User_Id   int
	Rate      float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	// GetAll(ctx context.Context) ([]Rating, error)
	// Detail(ctx context.Context, id int) (Rating, error)
	Update(ctx context.Context, Rating Rating) error
	Delete(ctx context.Context, Rating Rating) error
	Create(ctx context.Context, Rating Rating) (Rating, error)
}

type Repository interface {
	// GetAll(ctx context.Context) ([]Rating, error)
	// Detail(ctx context.Context, id int) (Rating, error)
	Update(ctx context.Context, Movie_Id int, User_Id int, Rate float32) error
	Delete(ctx context.Context, Movie_Id int, User_Id int) error
	Create(ctx context.Context, Movie_Id int, User_Id int, Rate float32) (Rating, error)
}
