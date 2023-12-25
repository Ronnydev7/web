package rule

import (
	"api/ent/entgenerated/privacy"
	"context"
)

func AlwaysSkipRule() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(_ context.Context) error {
		return privacy.Skip
	})
}
