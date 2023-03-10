// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	msgbus "github.com/scrapnode/scrapcore/msgbus"
	mock "github.com/stretchr/testify/mock"
)

// SubscribeFn is an autogenerated mock type for the SubscribeFn type
type SubscribeFn struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, event
func (_m *SubscribeFn) Execute(ctx context.Context, event *msgbus.Event) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *msgbus.Event) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSubscribeFn interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubscribeFn creates a new instance of SubscribeFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubscribeFn(t mockConstructorTestingTNewSubscribeFn) *SubscribeFn {
	mock := &SubscribeFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
