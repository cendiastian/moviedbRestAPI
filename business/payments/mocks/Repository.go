// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	payments "project/business/payments"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int) (payments.Payment_method, error) {
	ret := _m.Called(ctx, id)

	var r0 payments.Payment_method
	if rf, ok := ret.Get(0).(func(context.Context, int) payments.Payment_method); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(payments.Payment_method)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Detail provides a mock function with given fields: ctx, id
func (_m *Repository) Detail(ctx context.Context, id int) (payments.Payment_method, error) {
	ret := _m.Called(ctx, id)

	var r0 payments.Payment_method
	if rf, ok := ret.Get(0).(func(context.Context, int) payments.Payment_method); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(payments.Payment_method)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]payments.Payment_method, error) {
	ret := _m.Called(ctx)

	var r0 []payments.Payment_method
	if rf, ok := ret.Get(0).(func(context.Context) []payments.Payment_method); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payments.Payment_method)
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

// Register provides a mock function with given fields: ctx, Payment_method
func (_m *Repository) Register(ctx context.Context, Payment_method payments.Payment_method) (payments.Payment_method, error) {
	ret := _m.Called(ctx, Payment_method)

	var r0 payments.Payment_method
	if rf, ok := ret.Get(0).(func(context.Context, payments.Payment_method) payments.Payment_method); ok {
		r0 = rf(ctx, Payment_method)
	} else {
		r0 = ret.Get(0).(payments.Payment_method)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, payments.Payment_method) error); ok {
		r1 = rf(ctx, Payment_method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, Payment_method
func (_m *Repository) Update(ctx context.Context, Payment_method payments.Payment_method) (payments.Payment_method, error) {
	ret := _m.Called(ctx, Payment_method)

	var r0 payments.Payment_method
	if rf, ok := ret.Get(0).(func(context.Context, payments.Payment_method) payments.Payment_method); ok {
		r0 = rf(ctx, Payment_method)
	} else {
		r0 = ret.Get(0).(payments.Payment_method)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, payments.Payment_method) error); ok {
		r1 = rf(ctx, Payment_method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
