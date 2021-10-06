package genres

import (
	"context"
	"time"
)

type Genre struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAllGenre(ctx context.Context) ([]Genre, error)
}
type Repository interface {
	GetAllGenre(ctx context.Context) ([]Genre, error)
	FirstOrCreate(ctx context.Context, name string) (Genre, error)
}
