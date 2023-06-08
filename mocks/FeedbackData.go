// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback"
	mock "github.com/stretchr/testify/mock"
)

// FeedbackData is an autogenerated mock type for the FeedbackData type
type FeedbackData struct {
	mock.Mock
}

// RegisterFeedbackMentee provides a mock function with given fields: request, userId
func (_m *FeedbackData) RegisterFeedbackMentee(request feedback.FeedbackCore, userId string) (feedback.FeedbackCore, error) {
	ret := _m.Called(request, userId)

	var r0 feedback.FeedbackCore
	var r1 error
	if rf, ok := ret.Get(0).(func(feedback.FeedbackCore, string) (feedback.FeedbackCore, error)); ok {
		return rf(request, userId)
	}
	if rf, ok := ret.Get(0).(func(feedback.FeedbackCore, string) feedback.FeedbackCore); ok {
		r0 = rf(request, userId)
	} else {
		r0 = ret.Get(0).(feedback.FeedbackCore)
	}

	if rf, ok := ret.Get(1).(func(feedback.FeedbackCore, string) error); ok {
		r1 = rf(request, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeedbackData interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeedbackData creates a new instance of FeedbackData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeedbackData(t mockConstructorTestingTNewFeedbackData) *FeedbackData {
	mock := &FeedbackData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
