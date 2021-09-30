package responses

import (
	"project/ca/business/users"
	transactions "project/ca/controllers/transactions/responses"
	"time"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Token       string                            `json:"token"`
	Transaction []transactions.TransactionRespone `json:"transaction"`
	// Ratings   []ratings.Ratings
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain users.User) UserResponse {
	return UserResponse{
		Id:    domain.Id,
		Name:  domain.Name,
		Email: domain.Email,
		// Token:       domain.Token,
		Transaction: transactions.ToListDomain(domain.Transaction),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ToListDomain(domain []users.User) (response []UserResponse) {
	for _, user := range domain {
		response = append(response, FromDomain(user))
	}
	return
}
