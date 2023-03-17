// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	users "airbnb/feature/users"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceInterface is an autogenerated mock type for the UserServiceInterface type
type UserServiceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: userEntity
func (_m *UserServiceInterface) Create(userEntity users.UserEntity) (users.UserEntity, error) {
	ret := _m.Called(userEntity)

	var r0 users.UserEntity
	if rf, ok := ret.Get(0).(func(users.UserEntity) users.UserEntity); ok {
		r0 = rf(userEntity)
	} else {
		r0 = ret.Get(0).(users.UserEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserEntity) error); ok {
		r1 = rf(userEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *UserServiceInterface) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *UserServiceInterface) GetAll() ([]users.UserEntity, error) {
	ret := _m.Called()

	var r0 []users.UserEntity
	if rf, ok := ret.Get(0).(func() []users.UserEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.UserEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *UserServiceInterface) GetById(id uint) (users.UserEntity, error) {
	ret := _m.Called(id)

	var r0 users.UserEntity
	if rf, ok := ret.Get(0).(func(uint) users.UserEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.UserEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userEntity, id
func (_m *UserServiceInterface) Update(userEntity users.UserEntity, id uint) (users.UserEntity, error) {
	ret := _m.Called(userEntity, id)

	var r0 users.UserEntity
	if rf, ok := ret.Get(0).(func(users.UserEntity, uint) users.UserEntity); ok {
		r0 = rf(userEntity, id)
	} else {
		r0 = ret.Get(0).(users.UserEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserEntity, uint) error); ok {
		r1 = rf(userEntity, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServiceInterface creates a new instance of UserServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServiceInterface(t mockConstructorTestingTNewUserServiceInterface) *UserServiceInterface {
	mock := &UserServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
