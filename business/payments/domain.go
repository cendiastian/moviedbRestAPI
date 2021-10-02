package payments

import (
	"context"
	"time"
)

type Payment_method struct {
	Id     int
	Name   string
	Status int
	// Transaction []transactions.Transaction
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Payment_method, error)
	Detail(ctx context.Context, id int) (Payment_method, error)
	Update(ctx context.Context, Payment_method Payment_method) error
	Delete(ctx context.Context, id int) (Payment_method, error)
	Register(ctx context.Context, Payment_method Payment_method) (Payment_method, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Payment_method, error)
	Detail(ctx context.Context, id int) (Payment_method, error)
	Update(ctx context.Context, Payment_method Payment_method) error
	Delete(ctx context.Context, id int) (Payment_method, error)
	Register(ctx context.Context, Payment_method Payment_method) (Payment_method, error)
}
