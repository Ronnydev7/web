package token

import "context"

type (
	SignupToken interface {
		PrivacyToken
	}

	BaseSignupToken struct {
		SignupToken
	}
)

type signupTokenKey struct{}

func (BaseSignupToken) GetContextKey() interface{} {
	return signupTokenKey{}
}

func NewContextWithSignupToken(parent context.Context, token SignupToken) context.Context {
	return context.WithValue(parent, signupTokenKey{}, token)
}
