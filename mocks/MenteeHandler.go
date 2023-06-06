// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MenteeHandler is an autogenerated mock type for the MenteeHandler type
type MenteeHandler struct {
	mock.Mock
}

// RegisterMentee provides a mock function with given fields:
func (_m *MenteeHandler) RegisterMentee() echo.HandlerFunc {
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

type mockConstructorTestingTNewMenteeHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenteeHandler creates a new instance of MenteeHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenteeHandler(t mockConstructorTestingTNewMenteeHandler) *MenteeHandler {
	mock := &MenteeHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
