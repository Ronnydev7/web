package lib

import (
	"api/config"
	"api/intl"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	EMAIL_SIGNUP_TOKEN_TTL_HOUR = 1
	EMAIL_CLAIM_KEY             = "email"
	EXPIRATION_CLAIM_KEY        = "exp"
)

type (
	EmailSignupTokenClaims interface {
		GetEmail() (string, bool)
		GetExpiresAt() (time.Time, bool)
	}
	emailSignupTokenClaims struct {
		EmailSignupTokenClaims `json:"-"`
		jwt.StandardClaims
	}

	EmailSignupJwtTokenFactory interface {
		Create(email string) (string, intl.IntlError)
		Parse(stringToken string) (EmailSignupTokenClaims, intl.IntlError)
		GetSignatureSecret(claims EmailSignupTokenClaims) ([]byte, intl.IntlError)
	}

	EmailSignupJwtTokenFactoryWithConfig struct {
		EmailSignupJwtTokenFactory
		hmacConfig config.HmacConfig
	}

	NewEmailSignupJwtTokenFactoryFunc func(config.HmacConfig) EmailSignupJwtTokenFactory
)

var NewEmailSignupJwtTokenFactory NewEmailSignupJwtTokenFactoryFunc = func(c config.HmacConfig) EmailSignupJwtTokenFactory {
	return &EmailSignupJwtTokenFactoryWithConfig{
		hmacConfig: c,
	}
}

func (t emailSignupTokenClaims) GetEmail() (string, bool) {
	if t.Audience == "" {
		return "", false
	}
	return t.Audience, true
}

func (t emailSignupTokenClaims) GetExpiresAt() (time.Time, bool) {
	exp := t.ExpiresAt
	if exp == 0 {
		return time.Time{}, false
	}
	return time.Unix(exp, 0), true
}

func (builder EmailSignupJwtTokenFactoryWithConfig) Create(email string) (string, intl.IntlError) {
	expUnixTime := TimeNow().Add(time.Hour * EMAIL_SIGNUP_TOKEN_TTL_HOUR).Unix()
	claims := emailSignupTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  email,
			ExpiresAt: expUnixTime,
		},
	}
	emailSignupToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := builder.GetSignatureSecret(&claims)
	if err != nil {
		return "", err
	}
	result, jwtErr := emailSignupToken.SignedString(secret)
	if jwtErr != nil {
		return "", intl.HandleErrorFromJwt(jwtErr)
	}
	return result, nil
}

func (builder EmailSignupJwtTokenFactoryWithConfig) Parse(stringToken string) (EmailSignupTokenClaims, intl.IntlError) {
	claims := emailSignupTokenClaims{}
	_, err := jwt.ParseWithClaims(stringToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &intl.InvalidJwtTokenError{}
		}

		return builder.GetSignatureSecret(&claims)
	})
	if err != nil {
		return nil, intl.HandleErrorFromJwt(err)
	}

	return &claims, nil
}

func (builder EmailSignupJwtTokenFactoryWithConfig) GetSignatureSecret(claims EmailSignupTokenClaims) ([]byte, intl.IntlError) {
	email, exists := claims.GetEmail()
	if !exists {
		return nil, &intl.InvalidJwtTokenError{}
	}

	expiresAt, exists := claims.GetExpiresAt()
	if !exists {
		return nil, &intl.InvalidJwtTokenError{}
	}

	hmacSecret := builder.hmacConfig.GetEmailSignupTokenSecret()
	hmacData := fmt.Sprintf("%s:%d", email, expiresAt.Unix())

	return GetHmac([]byte(hmacSecret), []byte(hmacData)), nil
}
