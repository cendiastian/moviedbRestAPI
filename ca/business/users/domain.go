package users

import (
	"context"
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	// Login(ctx context.Context, email string, password string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Login(ctx context.Context, User User) (User, error)
	UserDetail(ctx context.Context, id int) (User, error)
	// Update(ctx context.Context, ar *User) error
	// Delete(ctx context.Context, id int64) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Login(ctx context.Context, email string, password string) (User, error)
	UserDetail(ctx context.Context, id int) (User, error)
}
