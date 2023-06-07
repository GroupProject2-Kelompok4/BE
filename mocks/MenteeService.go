// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee"
	mock "github.com/stretchr/testify/mock"
)

// MenteeService is an autogenerated mock type for the MenteeService type
type MenteeService struct {
	mock.Mock
}

// ProfileMenteeAndFeedback provides a mock function with given fields: menteeId
func (_m *MenteeService) ProfileMenteeAndFeedback(menteeId string) (mentee.MenteeCore, error) {
	ret := _m.Called(menteeId)

	var r0 mentee.MenteeCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (mentee.MenteeCore, error)); ok {
		return rf(menteeId)
	}
	if rf, ok := ret.Get(0).(func(string) mentee.MenteeCore); ok {
		r0 = rf(menteeId)
	} else {
		r0 = ret.Get(0).(mentee.MenteeCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(menteeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterMentee provides a mock function with given fields: request
func (_m *MenteeService) RegisterMentee(request mentee.MenteeCore) (mentee.MenteeCore, error) {
	ret := _m.Called(request)

	var r0 mentee.MenteeCore
	var r1 error
	if rf, ok := ret.Get(0).(func(mentee.MenteeCore) (mentee.MenteeCore, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(mentee.MenteeCore) mentee.MenteeCore); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(mentee.MenteeCore)
	}

	if rf, ok := ret.Get(1).(func(mentee.MenteeCore) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchMentee provides a mock function with given fields: keyword, limit, offset
func (_m *MenteeService) SearchMentee(keyword string, limit int, offset int) ([]mentee.MenteeCore, uint, error) {
	ret := _m.Called(keyword, limit, offset)

	var r0 []mentee.MenteeCore
	var r1 uint
	var r2 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]mentee.MenteeCore, uint, error)); ok {
		return rf(keyword, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []mentee.MenteeCore); ok {
		r0 = rf(keyword, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee.MenteeCore)
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

type mockConstructorTestingTNewMenteeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenteeService creates a new instance of MenteeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenteeService(t mockConstructorTestingTNewMenteeService) *MenteeService {
	mock := &MenteeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
