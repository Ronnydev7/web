package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

type SuperuserProfile struct {
	ent.Schema
}

func (SuperuserProfile) Fields() []ent.Field {
	return nil
}

func (SuperuserProfile) Edges() []ent.Edge {
	return nil
}

func (SuperuserProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserOwnedMixin{
			Ref: "superuser_profile",
		},
		UserOwnedQueryPolicyMixin{},
		UserOwnedMutationPolicyMixin{
			AllowAdminMutation: false,
		},
	}
}

func (SuperuserProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(),
	}
}
