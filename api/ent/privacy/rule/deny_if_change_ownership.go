package rule

import (
	"api/ent/entgenerated"
	"api/ent/entgenerated/privacy"
	"context"
)

func DenyIfChangeOwnership() privacy.MutationRule {
	type UserOwnedEntMutation interface {
		OldOwnerID(context.Context) (id int, err error)
		OwnerID() (id int, exists bool)
	}
	return privacy.MutationRuleFunc(
		func(ctx context.Context, mutation entgenerated.Mutation) error {
			userOwnedMutation, ok := mutation.(UserOwnedEntMutation)
			if !ok {
				return privacy.Skipf("unexpected mutation type %T", mutation)
			}

			oldOwnerId, err := userOwnedMutation.OldOwnerID(ctx)
			if err != nil {
				return privacy.Skipf("unable to obtain old owner ID: %v", err)
			}

			newOwnerId, exists := userOwnedMutation.OwnerID()
			if !exists {
				return privacy.Skipf("not setting a new owner")
			}

			if oldOwnerId != newOwnerId {
				return privacy.Denyf("changing ownership is forbidden")
			}

			return privacy.Skip
		},
	)
}
