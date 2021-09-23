package responses

import (
	"project/ca/business/users"
	"time"
)

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain users.User) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
