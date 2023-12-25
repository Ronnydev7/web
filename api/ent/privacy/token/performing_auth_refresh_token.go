package token

import "context"

type (
	PerformingAuthRefreshToken struct {
		PrivacyToken
		LoginSessionId int
	}

	performingAuthRefreshTokenKey struct{}
)

func (PerformingAuthRefreshToken) GetContextKey() interface{} {
	return performingAuthRefreshTokenKey{}
}

func NewContextWithPerformingAuthRefreshToken(parent context.Context, loginSessionId int) context.Context {
	return context.WithValue(parent, performingAuthRefreshTokenKey{}, &PerformingAuthRefreshToken{
		LoginSessionId: loginSessionId,
	})
}

func PerformingAuthRefreshTokenFromContext(ctx context.Context) *PerformingAuthRefreshToken {
	token, _ := ctx.Value(performingAuthRefreshTokenKey{}).(*PerformingAuthRefreshToken)
	return token
}
