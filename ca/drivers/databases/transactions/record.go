package transactions

import (
	"project/ca/business/transactions"
	"time"
)

type Payment_method struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Status      int
	Transaction transactions.Transaction `gorm:"foreignKey:Payment_method_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type Transaction struct {
	Id                int `gorm:"primaryKey"`
	Payment_method_id int
	User_Id           int
	Plan_Id           int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (subs *Payment_method) ToDomain() transactions.Payment_method {
	return transactions.Payment_method{
		Id:          subs.Id,
		Name:        subs.Name,
		Status:      subs.Status,
		Transaction: subs.Transaction,
		CreatedAt:   subs.CreatedAt,
		UpdatedAt:   subs.UpdatedAt,
	}
}
func (subs *Transaction) ToDomainTransaction() transactions.Transaction {
	return transactions.Transaction{
		Id:                subs.Id,
		Payment_method_id: subs.Payment_method_id,
		User_Id:           subs.User_Id,
		Plan_Id:           subs.Plan_Id,
		CreatedAt:         subs.CreatedAt,
		UpdatedAt:         subs.UpdatedAt,
	}
}

func ToListDomain(data []Payment_method) (result []transactions.Payment_method) {
	result = []transactions.Payment_method{}
	for _, sub := range data {
		result = append(result, sub.ToDomain())
	}
	return
}

func FromDomain(domain transactions.Payment_method) Payment_method {
	return Payment_method{
		Id:          domain.Id,
		Name:        domain.Name,
		Status:      domain.Status,
		Transaction: domain.Transaction,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
