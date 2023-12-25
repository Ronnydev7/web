package lib

import (
	"api/config"
	"api/ent/entgenerated"
	"api/intl"
	"context"
	"fmt"
	"strconv"
	"time"

	_ "github.com/99designs/gqlgen/codegen/config"
	"github.com/golang-jwt/jwt/v4"
)

type (
	AuthTokenClaims interface {
		GetUserId() (int, bool)
		GetExpiresAt() (time.Time, bool)
	}

	authTokenClaims struct {
		AuthTokenClaims `json:"-"`
		jwt.RegisteredClaims
	}

	RefreshTokenClaims interface {
		GetLoginSessionId() (int, bool)
		GetIssuedAt() (time.Time, bool)
		GetExpiresAt() (time.Time, bool)
	}

	refreshTokenClaims struct {
		RefreshTokenClaims `json:"-"`
		jwt.RegisteredClaims
	}

	ResetPasswordTokenClaims interface {
		GetEmailCredentialId() (int, bool)
		GetExpiresAt() (time.Time, bool)
	}

	resetPasswordTokenClaims struct {
		ResetPasswordTokenClaims
		jwt.RegisteredClaims
	}

	UserTokenFactory interface {
		CreateAuthToken(user *entgenerated.User) (string, intl.IntlError)
		CreateRefreshToken(loginSession *entgenerated.LoginSession) (string, intl.IntlError)
		CreateResetPasswordToken(credential *entgenerated.EmailCredential) (string, intl.IntlError)
		GetAuthTokenSignature(AuthTokenClaims) ([]byte, intl.IntlError)
		GetRefreshTokenSignature(RefreshTokenClaims) ([]byte, intl.IntlError)
		GetResetPasswordTokenSignature(ResetPasswordTokenClaims, *entgenerated.EmailCredential) ([]byte, intl.IntlError)
		ParseAuthToken(authTokenString string) (AuthTokenClaims, intl.IntlError)
		ParseRefreshToken(refreshTokenString string) (RefreshTokenClaims, intl.IntlError)
		ParseResetPasswordToken(
			ctx context.Context,
			tokenString string,
			entClient *entgenerated.Client,
		) (ResetPasswordTokenClaims, intl.IntlError)
	}

	JwtUserTokenFactory struct {
		UserTokenFactory
	}

	NewUserTokenFactoryFunc = func() UserTokenFactory
)

const (
	AUTH_TOKEN_DURATION_MIN           = 5
	REFRESH_TOKEN_DURATION_MONTH      = 6
	RESET_PASSWORD_TOKEN_DURATION_MIN = 30
)

var NewUserTokenFactory NewUserTokenFactoryFunc = func() UserTokenFactory {
	return &JwtUserTokenFactory{}
}

func (factory JwtUserTokenFactory) CreateAuthToken(
	user *entgenerated.User,
) (string, intl.IntlError) {
	exp := TimeNow().Local().Add(time.Minute * AUTH_TOKEN_DURATION_MIN)
	authClaims := authTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(user.ID),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	authToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		authClaims,
	)
	signatureSecret, err := factory.GetAuthTokenSignature(authClaims)
	if err != nil {
		return "", err
	}

	authTokenString, jwtErr := authToken.SignedString(signatureSecret)
	if jwtErr != nil {
		return "", intl.HandleErrorFromJwt(jwtErr)
	}
	return authTokenString, nil
}

func (factory JwtUserTokenFactory) CreateRefreshToken(loginSession *entgenerated.LoginSession) (string, intl.IntlError) {
	jti := strconv.Itoa(loginSession.ID)
	iat := loginSession.LastLoginTime
	exp := loginSession.LastLoginTime.Add(time.Hour * 24 * 30 * REFRESH_TOKEN_DURATION_MONTH)
	claims := refreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signatureSecret, err := factory.GetRefreshTokenSignature(claims)
	if err != nil {
		return "", err
	}

	refreshTokenString, jwtErr := refreshToken.SignedString(signatureSecret)
	if jwtErr != nil {
		return "", intl.HandleErrorFromJwt(jwtErr)
	}
	return refreshTokenString, nil
}

func (factory JwtUserTokenFactory) CreateResetPasswordToken(emailCredential *entgenerated.EmailCredential) (string, intl.IntlError) {
	jti := strconv.Itoa(emailCredential.ID)
	exp := TimeNow().Add(time.Minute * RESET_PASSWORD_TOKEN_DURATION_MIN)
	claims := resetPasswordTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	signatureSecret, err := factory.GetResetPasswordTokenSignature(claims, emailCredential)
	if err != nil {
		return "", err
	}

	resetPasswordToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resetPasswordTokenString, jwtErr := resetPasswordToken.SignedString(signatureSecret)
	if jwtErr != nil {
		return "", intl.HandleErrorFromJwt(jwtErr)
	}
	return resetPasswordTokenString, nil
}

func (factory JwtUserTokenFactory) GetAuthTokenSignature(claims AuthTokenClaims) ([]byte, intl.IntlError) {
	jti, jtiExists := claims.GetUserId()
	exp, expExists := claims.GetExpiresAt()
	if jtiExists && expExists {
		authHmacData := fmt.Sprintf("%d:%d", jti, exp.Unix())
		authHmacSecret := config.GetHmacConfig().GetAuthTokenSecret()
		return GetHmac([]byte(authHmacSecret), []byte(authHmacData)), nil
	}
	return nil, &intl.InvalidJwtTokenError{}
}

func (factory JwtUserTokenFactory) GetRefreshTokenSignature(claims RefreshTokenClaims) ([]byte, intl.IntlError) {
	jti, jtiExists := claims.GetLoginSessionId()
	iat, iatExists := claims.GetIssuedAt()
	exp, expExists := claims.GetExpiresAt()

	if jtiExists && iatExists && expExists {
		refreshHmacSecret := config.GetHmacConfig().GetRefreshTokenSecret()
		refreshHmacData := fmt.Sprintf(
			"%d:%d:%d",
			jti,
			iat.Unix(),
			exp.Unix(),
		)
		return GetHmac([]byte(refreshHmacSecret), []byte(refreshHmacData)), nil
	}
	return nil, &intl.InvalidJwtTokenError{}
}

func (factory JwtUserTokenFactory) GetResetPasswordTokenSignature(
	claims ResetPasswordTokenClaims,
	credential *entgenerated.EmailCredential,
) ([]byte, intl.IntlError) {
	jti, jtiExists := claims.GetEmailCredentialId()
	exp, expExists := claims.GetExpiresAt()

	if jtiExists && expExists {
		// Use the password has as the secret. Once the password is reset, the token gets automatically
		// invalidated
		hmacSecret := credential.PasswordHash
		hmacData := fmt.Sprintf(
			"%d:%d",
			jti,
			exp.Unix(),
		)
		return GetHmac([]byte(hmacSecret), []byte(hmacData)), nil
	}
	return nil, &intl.InvalidJwtTokenError{}
}

func (factory JwtUserTokenFactory) ParseAuthToken(jwtTokenString string) (AuthTokenClaims, intl.IntlError) {
	claims := authTokenClaims{}
	_, err := jwt.ParseWithClaims(jwtTokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &intl.InvalidJwtTokenError{}
		}

		return factory.GetAuthTokenSignature(claims)
	})
	if err != nil {
		return nil, intl.HandleErrorFromJwt(err)
	}

	return &claims, nil
}

func (factory JwtUserTokenFactory) ParseRefreshToken(jwtTokenString string) (RefreshTokenClaims, intl.IntlError) {
	claims := refreshTokenClaims{}
	_, err := jwt.ParseWithClaims(jwtTokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &intl.InvalidJwtTokenError{}
		}

		return factory.GetRefreshTokenSignature(claims)
	})
	if err != nil {
		return nil, intl.HandleErrorFromJwt(err)
	}

	return &claims, nil
}

func (factory JwtUserTokenFactory) ParseResetPasswordToken(
	ctx context.Context,
	jwtTokenString string,
	entClient *entgenerated.Client,
) (ResetPasswordTokenClaims, intl.IntlError) {
	claims := resetPasswordTokenClaims{}
	_, err := jwt.ParseWithClaims(jwtTokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &intl.InvalidJwtTokenError{}
		}

		id, exists := claims.GetEmailCredentialId()
		if !exists {
			return nil, &intl.InvalidJwtTokenError{}
		}

		emailCredential, err := entClient.EmailCredential.Get(ctx, id)
		if err != nil {
			if entgenerated.IsNotFound(err) {
				return nil, &intl.InvalidJwtTokenError{}
			}
			return nil, err
		}

		return factory.GetResetPasswordTokenSignature(claims, emailCredential)
	})
	if err != nil {
		return nil, intl.HandleErrorFromJwt(err)
	}

	return &claims, nil
}

func (token authTokenClaims) GetUserId() (int, bool) {
	jti := token.ID
	if jti == "" {
		return 0, false
	}
	value, err := strconv.Atoi(jti)
	if err != nil {
		return 0, false
	}
	return value, true
}

func (token authTokenClaims) GetExpiresAt() (time.Time, bool) {
	exp := token.ExpiresAt
	if exp == nil {
		return time.Time{}, false
	}
	return exp.Time, true
}

func (token refreshTokenClaims) GetLoginSessionId() (int, bool) {
	id, err := strconv.Atoi(token.ID)
	if err != nil {
		return -1, false
	}
	return id, true
}

func (token refreshTokenClaims) GetIssuedAt() (time.Time, bool) {
	iat := token.IssuedAt
	if iat == nil {
		return time.Time{}, false
	}
	return iat.Time, true
}

func (token refreshTokenClaims) GetExpiresAt() (time.Time, bool) {
	exp := token.ExpiresAt
	if exp == nil {
		return time.Time{}, false
	}
	return exp.Time, true
}

func (token resetPasswordTokenClaims) GetEmailCredentialId() (int, bool) {
	id, err := strconv.Atoi(token.ID)
	if err != nil {
		return -1, false
	}
	return id, true
}

func (token resetPasswordTokenClaims) GetExpiresAt() (time.Time, bool) {
	exp := token.ExpiresAt
	if exp == nil {
		return time.Time{}, false
	}
	return exp.Time, true
}
