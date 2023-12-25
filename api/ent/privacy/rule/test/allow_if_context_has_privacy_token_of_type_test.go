package test

import (
	"api/ent/entgenerated/privacy"
	"api/ent/internal"
	"api/ent/privacy/rule"
	"api/ent/privacy/token"
	"context"
	"strings"
	"testing"
)

func TestAllowIfContextHasPrivacyTokenOfTypeRulePassWhenContextHasToken(t *testing.T) {
	c := context.Background()
	c = token.NewContextWithPerformingLoginToken(c, "test@example.com")
	rule := rule.AllowIfContextHasPrivacyTokenOfType(&token.PerformingLoginToken{})
	client := internal.CreateEntClientForTest(t)
	defer client.Close()
	result := rule.EvalQuery(c, client.User.Query())
	if result != privacy.Allow {
		t.Fatalf("expected to get Allow decision, received %v", result)
	}
}

func TestAllowIfContextHasPrivacyTokenOfTypeRuleRuleSkipWhenContextHasNoToken(t *testing.T) {
	c := context.Background()
	rule := rule.AllowIfContextHasPrivacyTokenOfType(token.PerformingLoginToken{})
	client := internal.CreateEntClientForTest(t)
	defer client.Close()
	result := rule.EvalQuery(c, client.User.Query())
	if !strings.Contains(result.Error(), privacy.Skip.Error()) {
		t.Fatalf("expected to get SKip decision, received %v", result)
	}
}
