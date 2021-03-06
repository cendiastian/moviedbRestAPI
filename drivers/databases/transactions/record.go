package transactions

import (
	"project/business/transactions"
	"project/drivers/databases/payments"
	"project/drivers/databases/subscription"

	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id                int `gorm:"primaryKey"`
	Payment_method_id int
	Payment_method    payments.Payment_method `gorm:"foreignKey:Payment_method_id"`
	User_Id           int
	// User              users.Users
	Plan_Id           int
	Subscription_Plan subscription.SubcriptionPlan `gorm:"foreignKey:Plan_Id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (subs *Transaction) ToDomainTransaction() transactions.Transaction {
	return transactions.Transaction{
		Id:                subs.Id,
		Payment_method_id: subs.Payment_method_id,
		Payment:           subs.Payment_method.ToDomain(),
		User_Id:           subs.User_Id,
		// Username: subs.User.Name,
		Plan_Id:           subs.Plan_Id,
		Subscription_Plan: subs.Subscription_Plan.ToDomain(),
		Price:             subs.Subscription_Plan.Price,
		CreatedAt:         subs.CreatedAt,
		UpdatedAt:         subs.UpdatedAt,
	}
}

func ToListDomain(data []Transaction) (result []transactions.Transaction) {
	result = []transactions.Transaction{}
	for _, trans := range data {
		result = append(result, trans.ToDomainTransaction())
	}
	return
}

func FromDomainTransaction(domain transactions.Transaction) Transaction {
	return Transaction{
		Id:                domain.Id,
		Payment_method_id: domain.Payment_method_id,
		User_Id:           domain.User_Id,
		Plan_Id:           domain.Plan_Id,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
