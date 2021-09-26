package subscription

import (
	"context"
	"project/ca/business/transactions"
	"time"
)

type SubcriptionPlan struct {
	Id          int
	Name        string
	Expired     string
	Price       int
	Transaction transactions.Transaction
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]SubcriptionPlan, error)
	Detail(ctx context.Context, id int) (SubcriptionPlan, error)
	Update(ctx context.Context, subs SubcriptionPlan) error
	Delete(ctx context.Context, id int) (SubcriptionPlan, error)
	CreatePlan(ctx context.Context, subs SubcriptionPlan) (SubcriptionPlan, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]SubcriptionPlan, error)
	Detail(ctx context.Context, id int) (SubcriptionPlan, error)
	Update(ctx context.Context, id int, name string, price int) error
	Delete(ctx context.Context, id int) (SubcriptionPlan, error)
	CreatePlan(ctx context.Context, name string, expired string, price int) (SubcriptionPlan, error)
}
