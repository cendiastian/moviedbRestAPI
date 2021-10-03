package subscription_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"project/business/subscription"
	_mockpaymentRepository "project/business/subscription/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var paymentRepository _mockpaymentRepository.Repository
var subscriptionervice subscription.Usecase
var paymentDomain subscription.SubcriptionPlan
var subscriptionDomain []subscription.SubcriptionPlan

func setup() {
	subscriptionervice = subscription.NewSubsUsecase(&paymentRepository, time.Hour*1)

	paymentDomain = subscription.SubcriptionPlan{
		Id:      1,
		Name:    "cen",
		Expired: "1 hari",
		Exp:     time.Time{},
		Price:   1,
	}
	subscriptionDomain = append(subscriptionDomain, paymentDomain)
}

func TestDetail(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()

		payment, err := subscriptionervice.Detail(context.Background(), paymentDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, payment)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(subscription.SubcriptionPlan{}, errors.New("Unexpected Error")).Once()

		payment, err := subscriptionervice.Detail(context.Background(), paymentDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, payment, subscription.SubcriptionPlan{})

		paymentRepository.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		paymentRepository.On("GetAll",
			mock.Anything).Return(subscriptionDomain, nil).Once()

		payment, err := subscriptionervice.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, payment)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("GetAll",
			mock.Anything).Return([]subscription.SubcriptionPlan{}, errors.New("Unexpected Error")).Once()

		payment, err := subscriptionervice.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, payment, []subscription.SubcriptionPlan{})

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
		err := subscriptionervice.Delete(context.Background(), paymentDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, paymentDomain)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		paymentRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, errors.New("Unexpected Error")).Once()
		err := subscriptionervice.Delete(context.Background(), paymentDomain.Id)

		assert.Error(t, err)
		// assert.Equal(t,  subscription.SubcriptionPlan{})

		paymentRepository.AssertExpectations(t)
	})
}

func TestCreatePlan(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		paymentRepository.On("CreatePlan",
			mock.Anything,
			mock.AnythingOfType("subscription.SubcriptionPlan")).Return(paymentDomain, nil).Once()
		payment, err := subscriptionervice.CreatePlan(context.Background(), subscription.SubcriptionPlan{
			// Id:      1,
			Name:    "cen",
			Expired: "1 hari",
			Exp:     time.Now(),
			Price:   1,
		})

		assert.Nil(t, err)
		assert.Equal(t, "cen", payment.Name)
	})

	// t.Run("Test case 2 | Error Registry", func(t *testing.T) {
	// 	paymentRepository.On("CreatePlan",
	// 		mock.Anything,
	// 		mock.AnythingOfType("subscription.SubcriptionPlan")).Return(subscription.SubcriptionPlan{}, errors.New("Unexpected Error")).Once()
	// 	payment, err := subscriptionervice.CreatePlan(context.Background(), subscription.SubcriptionPlan{
	// 		// Id:      1,
	// 		Name:    "cen",
	// 		Expired: "1 hari",
	// 		Exp:     time.Time{},
	// 		Price:   1,
	// 	})

	// 	assert.Error(t, err)
	// 	assert.Equal(t, payment, subscription.SubcriptionPlan{})
	// })

	// t.Run("Test Case 3 | Invalid Name Empty", func(t *testing.T) {
	// 	paymentRepository.On("CreatePlan",
	// 		mock.Anything,
	// 		mock.AnythingOfType("subscription.SubcriptionPlan")).Return(paymentDomain, nil).Once()
	// 	_, err := subscriptionervice.CreatePlan(context.Background(), subscription.SubcriptionPlan{
	// 		// Id:      1,
	// 		Name:    "cen",
	// 		Expired: "1 hari",
	// 		Exp:     time.Time{},
	// 		Price:   1,
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
			mock.AnythingOfType("subscription.SubcriptionPlan")).Return(nil).Once()

		err := subscriptionervice.Update(context.Background(), subscription.SubcriptionPlan{
			Id:      1,
			Name:    "cen",
			Expired: "1 hari",
			Exp:     time.Time{},
			Price:   1,
		})

		assert.NoError(t, err)
		// assert.Equal(t, payment.Id, 1)

		paymentRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		paymentRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(paymentDomain, nil).Once()
		paymentRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("subscription.SubcriptionPlan")).Return(errors.New("Unexpected Error")).Once()

		err := subscriptionervice.Update(context.Background(), subscription.SubcriptionPlan{
			Id:      1,
			Name:    "cen",
			Expired: "1 hari",
			Exp:     time.Time{},
			Price:   1,
		})

		assert.Error(t, err)
		// assert.Equal(t, payment, subscription.SubcriptionPlan{})

		paymentRepository.AssertExpectations(t)
	})
}
