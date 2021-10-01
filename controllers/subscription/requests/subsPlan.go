package requests

import (
	"project/business/subscription"
	"time"
)

type SubcriptionPlan struct {
	// Id      int    `json:"id"`
	Name    string `json:"name"`
	Expired string `json:"expired"`
	Exp     int    `json:"exp"`
	Price   int    `json:"price"`
}

func (subs *SubcriptionPlan) ToDomain() subscription.SubcriptionPlan {
	t := time.Now()
	return subscription.SubcriptionPlan{
		Name:    subs.Name,
		Expired: subs.Expired,
		Exp:     t.AddDate(0, 0, subs.Exp),
		Price:   subs.Price,
	}
}
