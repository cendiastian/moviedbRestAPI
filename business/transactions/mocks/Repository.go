// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	transactions "project/business/transactions"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: ctx, Transaction
func (_m *Repository) CreateTransaction(ctx context.Context, Transaction transactions.Transaction) (transactions.Transaction, error) {
	ret := _m.Called(ctx, Transaction)

	var r0 transactions.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, transactions.Transaction) transactions.Transaction); ok {
		r0 = rf(ctx, Transaction)
	} else {
		r0 = ret.Get(0).(transactions.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, transactions.Transaction) error); ok {
		r1 = rf(ctx, Transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetailTrans provides a mock function with given fields: ctx, id
func (_m *Repository) DetailTrans(ctx context.Context, id int) (transactions.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 transactions.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, int) transactions.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(transactions.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
