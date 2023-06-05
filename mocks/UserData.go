// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	user "github.com/GroupProject2-Kelompok4/BE/features/user"
	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// DeactiveUser provides a mock function with given fields: userId
func (_m *UserData) DeactiveUser(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: request
func (_m *UserData) Login(request user.UserCore) (user.UserCore, string, error) {
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

// ProfileUser provides a mock function with given fields: userId
func (_m *UserData) ProfileUser(userId string) (user.UserCore, error) {
	ret := _m.Called(userId)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.UserCore, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) user.UserCore); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: request
func (_m *UserData) Register(request user.UserCore) (user.UserCore, error) {
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

// SearchUser provides a mock function with given fields: keyword, limit, offset
func (_m *UserData) SearchUser(keyword string, limit int, offset int) ([]user.UserCore, uint, error) {
	ret := _m.Called(keyword, limit, offset)

	var r0 []user.UserCore
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]user.UserCore, uint, error)); ok {
		return rf(keyword, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []user.UserCore); ok {
		r0 = rf(keyword, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.UserCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) uint); ok {
		r1 = rf(keyword, limit, offset)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(string, int, int) error); ok {
		r2 = rf(keyword, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateProfile provides a mock function with given fields: userId, request
func (_m *UserData) UpdateProfile(userId string, request user.UserCore) (user.UserCore, error) {
	ret := _m.Called(userId, request)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, user.UserCore) (user.UserCore, error)); ok {
		return rf(userId, request)
	}
	if rf, ok := ret.Get(0).(func(string, user.UserCore) user.UserCore); ok {
		r0 = rf(userId, request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(string, user.UserCore) error); ok {
		r1 = rf(userId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserProfile provides a mock function with given fields: userId, request
func (_m *UserData) UpdateUserProfile(userId string, request user.UserCore) (user.UserCore, error) {
	ret := _m.Called(userId, request)

	var r0 user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, user.UserCore) (user.UserCore, error)); ok {
		return rf(userId, request)
	}
	if rf, ok := ret.Get(0).(func(string, user.UserCore) user.UserCore); ok {
		r0 = rf(userId, request)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(string, user.UserCore) error); ok {
		r1 = rf(userId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
