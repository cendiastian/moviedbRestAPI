package genres_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"project/business/genres"
	_mockgenreRepository "project/business/genres/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var genreRepository _mockgenreRepository.Repository
var genreService genres.Usecase
var genreDomain genres.Genre
var genresDomain []genres.Genre

func setup() {
	genreService = genres.NewGenreUsecase(&genreRepository, time.Hour*1)

	genreDomain = genres.Genre{
		Id:   1,
		Name: "cen",
	}
	genresDomain = append(genresDomain, genreDomain)
}

func TestGetAllGenre(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		genreRepository.On("GetAllGenre",
			mock.Anything).Return(genresDomain, nil).Once()

		user, err := genreService.GetAllGenre(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		genreRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		genreRepository.On("GetAllGenre",
			mock.Anything).Return([]genres.Genre{}, errors.New("Unexpected Error")).Once()

		user, err := genreService.GetAllGenre(context.Background())

		assert.Error(t, err)
		assert.Equal(t, user, []genres.Genre{})

		genreRepository.AssertExpectations(t)
	})
}
