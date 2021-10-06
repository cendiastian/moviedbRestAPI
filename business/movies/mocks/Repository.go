// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	genres "project/business/genres"

	mock "github.com/stretchr/testify/mock"

	movies "project/business/movies"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateMovie provides a mock function with given fields: ctx, Movie, array
func (_m *Repository) CreateMovie(ctx context.Context, Movie movies.Movie, array []genres.Genre) (movies.Movie, error) {
	ret := _m.Called(ctx, Movie, array)

	var r0 movies.Movie
	if rf, ok := ret.Get(0).(func(context.Context, movies.Movie, []genres.Genre) movies.Movie); ok {
		r0 = rf(ctx, Movie, array)
	} else {
		r0 = ret.Get(0).(movies.Movie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, movies.Movie, []genres.Genre) error); ok {
		r1 = rf(ctx, Movie, array)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAll provides a mock function with given fields: ctx
func (_m *Repository) DeleteAll(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMovie provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteMovie(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterGenre provides a mock function with given fields: ctx, genre
func (_m *Repository) FilterGenre(ctx context.Context, genre string) ([]movies.Movie, error) {
	ret := _m.Called(ctx, genre)

	var r0 []movies.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string) []movies.Movie); ok {
		r0 = rf(ctx, genre)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, genre)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOrder provides a mock function with given fields: ctx, order
func (_m *Repository) FilterOrder(ctx context.Context, order string) ([]movies.Movie, error) {
	ret := _m.Called(ctx, order)

	var r0 []movies.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string) []movies.Movie); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllMovie provides a mock function with given fields: ctx
func (_m *Repository) GetAllMovie(ctx context.Context) ([]movies.Movie, error) {
	ret := _m.Called(ctx)

	var r0 []movies.Movie
	if rf, ok := ret.Get(0).(func(context.Context) []movies.Movie); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MovieDetail provides a mock function with given fields: ctx, id
func (_m *Repository) MovieDetail(ctx context.Context, id int) (movies.Movie, error) {
	ret := _m.Called(ctx, id)

	var r0 movies.Movie
	if rf, ok := ret.Get(0).(func(context.Context, int) movies.Movie); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(movies.Movie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchMovie provides a mock function with given fields: ctx, title
func (_m *Repository) SearchMovie(ctx context.Context, title string) ([]movies.Movie, error) {
	ret := _m.Called(ctx, title)

	var r0 []movies.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string) []movies.Movie); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMovie provides a mock function with given fields: ctx, Movie
func (_m *Repository) UpdateMovie(ctx context.Context, Movie movies.Movie) error {
	ret := _m.Called(ctx, Movie)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, movies.Movie) error); ok {
		r0 = rf(ctx, Movie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}