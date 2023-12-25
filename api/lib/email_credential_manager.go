package lib

import (
	"api/config"
	"api/ent/entgenerated"
	"api/ent/privacy/entviewer"
	"api/ent/privacy/token"
	"api/intl"
	"api/intl/intlgenerated"
	"api/lib/cookies"
	"api/privacy/viewer"
	"context"
	"errors"

	"entgo.io/ent/entql"
	"golang.org/x/crypto/bcrypt"
)

type (
	EmailCredentialManagerFactory = func(config.AuthConfig) EmailCredentialManager

	EmailCredentialLoginResult struct {
		AuthToken    string
		RefreshToken string
	}

	EmailCredentialManager interface {
		Login(
			ctx context.Context,
			client *entgenerated.Client,
			email string,
			rawPassword string,
		) (*EmailCredentialLoginResult, error)
	}

	EmailCredentialManagerWithConfig struct {
		EmailCredentialManager
		authC config.AuthConfig
	}

	EmailCredentialNotFoundError struct {
		intl.IntlError
	}
)

const EMAIL_CREDENTIAL_MANAGER_LOGGER_NAME = "lib.email_credential_manager"

var NewEmailCredentialManager EmailCredentialManagerFactory = func(authConfig config.AuthConfig) EmailCredentialManager {
	return &EmailCredentialManagerWithConfig{
		authC: authConfig,
	}
}

func (manager EmailCredentialManagerWithConfig) Login(
	ctx context.Context,
	client *entgenerated.Client,
	email string,
	rawPassword string,
) (*EmailCredentialLoginResult, error) {
	ctxWithToken := token.NewContextWithPerformingLoginToken(ctx, email)
	query := client.EmailCredential.Query()
	query.Filter().WhereEmail(entql.StringEQ(email))
	userCredential, err := query.WithOwner().Only(ctxWithToken)
	if err != nil {
		if entgenerated.IsNotFound(err) {
			return nil, &EmailCredentialNotFoundError{}
		}
		NewLogger(EMAIL_CREDENTIAL_MANAGER_LOGGER_NAME).LogError(err)
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword(userCredential.PasswordHash, []byte(rawPassword))
	if err != nil {
		return nil, &EmailCredentialNotFoundError{}
	}
	user := userCredential.Edges.Owner

	ctxWithToken = viewer.NewContext(ctx, entviewer.NewUserViewerFromUser(user))
	// Create and return user token
	loginSession, err := client.LoginSession.
		Create().
		SetOwner(user).
		SetLastLoginTime(TimeNow()).
		Save(ctxWithToken)
	if err != nil {
		entUtils := NewEntUtils()
		return nil, entUtils.HandlePermissionError(err)
	}

	userTokenFactory := NewUserTokenFactory()

	refreshToken, err := userTokenFactory.CreateRefreshToken(loginSession)
	if err != nil {
		return doEmailLoginMysteriousDeath(err)
	}
	refreshTokenClaims, err := userTokenFactory.ParseRefreshToken(refreshToken)
	if err != nil {
		return doEmailLoginMysteriousDeath(err)
	}
	expiredAt, exists := refreshTokenClaims.GetExpiresAt()
	if !exists {
		return doEmailLoginMysteriousDeath(errors.New("exp claim should exist but it's not found"))
	}

	cookieWriter, err := RequireCookieWriterFromContext(ctx)
	if err != nil {
		return doEmailLoginMysteriousDeath(err)
	}
	cookieWriter.Write(cookies.NewRefreshTokenCookie(
		refreshToken,
		expiredAt,
		manager.authC.GetRefreshTokenCookieSameSiteMode()),
	)

	authToken, err := userTokenFactory.CreateAuthToken(user)
	if err != nil {
		return doEmailLoginMysteriousDeath(err)
	}

	return &EmailCredentialLoginResult{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}

func doEmailLoginMysteriousDeath(err error) (*EmailCredentialLoginResult, error) {
	NewLogger(EMAIL_CREDENTIAL_MANAGER_LOGGER_NAME).LogError(err)
	return nil, errors.New(intlgenerated.COMMON_STRINGS__UNKNOWN_SERVER_ERROR)
}

func (EmailCredentialNotFoundError) Error() string {
	return "User Error: incorrect credential (either email is not found or password is incorrect"
}

func (EmailCredentialNotFoundError) GetIntlKey() string {
	return intlgenerated.EMAIL_CREDENTIAL__INVALID_EMAIL_CREDENTIAL
}
