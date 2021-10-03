package payments_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"project/business/payments"
	_mockpaymentRepository "project/business/payments/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var paymentRepository _mockpaymentRepository.Repository
var paymentService payments.Usecase
var paymentDomain payments.Payment_method
var paymentsDomain []payments.Payment_method

func setup() {
	paymentService = payments.NewPaymentUsecase(&paymentRepository, time.Hour*1)

	paymentDomain = payments.Payment_method{
		Id:     1,
		Name:   "cen",
		Status: 1,
	}
	paymentsDomain = append(paymentsDomain, paymentDomain)
}

func TestDetail(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()

		payment, err := paymentService.Detail(context.Background(), paymentDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, payment)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(payments.Payment_method{}, errors.New("Unexpected Error")).Once()

		payment, err := paymentService.Detail(context.Background(), paymentDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, payment, payments.Payment_method{})

		paymentRepository.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		paymentRepository.On("GetAll",
			mock.Anything).Return(paymentsDomain, nil).Once()

		payment, err := paymentService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, payment)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("GetAll",
			mock.Anything).Return([]payments.Payment_method{}, errors.New("Unexpected Error")).Once()

		payment, err := paymentService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, payment, []payments.Payment_method{})

		paymentRepository.AssertExpectations(t)
	})
}
func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		paymentRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		payment, err := paymentService.Delete(context.Background(), paymentDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, payment, paymentDomain)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		paymentRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int")).Return(payments.Payment_method{}, errors.New("Unexpected Error")).Once()
		payment, err := paymentService.Delete(context.Background(), paymentDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, payment, payments.Payment_method{})

		paymentRepository.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		paymentRepository.On("Register",
			mock.Anything,
			mock.AnythingOfType("payments.Payment_method")).Return(paymentDomain, nil).Once()
		payment, err := paymentService.Register(context.Background(), payments.Payment_method{
			Name:   "cen",
			Status: 1,
		})

		assert.Nil(t, err)
		assert.Equal(t, "cen", payment.Name)
	})

	t.Run("Test case 2 | Error Registry", func(t *testing.T) {
		paymentRepository.On("Register",
			mock.Anything,
			mock.AnythingOfType("payments.Payment_method")).Return(payments.Payment_method{}, errors.New("Unexpected Error")).Once()
		payment, err := paymentService.Register(context.Background(), payments.Payment_method{
			Name:   "asd",
			Status: 1,
		})

		assert.Error(t, err)
		assert.Equal(t, payment, payments.Payment_method{})
	})

	// t.Run("Test Case 3 | Invalid Name Empty", func(t *testing.T) {
	// 	paymentRepository.On("Register",
	// 		mock.Anything,
	// 		mock.AnythingOfType("payments.Payment_method")).Return(paymentDomain, nil).Once()
	// 	_, err := paymentService.Register(context.Background(), payments.Payment_method{
	// 		Name:   "",
	// 		Status: 1,
	// 	})
	// 	assert.NotNil(t, err)
	// })

}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		paymentRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("payments.Payment_method")).Return(paymentDomain, nil).Once()

		payment, err := paymentService.Update(context.Background(), payments.Payment_method{
			Id:     1,
			Name:   "asd",
			Status: 1,
		})

		assert.NoError(t, err)
		assert.Equal(t, payment.Id, 1)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		paymentRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("payments.Payment_method")).Return(payments.Payment_method{}, errors.New("Unexpected Error")).Once()

		payment, err := paymentService.Update(context.Background(), payments.Payment_method{
			Id:     1,
			Name:   "asd",
			Status: 1,
		})

		assert.Error(t, err)
		assert.Equal(t, payment, payments.Payment_method{})

		paymentRepository.AssertExpectations(t)
	})
}
