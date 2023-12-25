package intl

import (
	"api/intl/intlgenerated"

	"github.com/golang-jwt/jwt/v4"
)

type (
	ExpiredJwtTokenError struct {
		IntlError
	}

	InvalidJwtTokenSignatureError struct {
		IntlError
	}

	InvalidJwtTokenError struct {
		IntlError
	}

	UntranslatedJwtValidationError struct {
		UntranslatedError
		Inner error
	}
)

func HandleErrorFromJwt(err error) IntlError {
	if validationError, ok := err.(*jwt.ValidationError); ok {
		return HandleJwtValidationError(validationError)
	}
	return &UntranslatedJwtValidationError{
		Inner: err,
	}
}

func HandleJwtValidationError(err *jwt.ValidationError) IntlError {
	if err.Errors == jwt.ValidationErrorExpired {
		return &ExpiredJwtTokenError{}
	} else if err.Errors == jwt.ValidationErrorSignatureInvalid {
		return &InvalidJwtTokenSignatureError{}
	}

	if innerError, ok := err.Inner.(IntlError); ok {
		return innerError
	}

	return &UntranslatedJwtValidationError{
		Inner: err.Inner,
	}
}

func (ExpiredJwtTokenError) Error() string {
	return "expired jwt token"
}

func (ExpiredJwtTokenError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__TOKEN_EXPIRED
}

func (InvalidJwtTokenSignatureError) Error() string {
	return "invalid jwt token signature"
}

func (InvalidJwtTokenSignatureError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__INVALID_JWT_TOKEN
}

func (InvalidJwtTokenError) Error() string {
	return "invalid jwt token"
}

func (InvalidJwtTokenError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__INVALID_JWT_TOKEN
}

func (err UntranslatedJwtValidationError) Error() string {
	return err.Inner.Error()
}

func (UntranslatedJwtValidationError) GetIntlKey() string {
	return intlgenerated.COMMON_STRINGS__INVALID_JWT_TOKEN
}

func (err UntranslatedJwtValidationError) GetInner() error {
	return err.Inner
}
