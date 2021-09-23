package users

import (
	"project/ca/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.User {
	return users.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListDomain(data []Users) (result []users.User) {
	result = []users.User{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain users.User) Users {
	return Users{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
