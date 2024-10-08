// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	utils "api.ddd/pkgs/message_bus/utils"
	mock "github.com/stretchr/testify/mock"
)

// IEventHandler is an autogenerated mock type for the IEventHandler type
type IEventHandler struct {
	mock.Mock
}

type IEventHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *IEventHandler) EXPECT() *IEventHandler_Expecter {
	return &IEventHandler_Expecter{mock: &_m.Mock}
}

// OnMessage provides a mock function with given fields: event
func (_m *IEventHandler) OnMessage(event utils.IEvent) error {
	ret := _m.Called(event)

	if len(ret) == 0 {
		panic("no return value specified for OnMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(utils.IEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IEventHandler_OnMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnMessage'
type IEventHandler_OnMessage_Call struct {
	*mock.Call
}

// OnMessage is a helper method to define mock.On call
//   - event utils.IEvent
func (_e *IEventHandler_Expecter) OnMessage(event interface{}) *IEventHandler_OnMessage_Call {
	return &IEventHandler_OnMessage_Call{Call: _e.mock.On("OnMessage", event)}
}

func (_c *IEventHandler_OnMessage_Call) Run(run func(event utils.IEvent)) *IEventHandler_OnMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(utils.IEvent))
	})
	return _c
}

func (_c *IEventHandler_OnMessage_Call) Return(_a0 error) *IEventHandler_OnMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IEventHandler_OnMessage_Call) RunAndReturn(run func(utils.IEvent) error) *IEventHandler_OnMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewIEventHandler creates a new instance of IEventHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIEventHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IEventHandler {
	mock := &IEventHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
