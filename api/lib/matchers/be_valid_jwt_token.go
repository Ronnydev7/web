package matchers

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/onsi/gomega"
)

type beValidJwtTokenMatcher struct {
	gomega.OmegaMatcher
	expectedSignature []byte
}

func BeValidJwtToken(expectedSignature []byte) gomega.OmegaMatcher {
	return &beValidJwtTokenMatcher{expectedSignature: expectedSignature}
}

func (matcher beValidJwtTokenMatcher) Match(actual interface{}) (bool, error) {
	actualString, ok := actual.(string)
	if !ok {
		return false, errors.New("a JWT token must be a string")
	}

	_, err := jwt.Parse(actualString, func(_ *jwt.Token) (interface{}, error) {
		return matcher.expectedSignature, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
