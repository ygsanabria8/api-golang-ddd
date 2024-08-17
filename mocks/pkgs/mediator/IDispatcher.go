// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mediator "api.ddd/pkgs/mediator"
	mock "github.com/stretchr/testify/mock"
)

// IDispatcher is an autogenerated mock type for the IDispatcher type
type IDispatcher struct {
	mock.Mock
}

type IDispatcher_Expecter struct {
	mock *mock.Mock
}

func (_m *IDispatcher) EXPECT() *IDispatcher_Expecter {
	return &IDispatcher_Expecter{mock: &_m.Mock}
}

// RegisterHandler provides a mock function with given fields: _a0, _a1
func (_m *IDispatcher) RegisterHandler(_a0 mediator.IHandler, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RegisterHandler")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(mediator.IHandler, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IDispatcher_RegisterHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterHandler'
type IDispatcher_RegisterHandler_Call struct {
	*mock.Call
}

// RegisterHandler is a helper method to define mock.On call
//   - _a0 mediator.IHandler
//   - _a1 interface{}
func (_e *IDispatcher_Expecter) RegisterHandler(_a0 interface{}, _a1 interface{}) *IDispatcher_RegisterHandler_Call {
	return &IDispatcher_RegisterHandler_Call{Call: _e.mock.On("RegisterHandler", _a0, _a1)}
}

func (_c *IDispatcher_RegisterHandler_Call) Run(run func(_a0 mediator.IHandler, _a1 interface{})) *IDispatcher_RegisterHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(mediator.IHandler), args[1].(interface{}))
	})
	return _c
}

func (_c *IDispatcher_RegisterHandler_Call) Return(_a0 error) *IDispatcher_RegisterHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IDispatcher_RegisterHandler_Call) RunAndReturn(run func(mediator.IHandler, interface{}) error) *IDispatcher_RegisterHandler_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *IDispatcher) Send(_a0 *mediator.Message) (interface{}, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*mediator.Message) (interface{}, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*mediator.Message) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*mediator.Message) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IDispatcher_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type IDispatcher_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *mediator.Message
func (_e *IDispatcher_Expecter) Send(_a0 interface{}) *IDispatcher_Send_Call {
	return &IDispatcher_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *IDispatcher_Send_Call) Run(run func(_a0 *mediator.Message)) *IDispatcher_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*mediator.Message))
	})
	return _c
}

func (_c *IDispatcher_Send_Call) Return(_a0 interface{}, _a1 error) *IDispatcher_Send_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IDispatcher_Send_Call) RunAndReturn(run func(*mediator.Message) (interface{}, error)) *IDispatcher_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewIDispatcher creates a new instance of IDispatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDispatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDispatcher {
	mock := &IDispatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
