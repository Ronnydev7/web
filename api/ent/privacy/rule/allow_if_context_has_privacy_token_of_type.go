package rule

import (
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/token"
	"context"
	"reflect"
)

func AllowIfContextHasPrivacyTokenOfType(emptyToken token.PrivacyToken) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		actualTokenType := reflect.TypeOf(ctx.Value(emptyToken.GetContextKey()))
		expectedTokenType := reflect.TypeOf(emptyToken)
		if actualTokenType == expectedTokenType {
			return privacy.Allow
		}
		return privacy.Skipf("no token found from context with type %T", emptyToken)
	})
}
