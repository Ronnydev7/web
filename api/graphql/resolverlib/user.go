package resolverlib

import (
	"api/ent/entgenerated"
	"api/graphql/gqlgenerated"
	"api/intl"
	"api/intl/intlgenerated"
	"api/lib"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

const LOGGER_NAME = "user_resolver"

func NewUserResolverLogger() lib.Logger {
	return lib.NewLogger(LOGGER_NAME)
}

func UnknownRefreshUserTokenDeath(ctx context.Context, err error) (*gqlgenerated.UserToken, error) {
	logger := NewUserResolverLogger()
	logger.LogError(err)
	return nil, graphql.ErrorOnPath(ctx, intl.HandleIntlError(err, intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR))
}

func UnauthorizedRefreshUserTokenCall(
	ctx context.Context,
	httpErrorWriter lib.HttpErrorWriter,
	unexpectedError error,
) (*gqlgenerated.UserToken, error) {
	logger := NewUserResolverLogger()
	if unexpectedError != nil {
		logger.LogError(unexpectedError)
	}
	httpErrorWriter.WriteUnauthorized()
	return nil, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.COMMON_STRINGS__UNAUTHORIZED))
}

func UnknownLogoutDeath(ctx context.Context, err error) (bool, error) {
	NewUserResolverLogger().LogError(err)
	return false, graphql.ErrorOnPath(ctx, intl.HandleIntlError(err, intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR))
}

func UnauthorizedLogoutCall(ctx context.Context, httpErrorWriter lib.HttpErrorWriter, unexpectedError error) (bool, error) {
	logger := NewUserResolverLogger()
	if unexpectedError != nil {
		logger.LogError(unexpectedError)
	}
	httpErrorWriter.WriteUnauthorized()
	return false, graphql.ErrorOnPath(ctx, errors.New(intlgenerated.COMMON_STRINGS__UNAUTHORIZED))
}

func GetLoginSession(
	ctx context.Context,
	entClient *entgenerated.Client,
	cookieReader lib.CookieReader,
	userTokenFactory lib.UserTokenFactory,
) (*entgenerated.LoginSession, error) {
	factory := lib.NewLoginSessionFactory(ctx, entClient)
	return factory.FromRefreshTokenCookieBypassPrivacy(cookieReader, userTokenFactory)
}
