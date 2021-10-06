package responses

import (
	"project/business/users"
	transactions "project/controllers/transactions/responses"
	"time"
)

type LoginResponse struct {
	Id          int                               `json:"id"`
	Name        string                            `json:"name"`
	Email       string                            `json:"email"`
	Token       string                            `json:"token"`
	Transaction []transactions.TransactionRespone `json:"transaction"`
	Premium     ProResponse                       `json:"premium"`
	// Ratings   []ratings.Ratings
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomainLogin(domain users.User) LoginResponse {
	return LoginResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Email:       domain.Email,
		Token:       domain.Token,
		Transaction: transactions.ToListDomain(domain.Transaction),
		Premium:     FromDomainPro(domain.Premium),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
