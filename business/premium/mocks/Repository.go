// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	premium "project/business/premium"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Detail provides a mock function with given fields: ctx, user
func (_m *Repository) Detail(ctx context.Context, user int) (premium.Premium, error) {
	ret := _m.Called(ctx, user)

	var r0 premium.Premium
	if rf, ok := ret.Get(0).(func(context.Context, int) premium.Premium); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(premium.Premium)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, Premium
func (_m *Repository) Save(ctx context.Context, Premium premium.Premium) (premium.Premium, error) {
	ret := _m.Called(ctx, Premium)

	var r0 premium.Premium
	if rf, ok := ret.Get(0).(func(context.Context, premium.Premium) premium.Premium); ok {
		r0 = rf(ctx, Premium)
	} else {
		r0 = ret.Get(0).(premium.Premium)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, premium.Premium) error); ok {
		r1 = rf(ctx, Premium)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}