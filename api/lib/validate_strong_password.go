package lib

import (
	"api/intl/intlgenerated"
	"errors"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

const REQUIRED_PASSWORD_ENTROPY = 60

func ValidateStrongPassword(password string) error {
	entropy := passwordvalidator.GetEntropy(password)
	if entropy < REQUIRED_PASSWORD_ENTROPY {
		return errors.New(intlgenerated.EMAIL_CREDENTIAL__WEAK_PASSWORD)
	}
	return nil
}
