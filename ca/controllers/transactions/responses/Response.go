package responses

import (
	"project/ca/business/transactions"
	"time"
)

type TransactionRespone struct {
	Id int `json:"id"`
	// Payment_method_id int       `json:"payment_method_id"`
	// User_Id           int       `json:"user_id"`
	Username string `json:"username"`
	Payment  string `json:"payment"`
	// Plan_id           int       `json:"plan_id"`
	Subscription_Plan string    `json:"Subscription_Plan"`
	Price             int       `json:"price"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func FromDomainTransaction(domain transactions.Transaction) TransactionRespone {
	return TransactionRespone{
		Id:                domain.Id,
		Payment:           domain.Payment,
		Username:          domain.Username,
		Subscription_Plan: domain.Subscription_Plan,
		Price:             domain.Price,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
