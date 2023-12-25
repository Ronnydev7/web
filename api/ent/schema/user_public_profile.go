package schema

import (
	"api/ent/entgenerated/privacy"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

/**
 * This ent is always publicly accessible
 */
type UserPublicProfile struct {
	ent.Schema
}

func (UserPublicProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("handle_name").Unique(),
		field.String("photo_blob_key").
			Unique().
			Annotations(
				entgql.Skip(),
			).
			Optional(),
	}
}

func (UserPublicProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserOwnedMixin{
			Ref: "public_profile",
		},
		UserOwnedMutationPolicyMixin{
			AllowAdminMutation: true,
			AllowChangeOwner:   false,
		},
	}
}

func (UserPublicProfile) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
