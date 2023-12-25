package rule

import (
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/token"
	"context"
	"reflect"
)

func AllowAfterApplyingPrivacyTokenFilter(
	emptyToken token.PrivacyToken,
	applyFilter func(t token.PrivacyToken, filter privacy.Filter),
) privacy.QueryMutationRule {
	return privacy.FilterFunc(
		func(ctx context.Context, filter privacy.Filter) error {
			actualToken := ctx.Value(emptyToken.GetContextKey())
			actualTokenType := reflect.TypeOf(actualToken)
			expectedTokenType := reflect.TypeOf(emptyToken)
			if actualTokenType == expectedTokenType {
				applyFilter(actualToken.(token.PrivacyToken), filter)
				return privacy.Allowf("applied privacy token filter")
			}
			return privacy.Skipf("no token found from context with type %T", emptyToken)
		})
}
