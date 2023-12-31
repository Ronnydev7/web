// Code generated by mockery v2.36.1. DO NOT EDIT.

package configmocks

import mock "github.com/stretchr/testify/mock"

// AwsConfig is an autogenerated mock type for the AwsConfig type
type AwsConfig struct {
	mock.Mock
}

type AwsConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *AwsConfig) EXPECT() *AwsConfig_Expecter {
	return &AwsConfig_Expecter{mock: &_m.Mock}
}

// GetAccessKeyId provides a mock function with given fields:
func (_m *AwsConfig) GetAccessKeyId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AwsConfig_GetAccessKeyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessKeyId'
type AwsConfig_GetAccessKeyId_Call struct {
	*mock.Call
}

// GetAccessKeyId is a helper method to define mock.On call
func (_e *AwsConfig_Expecter) GetAccessKeyId() *AwsConfig_GetAccessKeyId_Call {
	return &AwsConfig_GetAccessKeyId_Call{Call: _e.mock.On("GetAccessKeyId")}
}

func (_c *AwsConfig_GetAccessKeyId_Call) Run(run func()) *AwsConfig_GetAccessKeyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AwsConfig_GetAccessKeyId_Call) Return(_a0 string) *AwsConfig_GetAccessKeyId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AwsConfig_GetAccessKeyId_Call) RunAndReturn(run func() string) *AwsConfig_GetAccessKeyId_Call {
	_c.Call.Return(run)
	return _c
}

// GetRegion provides a mock function with given fields:
func (_m *AwsConfig) GetRegion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AwsConfig_GetRegion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRegion'
type AwsConfig_GetRegion_Call struct {
	*mock.Call
}

// GetRegion is a helper method to define mock.On call
func (_e *AwsConfig_Expecter) GetRegion() *AwsConfig_GetRegion_Call {
	return &AwsConfig_GetRegion_Call{Call: _e.mock.On("GetRegion")}
}

func (_c *AwsConfig_GetRegion_Call) Run(run func()) *AwsConfig_GetRegion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AwsConfig_GetRegion_Call) Return(_a0 string) *AwsConfig_GetRegion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AwsConfig_GetRegion_Call) RunAndReturn(run func() string) *AwsConfig_GetRegion_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecretAccessKey provides a mock function with given fields:
func (_m *AwsConfig) GetSecretAccessKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AwsConfig_GetSecretAccessKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecretAccessKey'
type AwsConfig_GetSecretAccessKey_Call struct {
	*mock.Call
}

// GetSecretAccessKey is a helper method to define mock.On call
func (_e *AwsConfig_Expecter) GetSecretAccessKey() *AwsConfig_GetSecretAccessKey_Call {
	return &AwsConfig_GetSecretAccessKey_Call{Call: _e.mock.On("GetSecretAccessKey")}
}

func (_c *AwsConfig_GetSecretAccessKey_Call) Run(run func()) *AwsConfig_GetSecretAccessKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AwsConfig_GetSecretAccessKey_Call) Return(_a0 string) *AwsConfig_GetSecretAccessKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AwsConfig_GetSecretAccessKey_Call) RunAndReturn(run func() string) *AwsConfig_GetSecretAccessKey_Call {
	_c.Call.Return(run)
	return _c
}

// NewAwsConfig creates a new instance of AwsConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAwsConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *AwsConfig {
	mock := &AwsConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
