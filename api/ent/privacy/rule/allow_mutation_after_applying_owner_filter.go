package rule

import (
	"api/ent/entgenerated/predicate"
	"api/ent/entgenerated/privacy"
	"api/ent/entgenerated/user"
	"api/privacy/viewer"
	"context"
)

func AllowMutationAfterApplyingOwnerFilter() privacy.MutationRule {
	type OwnerFilter interface {
		WhereHasOwnerWith(predicates ...predicate.User)
	}

	return privacy.FilterFunc(
		func(ctx context.Context, f privacy.Filter) error {
			v := viewer.FromContext(ctx)
			ownerFilter, ok := f.(OwnerFilter)
			if !ok {
				return privacy.Deny
			}

			viewerId, exists := v.GetId()
			if !exists {
				return privacy.Skip
			}
			ownerFilter.WhereHasOwnerWith(user.ID(viewerId))
			return privacy.Allowf("applied owner filter")
		},
	)
}
