// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IConsumer is an autogenerated mock type for the IConsumer type
type IConsumer struct {
	mock.Mock
}

type IConsumer_Expecter struct {
	mock *mock.Mock
}

func (_m *IConsumer) EXPECT() *IConsumer_Expecter {
	return &IConsumer_Expecter{mock: &_m.Mock}
}

// SetUpConsumer provides a mock function with given fields:
func (_m *IConsumer) SetUpConsumer() {
	_m.Called()
}

// IConsumer_SetUpConsumer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUpConsumer'
type IConsumer_SetUpConsumer_Call struct {
	*mock.Call
}

// SetUpConsumer is a helper method to define mock.On call
func (_e *IConsumer_Expecter) SetUpConsumer() *IConsumer_SetUpConsumer_Call {
	return &IConsumer_SetUpConsumer_Call{Call: _e.mock.On("SetUpConsumer")}
}

func (_c *IConsumer_SetUpConsumer_Call) Run(run func()) *IConsumer_SetUpConsumer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IConsumer_SetUpConsumer_Call) Return() *IConsumer_SetUpConsumer_Call {
	_c.Call.Return()
	return _c
}

func (_c *IConsumer_SetUpConsumer_Call) RunAndReturn(run func()) *IConsumer_SetUpConsumer_Call {
	_c.Call.Return(run)
	return _c
}

// NewIConsumer creates a new instance of IConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *IConsumer {
	mock := &IConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}