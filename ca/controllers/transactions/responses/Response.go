package responses

import (
	"project/ca/business/transactions"
	"time"
)

type PaymentResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TransactionRespone struct {
	Id                int       `json:"id"`
	Payment_method_id int       `json:"payment_method_id"`
	User_Id           int       `json:"user_id"`
	Plan_id           int       `json:"plan_id"`
	Price             int       `json:"price"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func FromDomain(domain transactions.Payment_method) PaymentResponse {
	return PaymentResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(domain []transactions.Payment_method) (response []PaymentResponse) {
	for _, user := range domain {
		response = append(response, FromDomain(user))
	}
	return
}

func FromDomainTransaction(domain transactions.Transaction) TransactionRespone {
	return TransactionRespone{
		Payment_method_id: domain.Payment_method_id,
		User_Id:           domain.User_Id,
		Plan_id:           domain.Plan_Id,
		Price:             domain.Price,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
