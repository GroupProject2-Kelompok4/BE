// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	user "github.com/GroupProject2-Kelompok4/BE/features/user"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Login provides a mock function with given fields: request
func (_m *UserService) Login(request user.UserCore) (user.UserCore, string, error) {
	ret := _m.Called(request)

	var r0 user.UserCore
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(user.UserCore) (user.UserCore, string, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(user.UserCore) string); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(user.UserCore) error); ok {
		r2 = rf(request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: request
func (_m *UserService) Register(request user.UserCore) (user.UserCore, error) {
	ret := _m.Called(request)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(user.UserCore) (user.UserCore, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
