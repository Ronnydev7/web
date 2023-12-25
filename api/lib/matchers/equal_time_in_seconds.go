package matchers

import (
	"fmt"
	"time"

	"github.com/onsi/gomega"
)

type equalTimeMatcherInSeconds struct {
	gomega.OmegaMatcher
	expected time.Time
}

func EqualTimeInSeconds(expected time.Time) gomega.OmegaMatcher {
	return &equalTimeMatcherInSeconds{
		expected: expected,
	}
}

func (matcher equalTimeMatcherInSeconds) Match(actual interface{}) (bool, error) {
	actualTime, ok := actual.(time.Time)
	if !ok {
		return false, fmt.Errorf("expecting a time.Time struct value. Got %T", actual)
	}
	if actualTime.Unix() == matcher.expected.Unix() {
		return true, nil
	}
	return false, fmt.Errorf(
		"expecting unix time of %d (%v); got %d (%v)",
		matcher.expected.Unix(),
		matcher.expected,
		actualTime.Unix(),
		actualTime,
	)
}
