// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	message_bus "api.ddd/pkgs/message_bus"
	consumer "api.ddd/pkgs/message_bus/consumer"

	mock "github.com/stretchr/testify/mock"
)

// IMessageBus is an autogenerated mock type for the IMessageBus type
type IMessageBus struct {
	mock.Mock
}

type IMessageBus_Expecter struct {
	mock *mock.Mock
}

func (_m *IMessageBus) EXPECT() *IMessageBus_Expecter {
	return &IMessageBus_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with given fields:
func (_m *IMessageBus) Build() message_bus.IMessageBusClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 message_bus.IMessageBusClient
	if rf, ok := ret.Get(0).(func() message_bus.IMessageBusClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message_bus.IMessageBusClient)
		}
	}

	return r0
}

// IMessageBus_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type IMessageBus_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *IMessageBus_Expecter) Build() *IMessageBus_Build_Call {
	return &IMessageBus_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *IMessageBus_Build_Call) Run(run func()) *IMessageBus_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IMessageBus_Build_Call) Return(_a0 message_bus.IMessageBusClient) *IMessageBus_Build_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessageBus_Build_Call) RunAndReturn(run func() message_bus.IMessageBusClient) *IMessageBus_Build_Call {
	_c.Call.Return(run)
	return _c
}

// WitProducer provides a mock function with given fields: topic, eventType
func (_m *IMessageBus) WitProducer(topic string, eventType interface{}) message_bus.IMessageBus {
	ret := _m.Called(topic, eventType)

	if len(ret) == 0 {
		panic("no return value specified for WitProducer")
	}

	var r0 message_bus.IMessageBus
	if rf, ok := ret.Get(0).(func(string, interface{}) message_bus.IMessageBus); ok {
		r0 = rf(topic, eventType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message_bus.IMessageBus)
		}
	}

	return r0
}

// IMessageBus_WitProducer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WitProducer'
type IMessageBus_WitProducer_Call struct {
	*mock.Call
}

// WitProducer is a helper method to define mock.On call
//   - topic string
//   - eventType interface{}
func (_e *IMessageBus_Expecter) WitProducer(topic interface{}, eventType interface{}) *IMessageBus_WitProducer_Call {
	return &IMessageBus_WitProducer_Call{Call: _e.mock.On("WitProducer", topic, eventType)}
}

func (_c *IMessageBus_WitProducer_Call) Run(run func(topic string, eventType interface{})) *IMessageBus_WitProducer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *IMessageBus_WitProducer_Call) Return(_a0 message_bus.IMessageBus) *IMessageBus_WitProducer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessageBus_WitProducer_Call) RunAndReturn(run func(string, interface{}) message_bus.IMessageBus) *IMessageBus_WitProducer_Call {
	_c.Call.Return(run)
	return _c
}

// WithAppName provides a mock function with given fields: appName
func (_m *IMessageBus) WithAppName(appName string) message_bus.IMessageBus {
	ret := _m.Called(appName)

	if len(ret) == 0 {
		panic("no return value specified for WithAppName")
	}

	var r0 message_bus.IMessageBus
	if rf, ok := ret.Get(0).(func(string) message_bus.IMessageBus); ok {
		r0 = rf(appName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message_bus.IMessageBus)
		}
	}

	return r0
}

// IMessageBus_WithAppName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithAppName'
type IMessageBus_WithAppName_Call struct {
	*mock.Call
}

// WithAppName is a helper method to define mock.On call
//   - appName string
func (_e *IMessageBus_Expecter) WithAppName(appName interface{}) *IMessageBus_WithAppName_Call {
	return &IMessageBus_WithAppName_Call{Call: _e.mock.On("WithAppName", appName)}
}

func (_c *IMessageBus_WithAppName_Call) Run(run func(appName string)) *IMessageBus_WithAppName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IMessageBus_WithAppName_Call) Return(_a0 message_bus.IMessageBus) *IMessageBus_WithAppName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessageBus_WithAppName_Call) RunAndReturn(run func(string) message_bus.IMessageBus) *IMessageBus_WithAppName_Call {
	_c.Call.Return(run)
	return _c
}

// WithConsumer provides a mock function with given fields: topic, handler
func (_m *IMessageBus) WithConsumer(topic string, handler consumer.IEventHandler) message_bus.IMessageBus {
	ret := _m.Called(topic, handler)

	if len(ret) == 0 {
		panic("no return value specified for WithConsumer")
	}

	var r0 message_bus.IMessageBus
	if rf, ok := ret.Get(0).(func(string, consumer.IEventHandler) message_bus.IMessageBus); ok {
		r0 = rf(topic, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message_bus.IMessageBus)
		}
	}

	return r0
}

// IMessageBus_WithConsumer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithConsumer'
type IMessageBus_WithConsumer_Call struct {
	*mock.Call
}

// WithConsumer is a helper method to define mock.On call
//   - topic string
//   - handler consumer.IEventHandler
func (_e *IMessageBus_Expecter) WithConsumer(topic interface{}, handler interface{}) *IMessageBus_WithConsumer_Call {
	return &IMessageBus_WithConsumer_Call{Call: _e.mock.On("WithConsumer", topic, handler)}
}

func (_c *IMessageBus_WithConsumer_Call) Run(run func(topic string, handler consumer.IEventHandler)) *IMessageBus_WithConsumer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(consumer.IEventHandler))
	})
	return _c
}

func (_c *IMessageBus_WithConsumer_Call) Return(_a0 message_bus.IMessageBus) *IMessageBus_WithConsumer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IMessageBus_WithConsumer_Call) RunAndReturn(run func(string, consumer.IEventHandler) message_bus.IMessageBus) *IMessageBus_WithConsumer_Call {
	_c.Call.Return(run)
	return _c
}

// NewIMessageBus creates a new instance of IMessageBus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMessageBus(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMessageBus {
	mock := &IMessageBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
