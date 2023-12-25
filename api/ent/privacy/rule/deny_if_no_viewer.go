package rule

import (
	"api/ent/entgenerated/privacy"
	"api/privacy/viewer"
	"context"
)

func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		v := viewer.FromContext(ctx)
		if v == nil {
			return privacy.Denyf("viewer is missing")
		}
		return privacy.Skip
	})
}
