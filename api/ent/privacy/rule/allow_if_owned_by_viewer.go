package rule

import (
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/privacy"
	"api/ent/entgenerated/user"
	"api/privacy/viewer"
	"context"
)

func AllowIfOwnedByViewer() privacy.QueryMutationRule {
	type UserOwnedFilter interface {
		WhereHasOwnerWith(...predicate.User)
	}

	return privacy.FilterFunc(
		func(ctx context.Context, filter privacy.Filter) error {
			v := viewer.FromContext(ctx)
			if v == nil {
				return privacy.Skipf("missing viewer in context")
			}
			viewerId, exists := v.GetId()
			if !exists {
				return privacy.Skipf("anonymous viewer")
			}

			actualFilter, ok := filter.(UserOwnedFilter)
			if !ok {
				return privacy.Denyf("unexpected filter type %T", filter)
			}

			actualFilter.WhereHasOwnerWith(user.ID(viewerId))
			return privacy.Allowf("applied owner filter")
		},
	)
}
