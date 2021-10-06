package subscription

import (
	"project/business/subscription"
	"time"

	"gorm.io/gorm"
)

type SubcriptionPlan struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Expired   string
	Exp       int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (subs *SubcriptionPlan) ToDomain() subscription.SubcriptionPlan {
	return subscription.SubcriptionPlan{
		Id:        subs.Id,
		Name:      subs.Name,
		Expired:   subs.Expired,
		Exp:       subs.Exp,
		Price:     subs.Price,
		CreatedAt: subs.CreatedAt,
		UpdatedAt: subs.UpdatedAt,
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
		Id:        domain.Id,
		Name:      domain.Name,
		Expired:   domain.Expired,
		Exp:       domain.Exp,
		Price:     domain.Price,
		UpdatedAt: domain.UpdatedAt,
	}
}
