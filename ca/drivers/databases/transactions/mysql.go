package transactions

import (
	"context"
	"project/ca/business/transactions"

	"gorm.io/gorm"
)

type MysqlpayRepository struct {
	Connect *gorm.DB
}

func NewMysqlpayRepository(connect *gorm.DB) transactions.Repository {
	return &MysqlpayRepository{
		Connect: connect,
	}
}

func (rep *MysqlpayRepository) GetAll(ctx context.Context) ([]transactions.Payment_method, error) {
	var pay []Payment_method
	result := rep.Connect.Find(&pay)
	if result.Error != nil {
		return []transactions.Payment_method{}, result.Error
	}
	return ToListDomain(pay), nil
}

func (rep *MysqlpayRepository) Detail(ctx context.Context, id int) (transactions.Payment_method, error) {
	var pay Payment_method
	result := rep.Connect.First(&pay, "id= ?", id)
	if result.Error != nil {
		return transactions.Payment_method{}, result.Error
	}
	return pay.ToDomain(), nil
}

func (rep *MysqlpayRepository) Delete(ctx context.Context, id int) (transactions.Payment_method, error) {
	var pay Payment_method
	result := rep.Connect.Where("id = ?", id).Delete(&pay)

	if result.Error != nil {
		return transactions.Payment_method{}, result.Error
	}

	return pay.ToDomain(), nil
}

func (rep *MysqlpayRepository) Update(ctx context.Context, id int, name string, status int) error {
	result := rep.Connect.Where("id = ?", id).Updates(&Payment_method{Name: name, Status: status})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rep *MysqlpayRepository) Register(ctx context.Context, name string, status int) (transactions.Payment_method, error) {
	pay := Payment_method{
		Name:   name,
		Status: status,
	}
	result := rep.Connect.Create(&pay)

	if result.Error != nil {
		return transactions.Payment_method{}, result.Error
	}

	return pay.ToDomain(), nil
}
func (rep *MysqlpayRepository) CreateTransaction(ctx context.Context, payment_method_id int, user_id int, plan_id int) (transactions.Transaction, error) {
	pay := Transaction{
		Payment_method_id: payment_method_id,
		User_Id:           user_id,
		Plan_Id:           plan_id,
	}
	result := rep.Connect.Create(&pay)

	if result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return pay.ToDomainTransaction(), nil
}
