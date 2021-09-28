package requests

import "project/ca/business/payments"

type Update struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (pay *Update) ToDomain() payments.Payment_method {
	return payments.Payment_method{
		Id:     pay.Id,
		Name:   pay.Name,
		Status: pay.Status,
	}
}
