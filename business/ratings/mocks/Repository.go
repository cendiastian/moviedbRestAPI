// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	ratings "project/business/ratings"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, Ratings
func (_m *Repository) Create(ctx context.Context, Ratings ratings.Ratings) (ratings.Ratings, error) {
	ret := _m.Called(ctx, Ratings)

	var r0 ratings.Ratings
	if rf, ok := ret.Get(0).(func(context.Context, ratings.Ratings) ratings.Ratings); ok {
		r0 = rf(ctx, Ratings)
	} else {
		r0 = ret.Get(0).(ratings.Ratings)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ratings.Ratings) error); ok {
		r1 = rf(ctx, Ratings)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, MovieId, UserId
func (_m *Repository) Delete(ctx context.Context, MovieId int, UserId int) error {
	ret := _m.Called(ctx, MovieId, UserId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, MovieId, UserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Detail provides a mock function with given fields: ctx, movie, user
func (_m *Repository) Detail(ctx context.Context, movie int, user int) (ratings.Ratings, error) {
	ret := _m.Called(ctx, movie, user)

	var r0 ratings.Ratings
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ratings.Ratings); ok {
		r0 = rf(ctx, movie, user)
	} else {
		r0 = ret.Get(0).(ratings.Ratings)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, movie, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, Ratings
func (_m *Repository) Update(ctx context.Context, Ratings ratings.Ratings) error {
	ret := _m.Called(ctx, Ratings)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ratings.Ratings) error); ok {
		r0 = rf(ctx, Ratings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
