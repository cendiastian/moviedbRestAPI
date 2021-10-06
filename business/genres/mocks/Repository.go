// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	genres "project/business/genres"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FirstOrCreate provides a mock function with given fields: ctx, name
func (_m *Repository) FirstOrCreate(ctx context.Context, name string) (genres.Genre, error) {
	ret := _m.Called(ctx, name)

	var r0 genres.Genre
	if rf, ok := ret.Get(0).(func(context.Context, string) genres.Genre); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(genres.Genre)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllGenre provides a mock function with given fields: ctx
func (_m *Repository) GetAllGenre(ctx context.Context) ([]genres.Genre, error) {
	ret := _m.Called(ctx)

	var r0 []genres.Genre
	if rf, ok := ret.Get(0).(func(context.Context) []genres.Genre); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]genres.Genre)
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
