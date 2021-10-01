package premium

import (
	"context"
	"project/business/premium"

	"gorm.io/gorm"
)

type MysqlPremiumRepository struct {
	Connect *gorm.DB
}

func NewMysqlPremiumRepository(connect *gorm.DB) premium.Repository {
	return &MysqlPremiumRepository{
		Connect: connect,
	}
}

func (rep *MysqlPremiumRepository) Detail(ctx context.Context, id int) (premium.Premium, error) {
	var pay Premium
	result := rep.Connect.First(&pay, "user_id= ?", id)
	if result.Error != nil {
		return premium.Premium{}, result.Error
	}
	return pay.ToDomain(), nil
}

func (rep *MysqlPremiumRepository) Save(ctx context.Context, domain premium.Premium) (premium.Premium, error) {
	pay := FromDomain(domain)
	result := rep.Connect.Where("user_id= ?", domain.UserId).Save(&pay)
	// if result.Error != nil {
	// 	result = rep.Connect.Create(&pay)
	if result.Error != nil {
		return premium.Premium{}, result.Error
	}
	// }

	return pay.ToDomain(), nil
}
