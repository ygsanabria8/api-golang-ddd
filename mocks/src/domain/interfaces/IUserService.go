// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	aggregates "api.ddd/src/domain/aggregates"

	mock "github.com/stretchr/testify/mock"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

type IUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *IUserService) EXPECT() *IUserService_Expecter {
	return &IUserService_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *IUserService) CreateUser(user *aggregates.User) (*aggregates.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *aggregates.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*aggregates.User) (*aggregates.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*aggregates.User) *aggregates.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregates.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*aggregates.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUserService_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type IUserService_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user *aggregates.User
func (_e *IUserService_Expecter) CreateUser(user interface{}) *IUserService_CreateUser_Call {
	return &IUserService_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *IUserService_CreateUser_Call) Run(run func(user *aggregates.User)) *IUserService_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*aggregates.User))
	})
	return _c
}

func (_c *IUserService_CreateUser_Call) Return(_a0 *aggregates.User, _a1 error) *IUserService_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IUserService_CreateUser_Call) RunAndReturn(run func(*aggregates.User) (*aggregates.User, error)) *IUserService_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: userId
func (_m *IUserService) DeleteUser(userId string) error {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUserService_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type IUserService_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - userId string
func (_e *IUserService_Expecter) DeleteUser(userId interface{}) *IUserService_DeleteUser_Call {
	return &IUserService_DeleteUser_Call{Call: _e.mock.On("DeleteUser", userId)}
}

func (_c *IUserService_DeleteUser_Call) Run(run func(userId string)) *IUserService_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IUserService_DeleteUser_Call) Return(_a0 error) *IUserService_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUserService_DeleteUser_Call) RunAndReturn(run func(string) error) *IUserService_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUsers provides a mock function with given fields:
func (_m *IUserService) GetAllUsers() ([]*aggregates.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []*aggregates.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*aggregates.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*aggregates.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*aggregates.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUserService_GetAllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUsers'
type IUserService_GetAllUsers_Call struct {
	*mock.Call
}

// GetAllUsers is a helper method to define mock.On call
func (_e *IUserService_Expecter) GetAllUsers() *IUserService_GetAllUsers_Call {
	return &IUserService_GetAllUsers_Call{Call: _e.mock.On("GetAllUsers")}
}

func (_c *IUserService_GetAllUsers_Call) Run(run func()) *IUserService_GetAllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IUserService_GetAllUsers_Call) Return(_a0 []*aggregates.User, _a1 error) *IUserService_GetAllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IUserService_GetAllUsers_Call) RunAndReturn(run func() ([]*aggregates.User, error)) *IUserService_GetAllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: userId
func (_m *IUserService) GetUserById(userId string) (*aggregates.User, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 *aggregates.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*aggregates.User, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) *aggregates.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregates.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUserService_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type IUserService_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - userId string
func (_e *IUserService_Expecter) GetUserById(userId interface{}) *IUserService_GetUserById_Call {
	return &IUserService_GetUserById_Call{Call: _e.mock.On("GetUserById", userId)}
}

func (_c *IUserService_GetUserById_Call) Run(run func(userId string)) *IUserService_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IUserService_GetUserById_Call) Return(_a0 *aggregates.User, _a1 error) *IUserService_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IUserService_GetUserById_Call) RunAndReturn(run func(string) (*aggregates.User, error)) *IUserService_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: user
func (_m *IUserService) UpdateUser(user *aggregates.User) (*aggregates.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *aggregates.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*aggregates.User) (*aggregates.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*aggregates.User) *aggregates.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregates.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*aggregates.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUserService_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type IUserService_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - user *aggregates.User
func (_e *IUserService_Expecter) UpdateUser(user interface{}) *IUserService_UpdateUser_Call {
	return &IUserService_UpdateUser_Call{Call: _e.mock.On("UpdateUser", user)}
}

func (_c *IUserService_UpdateUser_Call) Run(run func(user *aggregates.User)) *IUserService_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*aggregates.User))
	})
	return _c
}

func (_c *IUserService_UpdateUser_Call) Return(_a0 *aggregates.User, _a1 error) *IUserService_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IUserService_UpdateUser_Call) RunAndReturn(run func(*aggregates.User) (*aggregates.User, error)) *IUserService_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewIUserService creates a new instance of IUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserService {
	mock := &IUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
