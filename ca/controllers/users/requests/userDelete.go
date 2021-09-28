package requests

import "project/ca/business/users"

type UserDelete struct {
	Id int `json:"id"`
}

func (user *UserDelete) ToDomain() users.User {
	return users.User{
		Id: user.Id,
	}
}
