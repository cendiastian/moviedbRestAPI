package movies_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"project/business/genres"
	_mockGenreRepository "project/business/genres/mocks"
	"project/business/movies"
	_mockmovieRepository "project/business/movies/mocks"
	"project/business/omdb"
	_mockAPIRepository "project/business/omdb/mocks"
	"project/business/premium"
	_mockProRepository "project/business/premium/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var movieRepository _mockmovieRepository.Repository
var proRepository _mockProRepository.Repository
var apiRepository _mockAPIRepository.Repository
var gnrRepository _mockGenreRepository.Repository
var movieService movies.Usecase
var movieDomain movies.Movie
var proDomain premium.Premium
var gnrDomain genres.Genre
var apiDomain omdb.GetAPI
var moviesDomain []movies.Movie

// var genresDomain []genres.Genre

func setup() {
	movieService = movies.NewMovieUsecase(&movieRepository, time.Hour*1, &gnrRepository, &apiRepository, &proRepository)
	proDomain = premium.Premium{
		UserId:  1,
		Type:    true,
		Expired: time.Time{},
	}
	gnrDomain = genres.Genre{
		Id:   1,
		Name: "cen",
	}
	// genresDomain := append(genresDomain, gnrDomain)
	apiDomain = omdb.GetAPI{
		Title:  "cen",
		Year:   "cen",
		ImdbId: "cen",
		Type:   "cen",
		Poster: "cen",
		Genre:  "cen can",
		Writer: "cen",
		Actors: "cen",
	}
	movieDomain = movies.Movie{
		Id:     1,
		Title:  "cen",
		Year:   "cen",
		ImdbId: "cen",
		Type:   "cen",
		Poster: "cen",
		// Genre:   movieDomain.Genre,
		// Ratings: movieDomain.Ratings,
		Rating: 5,
		Writer: "cen",
		Actors: "cen",
	}
	moviesDomain = append(moviesDomain, movieDomain)
}

func TestMovieDetail(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		proRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(proDomain, nil).Once()
		movieRepository.On("MovieDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(movieDomain, nil).Once()

		movie, err := movieService.MovieDetail(context.Background(), movieDomain.Id, proDomain.UserId)

		assert.NoError(t, err)
		assert.NotNil(t, movie)

		movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		proRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(proDomain, nil).Once()
		movieRepository.On("MovieDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(movies.Movie{}, errors.New("Unexpected Error")).Once()

		movie, err := movieService.MovieDetail(context.Background(), movieDomain.Id, proDomain.UserId)

		assert.Error(t, err)
		assert.Equal(t, movie, movies.Movie{})

		movieRepository.AssertExpectations(t)
	})
}

func TestGetAllMovie(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("GetAllMovie",
			mock.Anything).Return(moviesDomain, nil).Once()

		movie, err := movieService.GetAllMovie(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, movie)

		movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		movieRepository.On("GetAllMovie",
			mock.Anything).Return([]movies.Movie{}, errors.New("Unexpected Error")).Once()

		movie, err := movieService.GetAllMovie(context.Background())

		assert.Error(t, err)
		assert.Equal(t, movie, []movies.Movie{})

		movieRepository.AssertExpectations(t)
	})
}
func TestDeleteMovie(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("MovieDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(movieDomain, nil).Once()
		movieRepository.On("DeleteMovie",
			mock.Anything,
			mock.AnythingOfType("int")).Return(nil).Once()
		err := movieService.DeleteMovie(context.Background(), movieDomain.Id)

		assert.NoError(t, err)
		// assert.NotNil(t, movie, movieDomain)

		// movieRepository.AssertExpectations(t)
	})

	// t.Run("Test case 2 | DeleteMovie Error", func(t *testing.T) {
	// 	movieRepository.On("MovieDetail",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(movieDomain, nil).Once()
	// 	movieRepository.On("DeleteMovie",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(movies.Movie{}, errors.New("Unexpected Error")).Once()
	// 	err := movieService.DeleteMovie(context.Background(), movieDomain.Id)

	// 	assert.Error(t, err)
	// 	// assert.Equal(t, movie, movies.Movie{})

	// 	// movieRepository.AssertExpectations(t)
	// })
}

func TestCreateMovie(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		apiRepository.On("GetAPI",
			mock.Anything,
			mock.AnythingOfType("string")).Return(apiDomain, nil).Once()
		gnrRepository.On("FirstOrCreate",
			mock.Anything,
			mock.AnythingOfType("string")).Return(gnrDomain, nil).Once()
		movieRepository.On("CreateMovie",
			mock.Anything,
			mock.AnythingOfType("movies.Movie"),
			mock.AnythingOfType("[]genres.Genre")).Return(movieDomain, nil).Once()
		movie, err := movieService.CreateMovie(context.Background(), "1234")

		assert.Nil(t, err)
		assert.Equal(t, movie, movieDomain)
	})

	t.Run("Test case 2 | Error Registry", func(t *testing.T) {
		apiRepository.On("GetAPI",
			mock.Anything,
			mock.AnythingOfType("string")).Return(apiDomain, nil).Once()
		gnrRepository.On("FirstOrCreate",
			mock.Anything,
			mock.AnythingOfType("string")).Return(gnrDomain, nil).Once()
		movieRepository.On("CreateMovie",
			mock.Anything,
			mock.AnythingOfType("movies.Movie"),
			mock.AnythingOfType("[]genres.Genre")).Return(movies.Movie{}, errors.New("Unexpected Error")).Once()
		movie, err := movieService.CreateMovie(context.Background(), "1234")

		assert.NotNil(t, err)
		assert.Equal(t, movie, movies.Movie{})
	})

}

func TestUpdateMovie(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("MovieDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(movieDomain, nil).Once()
		movieRepository.On("UpdateMovie",
			mock.Anything,
			mock.AnythingOfType("movies.Movie")).Return(nil).Once()
		err := movieService.UpdateMovie(context.Background(), movies.Movie{
			Id:   1,
			Type: "test",
		})

		assert.NoError(t, err)
		// assert.Equal(t, 1, movie.Id)

		movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | UpdateMovie Error", func(t *testing.T) {
		movieRepository.On("MovieDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(movieDomain, nil).Once()
		movieRepository.On("UpdateMovie",
			mock.Anything,
			mock.AnythingOfType("movies.Movie")).Return(movies.Movie{}, errors.New("Unexpected Error")).Once()

		err := movieService.UpdateMovie(context.Background(), movies.Movie{
			Id:   1,
			Type: "test",
		})

		assert.Error(t, err)
		// 	// assert.Equal(t, movie, movies.Movie{})

		// 	movieRepository.AssertExpectations(t)
	})

	// t.Run("Test case 3 |Detail Error", func(t *testing.T) {
	// 	movieRepository.On("MovieDetail",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(movies.Movie{}, errors.New("Unexpected Error")).Once()
	// 	movieRepository.On("UpdateMovie",
	// 		mock.Anything,
	// 		mock.AnythingOfType("movies.Movie")).Return(movieDomain, errors.New("Unexpected Error")).Once()

	// 	movie, err := movieService.UpdateMovie(context.Background(), movies.Movie{
	// 		Id:       1,
	// 		Name:     "ddd",
	// 		Email:    "dd@dd.dd",
	// 		Password: "ddd",
	// 	})

	// 	assert.Error(t, err)
	// 	assert.Equal(t, movie, movies.Movie{})

	// 	movieRepository.AssertExpectations(t)
	// })
}
func TestDeleteAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("GetAllMovie",
			mock.Anything).Return(moviesDomain, nil).Once()
		movieRepository.On("DeleteAll",
			mock.Anything).Return(nil).Once()

		err := movieService.DeleteAll(context.Background())

		assert.NoError(t, err)

		// movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		movieRepository.On("GetAllMovie",
			mock.Anything).Return(moviesDomain, nil).Once()
		movieRepository.On("DeleteAll",
			mock.Anything).Return(errors.New(mock.Anything)).Once()

		err := movieService.DeleteAll(context.Background())

		assert.Error(t, err)

		// movieRepository.AssertExpectations(t)
	})
}

func TestSearchMovie(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("SearchMovie",
			mock.Anything,
			mock.AnythingOfType("string")).Return(moviesDomain, nil).Once()

		Job, err := movieService.SearchMovie(context.Background(), "cen")

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		movieRepository.On("SearchMovie",
			mock.Anything,
			mock.AnythingOfType("string")).Return([]movies.Movie{}, errors.New(mock.Anything)).Once()

		Job, err := movieService.SearchMovie(context.Background(), "cen")

		assert.Error(t, err)
		assert.Equal(t, Job, []movies.Movie{})

		// movieRepository.AssertExpectations(t)
	})
}

func TestFilterGenre(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("FilterGenre",
			mock.Anything,
			mock.AnythingOfType("string")).Return(moviesDomain, nil).Once()

		Job, err := movieService.FilterGenre(context.Background(), "cen")

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		movieRepository.On("FilterGenre",
			mock.Anything,
			mock.AnythingOfType("string")).Return([]movies.Movie{}, errors.New(mock.Anything)).Once()

		Job, err := movieService.FilterGenre(context.Background(), "cen")

		assert.Error(t, err)
		assert.Equal(t, Job, []movies.Movie{})

		// movieRepository.AssertExpectations(t)
	})
}

func TestFilterOrder(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		movieRepository.On("FilterOrder",
			mock.Anything,
			mock.AnythingOfType("string")).Return(moviesDomain, nil).Once()

		Job, err := movieService.FilterOrder(context.Background(), "cen")

		assert.NoError(t, err)
		assert.NotNil(t, Job)

		// movieRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		movieRepository.On("FilterOrder",
			mock.Anything,
			mock.AnythingOfType("string")).Return([]movies.Movie{}, errors.New(mock.Anything)).Once()

		Job, err := movieService.FilterOrder(context.Background(), "cen")

		assert.Error(t, err)
		assert.Equal(t, Job, []movies.Movie{})

		// movieRepository.AssertExpectations(t)
	})
}
