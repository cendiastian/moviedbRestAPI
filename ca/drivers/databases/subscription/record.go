package subscription

import (
	"project/ca/business/subscription"
	"project/ca/business/transactions"
	"time"

	"gorm.io/gorm"
)

type SubcriptionPlan struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Expired     string
	Price       int
	Transaction transactions.Transaction `gorm:"foreignKey:Plan_id;references:id;foreignKey:Price;references:Price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (subs *SubcriptionPlan) ToDomain() subscription.SubcriptionPlan {
	return subscription.SubcriptionPlan{
		Id:          subs.Id,
		Name:        subs.Name,
		Expired:     subs.Expired,
		Transaction: subs.Transaction,
		Price:       subs.Price,
		CreatedAt:   subs.CreatedAt,
		UpdatedAt:   subs.UpdatedAt,
	}
}

func ToListDomain(data []SubcriptionPlan) (result []subscription.SubcriptionPlan) {
	result = []subscription.SubcriptionPlan{}
	for _, sub := range data {
		result = append(result, sub.ToDomain())
	}
	return
}

func FromDomain(domain subscription.SubcriptionPlan) SubcriptionPlan {
	return SubcriptionPlan{
		Id:          domain.Id,
		Name:        domain.Name,
		Expired:     domain.Expired,
		Price:       domain.Price,
		Transaction: domain.Transaction,
		UpdatedAt:   domain.UpdatedAt,
	}
}
