// Code generated by mockery v2.36.1. DO NOT EDIT.

package libmocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// RefreshTokenClaims is an autogenerated mock type for the RefreshTokenClaims type
type RefreshTokenClaims struct {
	mock.Mock
}

type RefreshTokenClaims_Expecter struct {
	mock *mock.Mock
}

func (_m *RefreshTokenClaims) EXPECT() *RefreshTokenClaims_Expecter {
	return &RefreshTokenClaims_Expecter{mock: &_m.Mock}
}

// GetExpiresAt provides a mock function with given fields:
func (_m *RefreshTokenClaims) GetExpiresAt() (time.Time, bool) {
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

// RefreshTokenClaims_GetExpiresAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpiresAt'
type RefreshTokenClaims_GetExpiresAt_Call struct {
	*mock.Call
}

// GetExpiresAt is a helper method to define mock.On call
func (_e *RefreshTokenClaims_Expecter) GetExpiresAt() *RefreshTokenClaims_GetExpiresAt_Call {
	return &RefreshTokenClaims_GetExpiresAt_Call{Call: _e.mock.On("GetExpiresAt")}
}

func (_c *RefreshTokenClaims_GetExpiresAt_Call) Run(run func()) *RefreshTokenClaims_GetExpiresAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RefreshTokenClaims_GetExpiresAt_Call) Return(_a0 time.Time, _a1 bool) *RefreshTokenClaims_GetExpiresAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RefreshTokenClaims_GetExpiresAt_Call) RunAndReturn(run func() (time.Time, bool)) *RefreshTokenClaims_GetExpiresAt_Call {
	_c.Call.Return(run)
	return _c
}

// GetIssuedAt provides a mock function with given fields:
func (_m *RefreshTokenClaims) GetIssuedAt() (time.Time, bool) {
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

// RefreshTokenClaims_GetIssuedAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIssuedAt'
type RefreshTokenClaims_GetIssuedAt_Call struct {
	*mock.Call
}

// GetIssuedAt is a helper method to define mock.On call
func (_e *RefreshTokenClaims_Expecter) GetIssuedAt() *RefreshTokenClaims_GetIssuedAt_Call {
	return &RefreshTokenClaims_GetIssuedAt_Call{Call: _e.mock.On("GetIssuedAt")}
}

func (_c *RefreshTokenClaims_GetIssuedAt_Call) Run(run func()) *RefreshTokenClaims_GetIssuedAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RefreshTokenClaims_GetIssuedAt_Call) Return(_a0 time.Time, _a1 bool) *RefreshTokenClaims_GetIssuedAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RefreshTokenClaims_GetIssuedAt_Call) RunAndReturn(run func() (time.Time, bool)) *RefreshTokenClaims_GetIssuedAt_Call {
	_c.Call.Return(run)
	return _c
}

// GetLoginSessionId provides a mock function with given fields:
func (_m *RefreshTokenClaims) GetLoginSessionId() (int, bool) {
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

// RefreshTokenClaims_GetLoginSessionId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLoginSessionId'
type RefreshTokenClaims_GetLoginSessionId_Call struct {
	*mock.Call
}

// GetLoginSessionId is a helper method to define mock.On call
func (_e *RefreshTokenClaims_Expecter) GetLoginSessionId() *RefreshTokenClaims_GetLoginSessionId_Call {
	return &RefreshTokenClaims_GetLoginSessionId_Call{Call: _e.mock.On("GetLoginSessionId")}
}

func (_c *RefreshTokenClaims_GetLoginSessionId_Call) Run(run func()) *RefreshTokenClaims_GetLoginSessionId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RefreshTokenClaims_GetLoginSessionId_Call) Return(_a0 int, _a1 bool) *RefreshTokenClaims_GetLoginSessionId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RefreshTokenClaims_GetLoginSessionId_Call) RunAndReturn(run func() (int, bool)) *RefreshTokenClaims_GetLoginSessionId_Call {
	_c.Call.Return(run)
	return _c
}

// NewRefreshTokenClaims creates a new instance of RefreshTokenClaims. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRefreshTokenClaims(t interface {
	mock.TestingT
	Cleanup(func())
}) *RefreshTokenClaims {
	mock := &RefreshTokenClaims{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
