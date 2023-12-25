package rule

import (
	"api/ent/entgenerated/privacy"
	"api/privacy/viewer"
	"context"
)

func AllowIfSuperuser() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		v := viewer.FromContext(ctx)
		if v.IsSuperuser() {
			return privacy.Allow
		}
		return privacy.Skip
	})
}
