package lib

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/loginsession"
	"api/ent/privacy/token"
	"api/intl"
	"context"
)

type (
	LoginSessionFactory interface {
		FromRefreshTokenCookie(
			cookieReader CookieReader,
			userTokenFactory UserTokenFactory,
		) (*entgenerated.LoginSession, error)

		FromRefreshTokenCookieBypassPrivacy(
			cookieReader CookieReader,
			userTokenFactory UserTokenFactory,
		) (*entgenerated.LoginSession, error)
	}

	defaultLoginSessionFactory struct {
		LoginSessionFactory
		ctx    context.Context
		client *entgenerated.Client
	}

	NewLoginSessionFactoryFunc func(ctx context.Context, client *entgenerated.Client) LoginSessionFactory
)

var NewLoginSessionFactory NewLoginSessionFactoryFunc = func(
	ctx context.Context,
	client *entgenerated.Client,
) LoginSessionFactory {
	return &defaultLoginSessionFactory{
		ctx:    ctx,
		client: client,
	}
}

func (factory defaultLoginSessionFactory) FromRefreshTokenCookie(
	cookieReader CookieReader,
	userTokenFactory UserTokenFactory,
) (*entgenerated.LoginSession, error) {
	loginSessionId, err := getLoginSessionId(cookieReader, userTokenFactory)
	if err != nil {
		return nil, err
	}

	loginSession, err := factory.client.LoginSession.Get(factory.ctx, loginSessionId)
	if err != nil {
		return nil, err
	}
	return loginSession, nil
}

// Use this method only when the you have to get the login session without providing the
// legitimate owner viewer in the context, e.g when refreshing the auth token (this is
// probably the only use case). Note that this will eager load the owner edge.
func (factory defaultLoginSessionFactory) FromRefreshTokenCookieBypassPrivacy(
	cookieReader CookieReader,
	userTokenFactory UserTokenFactory,
) (*entgenerated.LoginSession, error) {
	loginSessionId, err := getLoginSessionId(cookieReader, userTokenFactory)
	if err != nil {
		return nil, err
	}
	liftedCtx := token.NewContextWithPerformingAuthRefreshToken(factory.ctx, loginSessionId)
	loginSession, err := factory.client.LoginSession.Query().
		Where(loginsession.ID(loginSessionId)).
		WithOwner().
		Only(liftedCtx)
	if err != nil {
		return nil, err
	}
	return loginSession, nil
}

func getLoginSessionId(cookieReader CookieReader, userTokenFactory UserTokenFactory) (int, error) {
	refreshTokenCookie, err := cookieReader.ReadRefreshToken()
	if err != nil {
		return 0, err
	}

	refreshToken, err := userTokenFactory.ParseRefreshToken(refreshTokenCookie.Value)
	if err != nil {
		return 0, err
	}

	loginSessionId, exists := refreshToken.GetLoginSessionId()
	if !exists {
		return 0, &intl.InvalidJwtTokenError{}
	}
	return loginSessionId, nil
}
