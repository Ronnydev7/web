package intl

import "errors"

type (
	IntlError interface {
		error
		GetIntlKey() string
	}
	UntranslatedError interface {
		IntlError
		GetInner() error
	}
)

func HandleIntlError(err error, fallbackKey string) error {
	intlError, ok := err.(IntlError)
	if ok {
		return errors.New(intlError.GetIntlKey())
	}
	return errors.New(fallbackKey)
}
