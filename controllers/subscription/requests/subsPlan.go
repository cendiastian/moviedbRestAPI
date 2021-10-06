package requests

import (
	"project/business/subscription"
)

type SubcriptionPlan struct {
	// Id      int    `json:"id"`
	Name    string `json:"name"`
	Expired string `json:"expired"`
	Exp     int    `json:"exp"`
	Price   int    `json:"price"`
}

func (subs *SubcriptionPlan) ToDomain() subscription.SubcriptionPlan {

	return subscription.SubcriptionPlan{
		Name:    subs.Name,
		Expired: subs.Expired,
		Exp:     subs.Exp,
		Price:   subs.Price,
	}
}
