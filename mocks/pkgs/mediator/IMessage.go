// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IMessage is an autogenerated mock type for the IMessage type
type IMessage struct {
	mock.Mock
}

type IMessage_Expecter struct {
	mock *mock.Mock
}

func (_m *IMessage) EXPECT() *IMessage_Expecter {
	return &IMessage_Expecter{mock: &_m.Mock}
}

// GetMessage provides a mock function with given fields:
func (_m *IMessage) GetMessage() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMessage")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// IMessage_GetMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessage'
type IMessage_GetMessage_Call struct {
	*mock.Call
}

// GetMessage is a helper method to define mock.On call
func (_e *IMessage_Expecter) GetMessage() *IMessage_GetMessage_Call {
	return &IMessage_GetMessage_Call{Call: _e.mock.On("GetMessage")}
}

func (_c *IMessage_GetMessage_Call) Run(run func()) *IMessage_GetMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessage_GetMessage_Call) Return(_a0 interface{}) *IMessage_GetMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessage_GetMessage_Call) RunAndReturn(run func() interface{}) *IMessage_GetMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessageString provides a mock function with given fields:
func (_m *IMessage) GetMessageString() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMessageString")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IMessage_GetMessageString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessageString'
type IMessage_GetMessageString_Call struct {
	*mock.Call
}

// GetMessageString is a helper method to define mock.On call
func (_e *IMessage_Expecter) GetMessageString() *IMessage_GetMessageString_Call {
	return &IMessage_GetMessageString_Call{Call: _e.mock.On("GetMessageString")}
}

func (_c *IMessage_GetMessageString_Call) Run(run func()) *IMessage_GetMessageString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessage_GetMessageString_Call) Return(_a0 string) *IMessage_GetMessageString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessage_GetMessageString_Call) RunAndReturn(run func() string) *IMessage_GetMessageString_Call {
	_c.Call.Return(run)
	return _c
}

// NewIMessage creates a new instance of IMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMessage(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMessage {
	mock := &IMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
