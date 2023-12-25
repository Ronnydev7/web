package lib

import (
	"api/ent/entgenerated"
	"context"
)

type (
	UserFactory interface {
		FromRefreshTokenCookieBypassPrivacy(context.Context, CookieReader, UserTokenFactory) (*entgenerated.User, error)
	}

	defaultUserFactory struct {
		UserFactory
		client *entgenerated.Client
	}

	NewUserFactoryFunc func(*entgenerated.Client) UserFactory
)

var NewUserFactory NewUserFactoryFunc = func(client *entgenerated.Client) UserFactory {
	return &defaultUserFactory{
		client: client,
	}
}

func (factory defaultUserFactory) FromRefreshTokenCookieBypassPrivacy(
	ctx context.Context,
	cookieReader CookieReader,
	userTokenFactory UserTokenFactory,
) (*entgenerated.User, error) {
	loginSessionFactory := NewLoginSessionFactory(ctx, factory.client)
	loginSession, err := loginSessionFactory.FromRefreshTokenCookieBypassPrivacy(cookieReader, userTokenFactory)
	if err != nil {
		return nil, err
	}
	return loginSession.Owner(ctx)
}
