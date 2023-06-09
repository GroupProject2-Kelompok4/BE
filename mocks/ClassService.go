// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	class "github.com/GroupProject2-Kelompok4/BE/features/class"
	mock "github.com/stretchr/testify/mock"
)

// ClassService is an autogenerated mock type for the ClassService type
type ClassService struct {
	mock.Mock
}

// DeleteClass provides a mock function with given fields: classId
func (_m *ClassService) DeleteClass(classId string) error {
	ret := _m.Called(classId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(classId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetClass provides a mock function with given fields: classId
func (_m *ClassService) GetClass(classId string) (class.ClassCore, error) {
	ret := _m.Called(classId)

	var r0 class.ClassCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (class.ClassCore, error)); ok {
		return rf(classId)
	}
	if rf, ok := ret.Get(0).(func(string) class.ClassCore); ok {
		r0 = rf(classId)
	} else {
		r0 = ret.Get(0).(class.ClassCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(classId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClasses provides a mock function with given fields: limit, offset
func (_m *ClassService) ListClasses(limit int, offset int) ([]class.ClassCore, uint, error) {
	ret := _m.Called(limit, offset)

	var r0 []class.ClassCore
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]class.ClassCore, uint, error)); ok {
		return rf(limit, offset)
	}
	if rf, ok := ret.Get(0).(func(int, int) []class.ClassCore); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class.ClassCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) uint); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Get(1).(uint)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterClass provides a mock function with given fields: request
func (_m *ClassService) RegisterClass(request class.ClassCore) (class.ClassCore, string, error) {
	ret := _m.Called(request)

	var r0 class.ClassCore
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(class.ClassCore) (class.ClassCore, string, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(class.ClassCore) class.ClassCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(class.ClassCore)
	}

	if rf, ok := ret.Get(1).(func(class.ClassCore) string); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(class.ClassCore) error); ok {
		r2 = rf(request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateClass provides a mock function with given fields: classId, request
func (_m *ClassService) UpdateClass(classId string, request class.ClassCore) error {
	ret := _m.Called(classId, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, class.ClassCore) error); ok {
		r0 = rf(classId, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClassService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassService creates a new instance of ClassService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassService(t mockConstructorTestingTNewClassService) *ClassService {
	mock := &ClassService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
