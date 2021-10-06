package payments

import (
	"context"
	"project/business/payments"

	"gorm.io/gorm"
)

type MysqlpayRepository struct {
	Connect *gorm.DB
}

func NewMysqlpayRepository(connect *gorm.DB) payments.Repository {
	return &MysqlpayRepository{
		Connect: connect,
	}
}

func (rep *MysqlpayRepository) GetAll(ctx context.Context) ([]payments.Payment_method, error) {
	var pay []Payment_method
	result := rep.Connect.Find(&pay)
	if result.Error != nil {
		return []payments.Payment_method{}, result.Error
	}
	return ToListDomain(pay), nil
}

func (rep *MysqlpayRepository) Detail(ctx context.Context, id int) (payments.Payment_method, error) {
	var pay Payment_method
	result := rep.Connect.First(&pay, "id= ?", id)
	if result.Error != nil {
		return payments.Payment_method{}, result.Error
	}
	return pay.ToDomain(), nil
}

func (rep *MysqlpayRepository) Delete(ctx context.Context, id int) (payments.Payment_method, error) {
	var pay Payment_method
	result := rep.Connect.Where("id = ?", id).Delete(&pay)

	if result.Error != nil {
		return payments.Payment_method{}, result.Error
	}

	return pay.ToDomain(), nil
}

func (rep *MysqlpayRepository) Update(ctx context.Context, domain payments.Payment_method) (payments.Payment_method, error) {
	pay := FromDomain(domain)
	result := rep.Connect.Where("id = ?", pay.Id).Updates(&Payment_method{Name: pay.Name, Status: pay.Status})

	if result.Error != nil {
		return payments.Payment_method{}, result.Error
	}

	return pay.ToDomain(), nil
}

func (rep *MysqlpayRepository) Register(ctx context.Context, domain payments.Payment_method) (payments.Payment_method, error) {
	pay := FromDomain(domain)
	result := rep.Connect.Create(&pay)

	if result.Error != nil {
		return payments.Payment_method{}, result.Error
	}

	return pay.ToDomain(), nil
}
