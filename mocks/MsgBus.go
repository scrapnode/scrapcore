// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	msgbus "github.com/scrapnode/scrapcore/msgbus"
	mock "github.com/stretchr/testify/mock"
)

// MsgBus is an autogenerated mock type for the MsgBus type
type MsgBus struct {
	mock.Mock
}

// Connect provides a mock function with given fields: ctx
func (_m *MsgBus) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields: ctx
func (_m *MsgBus) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Pub provides a mock function with given fields: ctx, event
func (_m *MsgBus) Pub(ctx context.Context, event *msgbus.Event) (*msgbus.PubRes, error) {
	ret := _m.Called(ctx, event)

	var r0 *msgbus.PubRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *msgbus.Event) (*msgbus.PubRes, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *msgbus.Event) *msgbus.PubRes); ok {
		r0 = rf(ctx, event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msgbus.PubRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *msgbus.Event) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sub provides a mock function with given fields: ctx, sample, queue, fn
func (_m *MsgBus) Sub(ctx context.Context, sample *msgbus.Event, queue string, fn msgbus.SubscribeFn) (func() error, error) {
	ret := _m.Called(ctx, sample, queue, fn)

	var r0 func() error
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *msgbus.Event, string, msgbus.SubscribeFn) (func() error, error)); ok {
		return rf(ctx, sample, queue, fn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *msgbus.Event, string, msgbus.SubscribeFn) func() error); ok {
		r0 = rf(ctx, sample, queue, fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *msgbus.Event, string, msgbus.SubscribeFn) error); ok {
		r1 = rf(ctx, sample, queue, fn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMsgBus interface {
	mock.TestingT
	Cleanup(func())
}

// NewMsgBus creates a new instance of MsgBus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMsgBus(t mockConstructorTestingTNewMsgBus) *MsgBus {
	mock := &MsgBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}