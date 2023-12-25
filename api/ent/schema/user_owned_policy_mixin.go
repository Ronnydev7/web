package schema

import (
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/rule"
	"api/ent/privacy/utils"

	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
)

type (
	UserOwnedMutationPolicyMixin struct {
		mixin.Schema
		AllowAdminMutation bool
		AllowChangeOwner   bool
	}

	UserOwnedQueryPolicyMixin struct {
		mixin.Schema
	}
)

func (mixin UserOwnedMutationPolicyMixin) Policy() ent.Policy {
	adminPolicy := rule.AllowIfSuperuser()
	if mixin.AllowAdminMutation {
		adminPolicy = rule.AllowIfAdmin()
	}

	var denyIfChangeOwnerRule privacy.MutationRule
	if mixin.AllowChangeOwner {
		denyIfChangeOwnerRule = rule.DenyIfChangeOwnership()
	}

	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				utils.NewMutationPolicyWithoutNil(privacy.MutationPolicy{
					rule.DenyIfNoViewer(),
					adminPolicy,
					denyIfChangeOwnerRule,
					rule.AllowMutationAfterApplyingOwnerFilter(),
					privacy.AlwaysDenyRule(),
				}),
				ent.OpCreate,
			),
			privacy.OnMutationOperation(
				utils.NewMutationPolicyWithoutNil(privacy.MutationPolicy{
					rule.DenyIfNoViewer(),
					adminPolicy,
					denyIfChangeOwnerRule,
					rule.AllowMutationAfterApplyingOwnerFilter(),
					privacy.AlwaysDenyRule(),
				}),
				ent.OpUpdateOne|ent.OpUpdate|ent.OpDeleteOne|ent.OpDelete,
			),
		},
	}
}

func (mixin UserOwnedQueryPolicyMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.AllowIfAdmin(),
			rule.AllowIfOwnedByViewer(),
			privacy.AlwaysDenyRule(),
		},
	}
}
