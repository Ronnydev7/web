package token

import (
	"context"
)

type EmailSignupToken struct {
	BaseSignupToken
	Email string
}

func (token EmailSignupToken) GetEmail() string {
	return token.Email
}

func EmailSignupTokenFromContext(ctx context.Context) *EmailSignupToken {
	token, _ := ctx.Value(signupTokenKey{}).(*EmailSignupToken)
	return token
}
