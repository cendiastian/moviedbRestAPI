package users

import (
	"project/ca/business/users"
	"project/ca/drivers/databases/transactions"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id          int    `gorm:"primaryKey"`
	Email       string `gorm:"unique"`
	Name        string
	Password    string
	Transaction []transactions.Transaction `gorm:"foreignKey:User_Id"`
	// Ratings     []ratings.Ratings          `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.User {
	return users.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Transaction: transactions.ToListDomain(user.Transaction),
		// Ratings:     user.Ratings,
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
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		// Transaction: transactions.ToListDomain(domain.Transaction),
		// Ratings:     domain.Ratings,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
