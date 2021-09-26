package requests

import "project/ca/business/transactions"

type Update struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (pay *Update) ToDomain() transactions.Payment_method {
	return transactions.Payment_method{
		Id:     pay.Id,
		Name:   pay.Name,
		Status: pay.Status,
	}
}
