package internal

import (
	"errors"
	"testing"
)

func TestingMatchPrivacyDecision(t *testing.T, actual error, expected error) {
	if !errors.Is(actual, expected) {
		t.Fatalf("expected decision '%v', received '%v'", expected, actual)
	}
}
