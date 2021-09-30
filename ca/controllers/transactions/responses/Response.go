package responses

import (
	"project/ca/business/transactions"
	payments "project/ca/controllers/payments/responses"
	subscriptions "project/ca/controllers/subscription/responses"
	"time"
)

type TransactionRespone struct {
	Id int `json:"id"`
	// Payment_method_id int       `json:"payment_method_id"`
	// User_Id           int       `json:"user_id"`
	// Username string `json:"username"`
	Payment payments.PaymentResponse `json:"payment"`
	// Plan_id           int       `json:"plan_id"`
	Subscription_Plan subscriptions.PlanResponse `json:"subscription_plan"`
	Price             int                        `json:"price"`
	CreatedAt         time.Time                  `json:"createdAt"`
	UpdatedAt         time.Time                  `json:"updatedAt"`
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

func ToListDomain(domain []transactions.Transaction) (response []TransactionRespone) {
	for _, rate := range domain {
		response = append(response, FromDomainTransaction(rate))
	}
	return
}
