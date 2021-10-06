package transactions_test

import (
	"context"
	"errors"
	"testing"
	"time"

	premium "project/business/premium"
	_mockProRepository "project/business/premium/mocks"
	"project/business/subscription"
	_mockSubsRepository "project/business/subscription/mocks"
	"project/business/transactions"
	_mocktransactionRepository "project/business/transactions/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepository _mocktransactionRepository.Repository
var proRepository _mockProRepository.Repository
var subsRepository _mockSubsRepository.Repository
var transactionService transactions.Usecase
var transactionDomain transactions.Transaction
var proDomain premium.Premium
var subsDomain subscription.SubcriptionPlan
var transactionsDomain []transactions.Transaction

func setup() {
	transactionService = transactions.NewTransUsecase(&transactionRepository, time.Hour*1, &proRepository, &subsRepository)
	proDomain = premium.Premium{
		UserId:  1,
		Type:    false,
		Expired: time.Time{},
	}
	subsDomain = subscription.SubcriptionPlan{
		Id:      1,
		Name:    "cen",
		Expired: "exp",
		Exp:     1,
		Price:   1,
	}
	transactionDomain = transactions.Transaction{
		Id:                1,
		Payment_method_id: 1,
		User_Id:           1,
		Plan_Id:           1,
		// Subscription_Plan subscription.SubcriptionPlan

	}
	transactionsDomain = append(transactionsDomain, transactionDomain)
}

func TestDetailTrans(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		transactionRepository.On("DetailTrans",
			mock.Anything,
			mock.AnythingOfType("int")).Return(transactionDomain, nil).Once()

		transaction, err := transactionService.DetailTrans(context.Background(), transactionDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, transaction)

		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		transactionRepository.On("DetailTrans",
			mock.Anything,
			mock.AnythingOfType("int")).Return(transactions.Transaction{}, errors.New("Unexpected Error")).Once()

		transaction, err := transactionService.DetailTrans(context.Background(), transactionDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, transaction, transactions.Transaction{})

		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test case 3 | Zero", func(t *testing.T) {
		// transactionRepository.On("DetailTrans",
		// 	mock.Anything,
		// 	mock.AnythingOfType("int")).Return(transactions.Transaction{}, nil).Once()

		transaction, _ := transactionService.DetailTrans(context.Background(), 0)

		assert.Zero(t, transaction.Id)
		// assert.Equal(t, transaction, transactions.Transaction{})

		// transactionRepository.AssertExpectations(t)
	})
}

func TestCreateTransaction(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		transactionRepository.On("CreateTransaction",
			mock.Anything,
			mock.AnythingOfType("transactions.Transaction")).Return(transactionDomain, nil).Once()
		subsRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(subsDomain, nil).Once()
		proRepository.On("Save",
			mock.Anything,
			mock.AnythingOfType("premium.Premium")).Return(proDomain, nil).Once()
		transaction, err := transactionService.CreateTransaction(context.Background(), transactions.Transaction{
			Payment_method_id: 1,
			User_Id:           1,
			Plan_Id:           1,
		})

		assert.Nil(t, err)
		assert.Equal(t, 1, transaction.User_Id)
	})

	t.Run("Test case 2 | Error Registry", func(t *testing.T) {
		transactionRepository.On("CreateTransaction",
			mock.Anything,
			mock.AnythingOfType("transactions.Transaction")).Return(transactions.Transaction{}, errors.New("Unexpected Error")).Once()
		proRepository.On("Save",
			mock.Anything,
			mock.AnythingOfType("premium.Premium")).Return(proDomain, nil).Once()
		transaction, err := transactionService.CreateTransaction(context.Background(), transactions.Transaction{
			Payment_method_id: 1,
			User_Id:           1,
			Plan_Id:           1,
		})

		assert.Error(t, err)
		assert.Equal(t, transaction, transactions.Transaction{})
	})

	t.Run("Test Case 3 | Invalid Empty", func(t *testing.T) {
		transactionRepository.On("CreateTransaction",
			mock.Anything,
			mock.AnythingOfType("transactions.Transaction")).Return(transactionDomain, nil).Once()
		_, err := transactionService.CreateTransaction(context.Background(), transactions.Transaction{
			Payment_method_id: 0,
			User_Id:           1,
			Plan_Id:           1,
		})
		assert.NotNil(t, err)
	})

	// t.Run("Test case 4 | Error Save", func(t *testing.T) {
	// 	transactionRepository.On("CreateTransaction",
	// 		mock.Anything,
	// 		mock.AnythingOfType("transactions.Transaction")).Return(transactionDomain, nil).Once()
	// 	subsRepository.On("Detail",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(subsDomain, nil).Once()
	// 	proRepository.On("Save",
	// 		mock.Anything,
	// 		mock.AnythingOfType("premium.Premium")).Return(premium.Premium{}, errors.New("Unexpected Error")).Once()
	// 		transaction, err := transactionService.CreateTransaction(context.Background(), transactions.Transaction{
	// 		Payment_method_id: 1,
	// 		User_Id:           1,
	// 		Plan_Id:           1,
	// 	})

	// 	assert.Error(t, err)
	// 	assert.Equal(t, transaction, transactions.Transaction{})
	// })

}
