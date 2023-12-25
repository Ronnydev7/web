package rule

import (
	"context"

	"api/ent/entgenerated/privacy"
	"api/privacy/viewer"
)

func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		v := viewer.FromContext(ctx)
		if v != nil && v.IsAdmin() {
			return privacy.Allow
		}
		return privacy.Skip
	})
}
