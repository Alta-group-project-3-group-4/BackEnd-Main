// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	comment "airbnb/feature/comment"

	mock "github.com/stretchr/testify/mock"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// Add provides a mock function with given fields: newComment
func (_m *CommentService) Add(newComment comment.Core) error {
	ret := _m.Called(newComment)

	var r0 error
	if rf, ok := ret.Get(0).(func(comment.Core) error); ok {
		r0 = rf(newComment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: userId, id
func (_m *CommentService) Delete(userId uint, id uint) error {
	ret := _m.Called(userId, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userId, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHomestayComments provides a mock function with given fields: homeId
func (_m *CommentService) GetHomestayComments(homeId uint) ([]comment.Core, error) {
	ret := _m.Called(homeId)

	var r0 []comment.Core
	if rf, ok := ret.Get(0).(func(uint) []comment.Core); ok {
		r0 = rf(homeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comment.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(homeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentService(t mockConstructorTestingTNewCommentService) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
