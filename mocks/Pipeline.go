// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	pipeline "github.com/scrapnode/scrapcore/pipeline"
	mock "github.com/stretchr/testify/mock"
)

// Pipeline is an autogenerated mock type for the Pipeline type
type Pipeline struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *Pipeline) Execute(_a0 pipeline.Pipe) pipeline.Pipe {
	ret := _m.Called(_a0)

	var r0 pipeline.Pipe
	if rf, ok := ret.Get(0).(func(pipeline.Pipe) pipeline.Pipe); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.Pipe)
		}
	}

	return r0
}

type mockConstructorTestingTNewPipeline interface {
	mock.TestingT
	Cleanup(func())
}

// NewPipeline creates a new instance of Pipeline. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPipeline(t mockConstructorTestingTNewPipeline) *Pipeline {
	mock := &Pipeline{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
