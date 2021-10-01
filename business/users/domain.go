package users

import (
	"context"
	"project/business/premium"
	"project/business/transactions"
	"time"
)

type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	Token       string
	Transaction []transactions.Transaction
	Premium     premium.Premium
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	// Login(ctx context.Context, email string, password string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Login(ctx context.Context, User User) (User, error)
	UserDetail(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, User User) error
	Delete(ctx context.Context, User User) error
	Register(ctx context.Context, User User) (User, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Login(ctx context.Context, User User) (User, error)
	UserDetail(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, User User) error
	Delete(ctx context.Context, id int) error
	Register(ctx context.Context, User User) (User, error)
}
