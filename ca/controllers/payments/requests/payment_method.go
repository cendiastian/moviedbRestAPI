package requests

import (
	"project/ca/business/payments"
	"time"
)

type Payment_method struct {
	Name      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (pay *Payment_method) ToDomain() payments.Payment_method {
	return payments.Payment_method{
		Name:   pay.Name,
		Status: pay.Status,
	}
}
