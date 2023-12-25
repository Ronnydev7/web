package token

import (
	"context"
	"errors"
	"fmt"
)

type PerformingLoginToken struct {
	PrivacyToken
	Email string
}

type performingLoginTokenKey struct{}

func (PerformingLoginToken) GetContextKey() interface{} {
	return performingLoginTokenKey{}
}

func NewContextWithPerformingLoginToken(parent context.Context, email string) context.Context {
	return context.WithValue(parent, performingLoginTokenKey{}, &PerformingLoginToken{Email: email})
}

func PerformingLoginTokenFromContext(ctx context.Context) *PerformingLoginToken {
	token, _ := ctx.Value(performingLoginTokenKey{}).(*PerformingLoginToken)
	return token
}

func GetPerformingLoginTokenValidatorFunc(email string) func(PrivacyToken) error {
	return func(t PrivacyToken) error {
		actualToken, ok := t.(*PerformingLoginToken)
		if !ok {
			return fmt.Errorf("invalid privacy token type %T", t)
		}
		if actualToken.Email == email {
			return nil
		}
		return errors.New("privacy token has incorrect email")
	}
}
