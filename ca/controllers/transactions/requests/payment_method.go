package requests

import (
	"project/ca/business/transactions"
	"time"
)

type Payment_method struct {
	Name      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (pay *Payment_method) ToDomain() transactions.Payment_method {
	return transactions.Payment_method{
		Name:   pay.Name,
		Status: pay.Status,
	}
}
