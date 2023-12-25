// Code generated by mockery v2.36.1. DO NOT EDIT.

package libmocks

import (
	intl "api/intl"
	lib "api/lib"

	mock "github.com/stretchr/testify/mock"
)

// BlobStorage is an autogenerated mock type for the BlobStorage type
type BlobStorage struct {
	mock.Mock
}

type BlobStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *BlobStorage) EXPECT() *BlobStorage_Expecter {
	return &BlobStorage_Expecter{mock: &_m.Mock}
}

// DeleteExternalMedia provides a mock function with given fields: key
func (_m *BlobStorage) DeleteExternalMedia(key string) intl.IntlError {
	ret := _m.Called(key)

	var r0 intl.IntlError
	if rf, ok := ret.Get(0).(func(string) intl.IntlError); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(intl.IntlError)
		}
	}

	return r0
}

// BlobStorage_DeleteExternalMedia_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteExternalMedia'
type BlobStorage_DeleteExternalMedia_Call struct {
	*mock.Call
}

// DeleteExternalMedia is a helper method to define mock.On call
//   - key string
func (_e *BlobStorage_Expecter) DeleteExternalMedia(key interface{}) *BlobStorage_DeleteExternalMedia_Call {
	return &BlobStorage_DeleteExternalMedia_Call{Call: _e.mock.On("DeleteExternalMedia", key)}
}

func (_c *BlobStorage_DeleteExternalMedia_Call) Run(run func(key string)) *BlobStorage_DeleteExternalMedia_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BlobStorage_DeleteExternalMedia_Call) Return(_a0 intl.IntlError) *BlobStorage_DeleteExternalMedia_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobStorage_DeleteExternalMedia_Call) RunAndReturn(run func(string) intl.IntlError) *BlobStorage_DeleteExternalMedia_Call {
	_c.Call.Return(run)
	return _c
}

// GetSignedExternalMediaDownloadUrl provides a mock function with given fields: key
func (_m *BlobStorage) GetSignedExternalMediaDownloadUrl(key string) (string, intl.IntlError) {
	ret := _m.Called(key)

	var r0 string
	var r1 intl.IntlError
	if rf, ok := ret.Get(0).(func(string) (string, intl.IntlError)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) intl.IntlError); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(intl.IntlError)
		}
	}

	return r0, r1
}

// BlobStorage_GetSignedExternalMediaDownloadUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSignedExternalMediaDownloadUrl'
type BlobStorage_GetSignedExternalMediaDownloadUrl_Call struct {
	*mock.Call
}

// GetSignedExternalMediaDownloadUrl is a helper method to define mock.On call
//   - key string
func (_e *BlobStorage_Expecter) GetSignedExternalMediaDownloadUrl(key interface{}) *BlobStorage_GetSignedExternalMediaDownloadUrl_Call {
	return &BlobStorage_GetSignedExternalMediaDownloadUrl_Call{Call: _e.mock.On("GetSignedExternalMediaDownloadUrl", key)}
}

func (_c *BlobStorage_GetSignedExternalMediaDownloadUrl_Call) Run(run func(key string)) *BlobStorage_GetSignedExternalMediaDownloadUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BlobStorage_GetSignedExternalMediaDownloadUrl_Call) Return(_a0 string, _a1 intl.IntlError) *BlobStorage_GetSignedExternalMediaDownloadUrl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlobStorage_GetSignedExternalMediaDownloadUrl_Call) RunAndReturn(run func(string) (string, intl.IntlError)) *BlobStorage_GetSignedExternalMediaDownloadUrl_Call {
	_c.Call.Return(run)
	return _c
}

// GetSignedExternalMediaUploadUrl provides a mock function with given fields: spec
func (_m *BlobStorage) GetSignedExternalMediaUploadUrl(spec *lib.BlobUploadSpec) (string, intl.IntlError) {
	ret := _m.Called(spec)

	var r0 string
	var r1 intl.IntlError
	if rf, ok := ret.Get(0).(func(*lib.BlobUploadSpec) (string, intl.IntlError)); ok {
		return rf(spec)
	}
	if rf, ok := ret.Get(0).(func(*lib.BlobUploadSpec) string); ok {
		r0 = rf(spec)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*lib.BlobUploadSpec) intl.IntlError); ok {
		r1 = rf(spec)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(intl.IntlError)
		}
	}

	return r0, r1
}

// BlobStorage_GetSignedExternalMediaUploadUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSignedExternalMediaUploadUrl'
type BlobStorage_GetSignedExternalMediaUploadUrl_Call struct {
	*mock.Call
}

// GetSignedExternalMediaUploadUrl is a helper method to define mock.On call
//   - spec *lib.BlobUploadSpec
func (_e *BlobStorage_Expecter) GetSignedExternalMediaUploadUrl(spec interface{}) *BlobStorage_GetSignedExternalMediaUploadUrl_Call {
	return &BlobStorage_GetSignedExternalMediaUploadUrl_Call{Call: _e.mock.On("GetSignedExternalMediaUploadUrl", spec)}
}

func (_c *BlobStorage_GetSignedExternalMediaUploadUrl_Call) Run(run func(spec *lib.BlobUploadSpec)) *BlobStorage_GetSignedExternalMediaUploadUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*lib.BlobUploadSpec))
	})
	return _c
}

func (_c *BlobStorage_GetSignedExternalMediaUploadUrl_Call) Return(_a0 string, _a1 intl.IntlError) *BlobStorage_GetSignedExternalMediaUploadUrl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlobStorage_GetSignedExternalMediaUploadUrl_Call) RunAndReturn(run func(*lib.BlobUploadSpec) (string, intl.IntlError)) *BlobStorage_GetSignedExternalMediaUploadUrl_Call {
	_c.Call.Return(run)
	return _c
}

// NewBlobStorage creates a new instance of BlobStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobStorage {
	mock := &BlobStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
