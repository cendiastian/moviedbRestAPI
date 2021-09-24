package requests

import "project/ca/business/users"

type UserUpdate struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserUpdate) ToDomain() users.User {
	return users.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}
}
