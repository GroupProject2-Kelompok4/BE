// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	class "github.com/GroupProject2-Kelompok4/BE/features/class"
	mock "github.com/stretchr/testify/mock"
)

// ClassData is an autogenerated mock type for the ClassData type
type ClassData struct {
	mock.Mock
}

// DeleteClass provides a mock function with given fields: classId
func (_m *ClassData) DeleteClass(classId string) error {
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
func (_m *ClassData) GetClass(classId string) (class.ClassCore, error) {
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
func (_m *ClassData) ListClasses(limit int, offset int) ([]class.ClassCore, uint, error) {
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
func (_m *ClassData) RegisterClass(request class.ClassCore) (class.ClassCore, string, error) {
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

type mockConstructorTestingTNewClassData interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassData creates a new instance of ClassData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassData(t mockConstructorTestingTNewClassData) *ClassData {
	mock := &ClassData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
