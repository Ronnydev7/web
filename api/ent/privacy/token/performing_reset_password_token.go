package token

import "context"

type (
	PerformingResetPasswordToken struct {
		PrivacyToken
		Email string
	}

	performingResetPasswordTokenKey struct{}
)

func (PerformingResetPasswordToken) GetContextKey() interface{} {
	return &performingResetPasswordTokenKey{}
}

func NewContextWithPerformingResetPasswordToken(parent context.Context, email string) context.Context {
	return context.WithValue(parent, &performingResetPasswordTokenKey{}, &PerformingResetPasswordToken{Email: email})
}

func PerformingResetPasswordTokenFromContext(ctx context.Context) *PerformingResetPasswordToken {
	token, _ := ctx.Value(&performingResetPasswordTokenKey{}).(*PerformingResetPasswordToken)
	return token
}
