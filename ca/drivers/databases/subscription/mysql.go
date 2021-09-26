package subscription

import (
	"context"
	"project/ca/business/subscription"

	"gorm.io/gorm"
)

type MysqlsubsRepository struct {
	Connect *gorm.DB
}

func NewMysqlsubsRepository(connect *gorm.DB) subscription.Repository {
	return &MysqlsubsRepository{
		Connect: connect,
	}
}

func (rep *MysqlsubsRepository) GetAll(ctx context.Context) ([]subscription.SubcriptionPlan, error) {
	var subs []SubcriptionPlan
	result := rep.Connect.Find(&subs)
	if result.Error != nil {
		return []subscription.SubcriptionPlan{}, result.Error
	}
	return ToListDomain(subs), nil
}

func (rep *MysqlsubsRepository) Detail(ctx context.Context, id int) (subscription.SubcriptionPlan, error) {
	var subs SubcriptionPlan
	result := rep.Connect.First(&subs, "id= ?", id)
	if result.Error != nil {
		return subscription.SubcriptionPlan{}, result.Error
	}
	return subs.ToDomain(), nil
}

func (rep *MysqlsubsRepository) Delete(ctx context.Context, id int) (subscription.SubcriptionPlan, error) {
	var subs SubcriptionPlan
	result := rep.Connect.Where("id = ?", id).Delete(&subs)

	if result.Error != nil {
		return subscription.SubcriptionPlan{}, result.Error
	}

	return subs.ToDomain(), nil
}

func (rep *MysqlsubsRepository) Update(ctx context.Context, id int, name string, price int) error {
	result := rep.Connect.Where("id = ?", id).Updates(&SubcriptionPlan{Name: name, Price: price})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlsubsRepository) CreatePlan(ctx context.Context, name string, expired string, price int) (subscription.SubcriptionPlan, error) {
	subs := SubcriptionPlan{
		Name:    name,
		Expired: expired,
		Price:   price,
	}
	result := rep.Connect.Create(&subs)

	if result.Error != nil {
		return subscription.SubcriptionPlan{}, result.Error
	}

	return subs.ToDomain(), nil
}
