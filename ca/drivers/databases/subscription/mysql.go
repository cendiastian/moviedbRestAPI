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

func (rep *MysqlsubsRepository) Update(ctx context.Context, domain subscription.SubcriptionPlan) error {
	subs := FromDomain(domain)
	result := rep.Connect.Where("id = ?", subs.Id).Updates(&SubcriptionPlan{Name: subs.Name, Price: subs.Price})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlsubsRepository) CreatePlan(ctx context.Context, domain subscription.SubcriptionPlan) (subscription.SubcriptionPlan, error) {
	subs := FromDomain(domain)
	result := rep.Connect.Create(&subs)

	if result.Error != nil {
		return subscription.SubcriptionPlan{}, result.Error
	}

	return subs.ToDomain(), nil
}

func (rep *MysqlsubsRepository) DetailTrans(ctx context.Context, id int) (subscription.SubcriptionPlan, error) {
	var subs SubcriptionPlan
	result := rep.Connect.Preload("Transaction", "Id = ?", id).First(&subs)
	if result.Error != nil {
		return subscription.SubcriptionPlan{}, result.Error
	}
	return subs.ToDomain(), nil
}
