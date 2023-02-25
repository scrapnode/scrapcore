// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Transport is an autogenerated mock type for the Transport type
type Transport struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx
func (_m *Transport) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *Transport) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: ctx
func (_m *Transport) Stop(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransport interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransport creates a new instance of Transport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransport(t mockConstructorTestingTNewTransport) *Transport {
	mock := &Transport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
