package requests

import "project/ca/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() users.User {
	return users.User{
		Email:    user.Email,
		Password: user.Password,
	}
}
