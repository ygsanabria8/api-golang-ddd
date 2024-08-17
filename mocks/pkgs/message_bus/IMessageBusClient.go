// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	utils "api.ddd/pkgs/message_bus/utils"
	mock "github.com/stretchr/testify/mock"
)

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

// SendMessage provides a mock function with given fields: eventType
func (_m *IMessageBusClient) SendMessage(eventType *utils.Event) {
	_m.Called(eventType)
}

// IMessageBusClient_SendMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessage'
type IMessageBusClient_SendMessage_Call struct {
	*mock.Call
}

// SendMessage is a helper method to define mock.On call
//   - eventType *utils.Event
func (_e *IMessageBusClient_Expecter) SendMessage(eventType interface{}) *IMessageBusClient_SendMessage_Call {
	return &IMessageBusClient_SendMessage_Call{Call: _e.mock.On("SendMessage", eventType)}
}

func (_c *IMessageBusClient_SendMessage_Call) Run(run func(eventType *utils.Event)) *IMessageBusClient_SendMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*utils.Event))
	})
	return _c
}

func (_c *IMessageBusClient_SendMessage_Call) Return() *IMessageBusClient_SendMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *IMessageBusClient_SendMessage_Call) RunAndReturn(run func(*utils.Event)) *IMessageBusClient_SendMessage_Call {
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
