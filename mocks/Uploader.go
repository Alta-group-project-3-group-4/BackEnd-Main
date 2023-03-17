// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// Uploader is an autogenerated mock type for the Uploader type
type Uploader struct {
	mock.Mock
}

// UploadImageS3 provides a mock function with given fields: file, homestayId
func (_m *Uploader) UploadImageS3(file multipart.FileHeader, homestayId string) (string, error) {
	ret := _m.Called(file, homestayId)

	var r0 string
	if rf, ok := ret.Get(0).(func(multipart.FileHeader, string) string); ok {
		r0 = rf(file, homestayId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(multipart.FileHeader, string) error); ok {
		r1 = rf(file, homestayId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUploader interface {
	mock.TestingT
	Cleanup(func())
}

// NewUploader creates a new instance of Uploader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUploader(t mockConstructorTestingTNewUploader) *Uploader {
	mock := &Uploader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
