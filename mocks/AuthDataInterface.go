// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	users "airbnb/feature/users"

	mock "github.com/stretchr/testify/mock"
)

// AuthDataInterface is an autogenerated mock type for the AuthDataInterface type
type AuthDataInterface struct {
	mock.Mock
}

// EditPassword provides a mock function with given fields: id, newPassword
func (_m *AuthDataInterface) EditPassword(id uint, newPassword string) error {
	ret := _m.Called(id, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(id, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmailOrId provides a mock function with given fields: email, id
func (_m *AuthDataInterface) GetUserByEmailOrId(email string, id uint) (users.UserEntity, error) {
	ret := _m.Called(email, id)

	var r0 users.UserEntity
	if rf, ok := ret.Get(0).(func(string, uint) users.UserEntity); ok {
		r0 = rf(email, id)
	} else {
		r0 = ret.Get(0).(users.UserEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(email, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: request
func (_m *AuthDataInterface) Register(request users.UserEntity) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.UserEntity) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuthDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthDataInterface creates a new instance of AuthDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthDataInterface(t mockConstructorTestingTNewAuthDataInterface) *AuthDataInterface {
	mock := &AuthDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
