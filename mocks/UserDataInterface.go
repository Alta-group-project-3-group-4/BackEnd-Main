// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	users "airbnb/feature/users"

	mock "github.com/stretchr/testify/mock"
)

// UserDataInterface is an autogenerated mock type for the UserDataInterface type
type UserDataInterface struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: id
func (_m *UserDataInterface) Destroy(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: userEntity, id
func (_m *UserDataInterface) Edit(userEntity users.UserEntity, id uint) (uint, error) {
	ret := _m.Called(userEntity, id)

	var r0 uint
	if rf, ok := ret.Get(0).(func(users.UserEntity, uint) uint); ok {
		r0 = rf(userEntity, id)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserEntity, uint) error); ok {
		r1 = rf(userEntity, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAll provides a mock function with given fields:
func (_m *UserDataInterface) SelectAll() ([]users.UserEntity, error) {
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

// SelectById provides a mock function with given fields: id
func (_m *UserDataInterface) SelectById(id uint) (users.UserEntity, error) {
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

// Store provides a mock function with given fields: userEntity
func (_m *UserDataInterface) Store(userEntity users.UserEntity) (uint, error) {
	ret := _m.Called(userEntity)

	var r0 uint
	if rf, ok := ret.Get(0).(func(users.UserEntity) uint); ok {
		r0 = rf(userEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserEntity) error); ok {
		r1 = rf(userEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserDataInterface creates a new instance of UserDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserDataInterface(t mockConstructorTestingTNewUserDataInterface) *UserDataInterface {
	mock := &UserDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
