package requests

import (
	"project/ca/business/subscription"
)

type Update struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (subs *Update) ToDomain() subscription.SubcriptionPlan {
	return subscription.SubcriptionPlan{
		Id:    subs.Id,
		Name:  subs.Name,
		Price: subs.Price,
	}
}
