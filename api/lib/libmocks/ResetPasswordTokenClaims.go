// Code generated by mockery v2.36.1. DO NOT EDIT.

package libmocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// ResetPasswordTokenClaims is an autogenerated mock type for the ResetPasswordTokenClaims type
type ResetPasswordTokenClaims struct {
	mock.Mock
}

type ResetPasswordTokenClaims_Expecter struct {
	mock *mock.Mock
}

func (_m *ResetPasswordTokenClaims) EXPECT() *ResetPasswordTokenClaims_Expecter {
	return &ResetPasswordTokenClaims_Expecter{mock: &_m.Mock}
}

// GetEmailCredentialId provides a mock function with given fields:
func (_m *ResetPasswordTokenClaims) GetEmailCredentialId() (int, bool) {
	ret := _m.Called()

	var r0 int
	var r1 bool
	if rf, ok := ret.Get(0).(func() (int, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ResetPasswordTokenClaims_GetEmailCredentialId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEmailCredentialId'
type ResetPasswordTokenClaims_GetEmailCredentialId_Call struct {
	*mock.Call
}

// GetEmailCredentialId is a helper method to define mock.On call
func (_e *ResetPasswordTokenClaims_Expecter) GetEmailCredentialId() *ResetPasswordTokenClaims_GetEmailCredentialId_Call {
	return &ResetPasswordTokenClaims_GetEmailCredentialId_Call{Call: _e.mock.On("GetEmailCredentialId")}
}

func (_c *ResetPasswordTokenClaims_GetEmailCredentialId_Call) Run(run func()) *ResetPasswordTokenClaims_GetEmailCredentialId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResetPasswordTokenClaims_GetEmailCredentialId_Call) Return(_a0 int, _a1 bool) *ResetPasswordTokenClaims_GetEmailCredentialId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResetPasswordTokenClaims_GetEmailCredentialId_Call) RunAndReturn(run func() (int, bool)) *ResetPasswordTokenClaims_GetEmailCredentialId_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpiresAt provides a mock function with given fields:
func (_m *ResetPasswordTokenClaims) GetExpiresAt() (time.Time, bool) {
	ret := _m.Called()

	var r0 time.Time
	var r1 bool
	if rf, ok := ret.Get(0).(func() (time.Time, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ResetPasswordTokenClaims_GetExpiresAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpiresAt'
type ResetPasswordTokenClaims_GetExpiresAt_Call struct {
	*mock.Call
}

// GetExpiresAt is a helper method to define mock.On call
func (_e *ResetPasswordTokenClaims_Expecter) GetExpiresAt() *ResetPasswordTokenClaims_GetExpiresAt_Call {
	return &ResetPasswordTokenClaims_GetExpiresAt_Call{Call: _e.mock.On("GetExpiresAt")}
}

func (_c *ResetPasswordTokenClaims_GetExpiresAt_Call) Run(run func()) *ResetPasswordTokenClaims_GetExpiresAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResetPasswordTokenClaims_GetExpiresAt_Call) Return(_a0 time.Time, _a1 bool) *ResetPasswordTokenClaims_GetExpiresAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResetPasswordTokenClaims_GetExpiresAt_Call) RunAndReturn(run func() (time.Time, bool)) *ResetPasswordTokenClaims_GetExpiresAt_Call {
	_c.Call.Return(run)
	return _c
}

// NewResetPasswordTokenClaims creates a new instance of ResetPasswordTokenClaims. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResetPasswordTokenClaims(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResetPasswordTokenClaims {
	mock := &ResetPasswordTokenClaims{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
