// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// FeedbackHandler is an autogenerated mock type for the FeedbackHandler type
type FeedbackHandler struct {
	mock.Mock
}

// RegisterFeedbackMentee provides a mock function with given fields:
func (_m *FeedbackHandler) RegisterFeedbackMentee() echo.HandlerFunc {
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

type mockConstructorTestingTNewFeedbackHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeedbackHandler creates a new instance of FeedbackHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeedbackHandler(t mockConstructorTestingTNewFeedbackHandler) *FeedbackHandler {
	mock := &FeedbackHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}