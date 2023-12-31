// Code generated by mockery v2.36.1. DO NOT EDIT.

package libmocks

import (
	intl "api/intl"
	lib "api/lib"

	mock "github.com/stretchr/testify/mock"
)

// Mailer is an autogenerated mock type for the Mailer type
type Mailer struct {
	mock.Mock
}

type Mailer_Expecter struct {
	mock *mock.Mock
}

func (_m *Mailer) EXPECT() *Mailer_Expecter {
	return &Mailer_Expecter{mock: &_m.Mock}
}

// SendConfirmSignupEmailEmail provides a mock function with given fields: receiver, emailSignupUrl
func (_m *Mailer) SendConfirmSignupEmailEmail(receiver string, emailSignupUrl string) intl.IntlError {
	ret := _m.Called(receiver, emailSignupUrl)

	var r0 intl.IntlError
	if rf, ok := ret.Get(0).(func(string, string) intl.IntlError); ok {
		r0 = rf(receiver, emailSignupUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(intl.IntlError)
		}
	}

	return r0
}

// Mailer_SendConfirmSignupEmailEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendConfirmSignupEmailEmail'
type Mailer_SendConfirmSignupEmailEmail_Call struct {
	*mock.Call
}

// SendConfirmSignupEmailEmail is a helper method to define mock.On call
//   - receiver string
//   - emailSignupUrl string
func (_e *Mailer_Expecter) SendConfirmSignupEmailEmail(receiver interface{}, emailSignupUrl interface{}) *Mailer_SendConfirmSignupEmailEmail_Call {
	return &Mailer_SendConfirmSignupEmailEmail_Call{Call: _e.mock.On("SendConfirmSignupEmailEmail", receiver, emailSignupUrl)}
}

func (_c *Mailer_SendConfirmSignupEmailEmail_Call) Run(run func(receiver string, emailSignupUrl string)) *Mailer_SendConfirmSignupEmailEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Mailer_SendConfirmSignupEmailEmail_Call) Return(_a0 intl.IntlError) *Mailer_SendConfirmSignupEmailEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mailer_SendConfirmSignupEmailEmail_Call) RunAndReturn(run func(string, string) intl.IntlError) *Mailer_SendConfirmSignupEmailEmail_Call {
	_c.Call.Return(run)
	return _c
}

// SendResetPasswordEmail provides a mock function with given fields: receiver, resetEmailUrl
func (_m *Mailer) SendResetPasswordEmail(receiver string, resetEmailUrl string) intl.IntlError {
	ret := _m.Called(receiver, resetEmailUrl)

	var r0 intl.IntlError
	if rf, ok := ret.Get(0).(func(string, string) intl.IntlError); ok {
		r0 = rf(receiver, resetEmailUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(intl.IntlError)
		}
	}

	return r0
}

// Mailer_SendResetPasswordEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendResetPasswordEmail'
type Mailer_SendResetPasswordEmail_Call struct {
	*mock.Call
}

// SendResetPasswordEmail is a helper method to define mock.On call
//   - receiver string
//   - resetEmailUrl string
func (_e *Mailer_Expecter) SendResetPasswordEmail(receiver interface{}, resetEmailUrl interface{}) *Mailer_SendResetPasswordEmail_Call {
	return &Mailer_SendResetPasswordEmail_Call{Call: _e.mock.On("SendResetPasswordEmail", receiver, resetEmailUrl)}
}

func (_c *Mailer_SendResetPasswordEmail_Call) Run(run func(receiver string, resetEmailUrl string)) *Mailer_SendResetPasswordEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Mailer_SendResetPasswordEmail_Call) Return(_a0 intl.IntlError) *Mailer_SendResetPasswordEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mailer_SendResetPasswordEmail_Call) RunAndReturn(run func(string, string) intl.IntlError) *Mailer_SendResetPasswordEmail_Call {
	_c.Call.Return(run)
	return _c
}

// SendTemplateMail provides a mock function with given fields: _a0
func (_m *Mailer) SendTemplateMail(_a0 lib.TemplateMail) intl.IntlError {
	ret := _m.Called(_a0)

	var r0 intl.IntlError
	if rf, ok := ret.Get(0).(func(lib.TemplateMail) intl.IntlError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(intl.IntlError)
		}
	}

	return r0
}

// Mailer_SendTemplateMail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTemplateMail'
type Mailer_SendTemplateMail_Call struct {
	*mock.Call
}

// SendTemplateMail is a helper method to define mock.On call
//   - _a0 lib.TemplateMail
func (_e *Mailer_Expecter) SendTemplateMail(_a0 interface{}) *Mailer_SendTemplateMail_Call {
	return &Mailer_SendTemplateMail_Call{Call: _e.mock.On("SendTemplateMail", _a0)}
}

func (_c *Mailer_SendTemplateMail_Call) Run(run func(_a0 lib.TemplateMail)) *Mailer_SendTemplateMail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(lib.TemplateMail))
	})
	return _c
}

func (_c *Mailer_SendTemplateMail_Call) Return(_a0 intl.IntlError) *Mailer_SendTemplateMail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mailer_SendTemplateMail_Call) RunAndReturn(run func(lib.TemplateMail) intl.IntlError) *Mailer_SendTemplateMail_Call {
	_c.Call.Return(run)
	return _c
}

// NewMailer creates a new instance of Mailer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMailer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mailer {
	mock := &Mailer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
