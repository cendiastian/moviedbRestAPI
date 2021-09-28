package transactions

import (
	"context"
	"time"
)

type Transaction struct {
	Id                int
	Payment_method_id int
	Payment           string
	User_Id           int
	Username          string
	Plan_Id           int
	Subscription_Plan string
	Price             int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Usecase interface {
	CreateTransaction(ctx context.Context, Transaction Transaction) (Transaction, error)
	DetailTrans(ctx context.Context, id int) (Transaction, error)
}

type Repository interface {
	CreateTransaction(ctx context.Context, Transaction Transaction) (Transaction, error)
	DetailTrans(ctx context.Context, id int) (Transaction, error)
}
