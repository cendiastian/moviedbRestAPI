package transactions

import (
	"context"
	"project/business/payments"
	"project/business/subscription"
	"time"
)

type Transaction struct {
	Id                int
	Payment_method_id int
	Payment           payments.Payment_method
	User_Id           int
	Username          string
	Plan_Id           int
	Subscription_Plan subscription.SubcriptionPlan
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
