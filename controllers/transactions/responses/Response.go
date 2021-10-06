package responses

import (
	"project/business/transactions"
	payments "project/controllers/payments/responses"
	subscriptions "project/controllers/subscription/responses"
	"time"
)

type TransactionRespone struct {
	Id                int                        `json:"id"`
	Payment           payments.PaymentResponse   `json:"payment"`
	Subscription_Plan subscriptions.PlanResponse `json:"subscription_plan"`
	Price             int                        `json:"price"`
	CreatedAt         time.Time                  `json:"createdAt"`
	UpdatedAt         time.Time                  `json:"updatedAt"`
}

type CreateTransactionRespone struct {
	Id                int       `json:"id"`
	Payment_method_id int       `json:"payment_method_id"`
	User_Id           int       `json:"user_id"`
	Plan_Id           int       `json:"plan_id"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func FromDomainTransaction(domain transactions.Transaction) TransactionRespone {
	return TransactionRespone{
		Id:      domain.Id,
		Payment: payments.FromDomain(domain.Payment),
		// Username:          domain.Username,
		Subscription_Plan: subscriptions.FromDomain(domain.Subscription_Plan),
		Price:             domain.Price,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func CreateTransaction(domain transactions.Transaction) CreateTransactionRespone {
	return CreateTransactionRespone{
		Id:                domain.Id,
		Payment_method_id: domain.Payment_method_id,
		User_Id:           domain.User_Id,
		Plan_Id:           domain.Plan_Id,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func ToListDomain(domain []transactions.Transaction) (response []TransactionRespone) {
	for _, rate := range domain {
		response = append(response, FromDomainTransaction(rate))
	}
	return
}
