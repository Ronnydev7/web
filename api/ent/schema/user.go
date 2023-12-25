package schema

import (
	"api/ent/entgenerated/privacy"
	"api/ent/privacy/rule"
	"api/ent/privacy/token"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
		),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("superuser_profile", SuperuserProfile.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			).
			Unique(),
		edge.
			To("email_credential", EmailCredential.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			).
			Unique(),
		edge.
			To("login_sessions", LoginSession.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.
			To("public_profile", UserPublicProfile.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			).
			Unique(),
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfContextHasPrivacyTokenOfType(&token.EmailSignupToken{}),
					rule.DenyIfNoViewer(),
					rule.AllowIfAdmin(),
					privacy.AlwaysDenyRule(),
				},
				ent.OpCreate,
			),
			privacy.OnMutationOperation(
				// TODO implement deletion policy
				privacy.MutationPolicy{
					rule.AllowIfAdmin(),
					privacy.AlwaysDenyRule(),
				},
				ent.OpUpdateOne|ent.OpUpdate|ent.OpDeleteOne|ent.OpDelete,
			),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
