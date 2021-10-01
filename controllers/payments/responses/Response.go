package responses

import (
	"project/business/payments"
	"time"
)

type PaymentResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain payments.Payment_method) PaymentResponse {
	return PaymentResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(domain []payments.Payment_method) (response []PaymentResponse) {
	for _, user := range domain {
		response = append(response, FromDomain(user))
	}
	return
}
