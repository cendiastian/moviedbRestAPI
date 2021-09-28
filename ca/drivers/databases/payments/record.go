package payments

import (
	"project/ca/business/payments"
	"time"

	"gorm.io/gorm"
)

type Payment_method struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Status int
	// Transaction []transactions.Transaction `gorm:"foreignKey:Payment_method_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (subs *Payment_method) ToDomain() payments.Payment_method {
	return payments.Payment_method{
		Id:     subs.Id,
		Name:   subs.Name,
		Status: subs.Status,
		// Transaction: subs.Transaction,
		CreatedAt: subs.CreatedAt,
		UpdatedAt: subs.UpdatedAt,
	}
}

func ToListDomain(data []Payment_method) (result []payments.Payment_method) {
	result = []payments.Payment_method{}
	for _, sub := range data {
		result = append(result, sub.ToDomain())
	}
	return
}

func FromDomain(domain payments.Payment_method) Payment_method {
	return Payment_method{
		Id:     domain.Id,
		Name:   domain.Name,
		Status: domain.Status,
		// Transaction: domain.Transaction,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
