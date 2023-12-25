package validator

import (
	"fmt"
	"regexp"
)

func IsAlphanumericUnderscore(fieldName string) func(string) error {
	alphanumericUnderscoreRegex := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	return func(value string) error {
		if alphanumericUnderscoreRegex.MatchString(value) {
			return nil
		}
		return fmt.Errorf("%s must only consist of alphanumeric or underscore characters", fieldName)
	}
}
