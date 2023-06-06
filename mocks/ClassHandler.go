// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// ClassHandler is an autogenerated mock type for the ClassHandler type
type ClassHandler struct {
	mock.Mock
}

// ListClasses provides a mock function with given fields:
func (_m *ClassHandler) ListClasses() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// RegisterClass provides a mock function with given fields:
func (_m *ClassHandler) RegisterClass() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewClassHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassHandler creates a new instance of ClassHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassHandler(t mockConstructorTestingTNewClassHandler) *ClassHandler {
	mock := &ClassHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
