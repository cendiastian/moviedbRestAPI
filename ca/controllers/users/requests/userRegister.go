package requests

import "project/ca/business/users"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserRegister) ToDomain() users.User {
	return users.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
