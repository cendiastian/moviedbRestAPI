package responses

import (
	"project/ca/business/users"
	"time"
)

type LoginResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	// Transaction []transactions.TransactionRespone `json:"transaction"`
	// Ratings   []ratings.Ratings
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomainLogin(domain users.User) LoginResponse {
	return LoginResponse{
		Id:    domain.Id,
		Name:  domain.Name,
		Email: domain.Email,
		// Token: domain.Token,
		// Transaction: transactions.ToListDomain(domain.Transaction),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
