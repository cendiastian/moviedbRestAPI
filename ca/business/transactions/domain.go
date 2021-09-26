package transactions

import (
	"context"
	"time"
)

type Payment_method struct {
	Id          int
	Name        string
	Status      int
	Transaction Transaction
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Transaction struct {
	Id                int
	Payment_method_id int
	User_Id           int
	Plan_Id           int
	Price             int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Payment_method, error)
	Detail(ctx context.Context, id int) (Payment_method, error)
	Update(ctx context.Context, Payment_method Payment_method) error
	Delete(ctx context.Context, id int) (Payment_method, error)
	Register(ctx context.Context, Payment_method Payment_method) (Payment_method, error)
	CreateTransaction(ctx context.Context, Transaction Transaction) (Transaction, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Payment_method, error)
	Detail(ctx context.Context, id int) (Payment_method, error)
	Update(ctx context.Context, id int, name string, password int) error
	Delete(ctx context.Context, id int) (Payment_method, error)
	Register(ctx context.Context, name string, status int) (Payment_method, error)
	CreateTransaction(ctx context.Context, payment_method_id int, user_Id int, plan_Id int) (Transaction, error)
}
