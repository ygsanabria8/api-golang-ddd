// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IMessageBusClient is an autogenerated mock type for the IMessageBusClient type
type IMessageBusClient struct {
	mock.Mock
}

type IMessageBusClient_Expecter struct {
	mock *mock.Mock
}

func (_m *IMessageBusClient) EXPECT() *IMessageBusClient_Expecter {
	return &IMessageBusClient_Expecter{mock: &_m.Mock}
}

// SendMessage provides a mock function with given fields: event
func (_m *IMessageBusClient) SendMessage(event interface{}) {
	_m.Called(event)
}

// IMessageBusClient_SendMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessage'
type IMessageBusClient_SendMessage_Call struct {
	*mock.Call
}

// SendMessage is a helper method to define mock.On call
//   - event interface{}
func (_e *IMessageBusClient_Expecter) SendMessage(event interface{}) *IMessageBusClient_SendMessage_Call {
	return &IMessageBusClient_SendMessage_Call{Call: _e.mock.On("SendMessage", event)}
}

func (_c *IMessageBusClient_SendMessage_Call) Run(run func(event interface{})) *IMessageBusClient_SendMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *IMessageBusClient_SendMessage_Call) Return() *IMessageBusClient_SendMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *IMessageBusClient_SendMessage_Call) RunAndReturn(run func(interface{})) *IMessageBusClient_SendMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewIMessageBusClient creates a new instance of IMessageBusClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMessageBusClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMessageBusClient {
	mock := &IMessageBusClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
