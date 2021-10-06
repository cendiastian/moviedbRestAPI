package ratings_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"project/business/ratings"
	_mockratingRepository "project/business/ratings/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ratingRepository _mockratingRepository.Repository
var ratingService ratings.Usecase
var ratingDomain ratings.Ratings

// var ratingsDomain []ratings.Ratings

func setup() {
	ratingService = ratings.NewRateUsecase(&ratingRepository, time.Hour*1)

	ratingDomain = ratings.Ratings{
		MovieId: 1,
		UserId:  1,
		Rate:    1,
	}
	// ratingsDomain = append(ratingsDomain, ratingDomain)
}

func TestDetail(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		ratingRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(ratingDomain, nil).Once()

		rating, err := ratingService.Detail(context.Background(), ratingDomain)

		assert.NoError(t, err)
		assert.NotNil(t, rating)

		ratingRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		ratingRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(ratings.Ratings{}, errors.New("Unexpected Error")).Once()

		rating, err := ratingService.Detail(context.Background(), ratingDomain)

		assert.Error(t, err)
		assert.Equal(t, rating, ratings.Ratings{})

		ratingRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		ratingRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(ratingDomain, nil).Once()
		ratingRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil).Once()
		err := ratingService.Delete(context.Background(), ratingDomain)

		assert.NoError(t, err)
		// assert.NotNil(t, rating, ratingDomain)

		ratingRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		ratingRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(ratingDomain, nil).Once()
		ratingRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(errors.New("Unexpected Error")).Once()
		err := ratingService.Delete(context.Background(), ratingDomain)

		assert.Error(t, err)
		// assert.NotNil(t, rating, ratingDomain)

		ratingRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		ratingRepository.On("Create",
			mock.Anything,
			mock.AnythingOfType("ratings.Ratings")).Return(ratingDomain, nil).Once()
		rating, err := ratingService.Create(context.Background(), ratings.Ratings{
			MovieId: 1,
			UserId:  1,
			Rate:    1,
		})

		assert.Nil(t, err)
		assert.Equal(t, 1, rating.MovieId)
	})

	t.Run("Test case 2 | Error Registry", func(t *testing.T) {
		ratingRepository.On("Create",
			mock.Anything,
			mock.AnythingOfType("ratings.Ratings")).Return(ratings.Ratings{}, errors.New("Unexpected Error")).Once()
		rating, err := ratingService.Create(context.Background(), ratings.Ratings{
			MovieId: 1,
			UserId:  1,
			Rate:    1,
		})

		assert.Error(t, err)
		assert.Equal(t, rating, ratings.Ratings{})
	})

	// t.Run("Test Case 3 | Invalid MovieId / UserId", func(t *testing.T) {
	// 	ratingRepository.On("Create",
	// 		mock.Anything,
	// 		mock.AnythingOfType("ratings.Ratings")).Return(ratingDomain, nil).Once()
	// 	_, err := ratingService.Create(context.Background(), ratings.Ratings{
	// 		MovieId: 0,
	// 		UserId:  1,
	// 		Rate:    1,
	// 	})
	// 	assert.NotNil(t, err)
	// })

}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		ratingRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(ratings.Ratings{}, nil).Once()
		ratingRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("ratings.Ratings")).Return(nil).Once()

		err := ratingService.Update(context.Background(), ratings.Ratings{
			MovieId: 1,
			UserId:  1,
			Rate:    1,
		})

		assert.NoError(t, err)
		// assert.Equal(t, rating.Id, 1)

		ratingRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		ratingRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(ratingDomain, nil).Once()
		ratingRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("ratings.Ratings")).Return(errors.New("Unexpected Error")).Once()

		err := ratingService.Update(context.Background(), ratings.Ratings{
			MovieId: 1,
			UserId:  1,
			Rate:    1,
		})

		assert.Error(t, err)
		// assert.Equal(t, rating.Id, 1)

		ratingRepository.AssertExpectations(t)
	})
}
