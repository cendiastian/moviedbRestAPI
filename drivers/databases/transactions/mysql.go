package transactions

import (
	"context"
	"fmt"
	"project/business/transactions"

	"gorm.io/gorm"
)

type MysqlTransRepository struct {
	Connect *gorm.DB
}

func NewMysqlTransRepository(connect *gorm.DB) transactions.Repository {
	return &MysqlTransRepository{
		Connect: connect,
	}
}

func (rep *MysqlTransRepository) CreateTransaction(ctx context.Context, domain transactions.Transaction) (transactions.Transaction, error) {
	trans := FromDomainTransaction(domain)

	fmt.Println(trans)
	result := rep.Connect.Preload("Subscription_Plan").Preload("Payment_method").Create(&trans)

	if result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return trans.ToDomainTransaction(), nil
}

func (rep *MysqlTransRepository) DetailTrans(ctx context.Context, id int) (transactions.Transaction, error) {
	var trans Transaction
	// var subs subscription.SubcriptionPlan
	result := rep.Connect.Preload("Payment_method").Preload("Subscription_Plan").Where("Id = ?", id).First(&trans)

	if result.Error != nil {
		return transactions.Transaction{}, result.Error
	}

	return trans.ToDomainTransaction(), nil
}

// func (rep *MysqlsubsRepository) DetailTrans(ctx context.Context, id int) (subscription.SubcriptionPlan, error) {
// 	var subs SubcriptionPlan
// 	result := rep.Connect.Preload("Transaction", "Id = ?", id).First(&subs)
// 	if result.Error != nil {
// 		return subscription.SubcriptionPlan{}, result.Error
// 	}
// 	return subs.ToDomain(), nil
// }
