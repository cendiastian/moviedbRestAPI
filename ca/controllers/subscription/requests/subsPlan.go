package requests

import (
	"project/ca/business/subscription"
)

type SubcriptionPlan struct {
	// Id      int    `json:"id"`
	Name    string `json:"name"`
	Expired string `json:"expired"`
	Price   int    `json:"price"`
}

func (subs *SubcriptionPlan) ToDomain() subscription.SubcriptionPlan {
	return subscription.SubcriptionPlan{
		Name:    subs.Name,
		Expired: subs.Expired,
		Price:   subs.Price,
	}
}
