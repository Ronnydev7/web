package validator

import (
	"fmt"
	"unicode/utf8"
)

func MaxUtf8RuneCount(fieldName string, count int) func(string) error {
	return func(value string) error {
		if utf8.RuneCountInString(value) > count {
			return fmt.Errorf("%s must not exceed %d characters", fieldName, count)
		}
		return nil
	}
}
